import * as pulumi from '@pulumi/pulumi'
import { CreateResult, UpdateResult } from '@pulumi/pulumi/dynamic'
import { S3Client, CreateBucketCommand, DeleteBucketCommand, PutBucketWebsiteCommand, PutPublicAccessBlockCommand, PutBucketAclCommand } from '@aws-sdk/client-s3'

export interface WebsiteBucketInputs {
    name: string
    commitHash: string
}

export interface WebsiteBucketResourceInputs {
    commitHash: string
}

export interface WebsiteBucketOutputs {
    currentBucketName: string
    previousBucketName: string
}

class WebsiteBucketProvider implements pulumi.dynamic.ResourceProvider {
  private async createBucket(client: S3Client, bucketName: string) {
    await client.send(new CreateBucketCommand({
        Bucket: bucketName,
        ObjectOwnership: 'ObjectWriter',
    }))

    await client.send(new PutPublicAccessBlockCommand({
        Bucket: bucketName,
        PublicAccessBlockConfiguration: {
            BlockPublicAcls: false,
            IgnorePublicAcls: false,
            BlockPublicPolicy: false,
            RestrictPublicBuckets: false,
        },
    }))

    await client.send(new PutBucketAclCommand({
        Bucket: bucketName,
        ACL: 'public-read',
    }))

    await client.send(new PutBucketWebsiteCommand({
        Bucket: bucketName,
        WebsiteConfiguration: {
            IndexDocument: {
                "Suffix": "index.html",
            },
        },
    }))
  }

  private async deleteBucket(client: S3Client, bucketName: string) {
    await client.send(new DeleteBucketCommand({
        Bucket: bucketName,
    }))
  }

  //*** CREATE ***//
  async create(inputs: WebsiteBucketInputs): Promise<CreateResult<WebsiteBucketOutputs>> {
    const client = new S3Client({ region: 'us-west-2' })
    const bucketName = `${inputs.name}-${inputs.commitHash}`
    await this.createBucket(client, bucketName)

    return {
        id: inputs.name,
        outs: {
            currentBucketName: bucketName,
            previousBucketName: '',
        },
    }
  }

  //*** UPDATE ***//
  async update(id: string, olds: WebsiteBucketOutputs, news: WebsiteBucketInputs): Promise<UpdateResult<WebsiteBucketOutputs>> {
    const client = new S3Client({ region: 'us-west-2' })
    const bucketName = `${news.name}-${news.commitHash}`
    await this.createBucket(client, bucketName)

    if (olds.previousBucketName && olds.previousBucketName !== '') {
        await this.deleteBucket(client, olds.previousBucketName)
    }

    return {
        outs: {
            currentBucketName: bucketName,
            previousBucketName: olds.currentBucketName,
        }
    }
  }

  //*** DELETE ***//
  async delete(id: string, outputs: WebsiteBucketOutputs) {
    const client = new S3Client({ region: 'us-west-2' })
    await this.deleteBucket(client, outputs.currentBucketName)
    await this.deleteBucket(client, outputs.previousBucketName)
  }
}

export class WebsiteBucket extends pulumi.dynamic.Resource {
  // @ts-ignore
  public readonly currentBucketName: pulumi.Output<string>
  // @ts-ignore
  public readonly previousBucketName: pulumi.Output<string>

  constructor(name: string, args: WebsiteBucketResourceInputs, opts?: pulumi.CustomResourceOptions) {
    super(new WebsiteBucketProvider(), name, {
        name,
        ...args,
        currentBucketName: undefined,
        previousBucketName: undefined,
    }, opts);
  }
}
