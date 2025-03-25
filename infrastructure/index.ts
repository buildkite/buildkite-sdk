import * as pulumi from '@pulumi/pulumi'
import * as aws from '@pulumi/aws'
import { WebsiteBucket } from './bucket'
import { Cdn, CdnBucket } from './cdn'
import * as synced from '@pulumi/synced-folder'

const commitHash = 'stu'

const typescriptBucket = new WebsiteBucket('typescript-docs-bucket', {
    commitHash,
})

const typescriptSync = new synced.S3BucketFolder('typescript-folder', {
    path: '../dist/docs/typescript',
    bucketName: typescriptBucket.currentBucketName,
    acl: aws.s3.PublicReadAcl,
    managedObjects: false,
})

const pythonBucket = new WebsiteBucket('python-docs-bucket', {
    commitHash,
})

const pythonSync = new synced.S3BucketFolder('python-folder', {
    path: '../dist/docs/python',
    bucketName: pythonBucket.currentBucketName,
    acl: aws.s3.PublicReadAcl,
    managedObjects: false,
})

const rubyBucket = new WebsiteBucket('ruby-docs-bucket', {
    commitHash,
})

const rubySync = new synced.S3BucketFolder('ruby-folder', {
    path: '../dist/docs/ruby',
    bucketName: rubyBucket.currentBucketName,
    acl: aws.s3.PublicReadAcl,
    managedObjects: false,
})

function createBucketCdnArgs(path: string) {
    return async (bucketName: string): Promise<CdnBucket> => {
        const bucket = await aws.s3.getBucket({ bucket: bucketName })
        return {
            path,
            name: bucketName,
            arn: bucket.arn,
            domain: bucket.bucketRegionalDomainName,
        }
    }
}

const typescriptBucketCdnArgs = typescriptBucket.currentBucketName.apply(createBucketCdnArgs('/typescript'))
const pythonBucketCdnArgs = pythonBucket.currentBucketName.apply(createBucketCdnArgs('/python'))
const rubyBucketCdnArgs = rubyBucket.currentBucketName.apply(createBucketCdnArgs('/ruby'))

const cdn = new Cdn('buildkite-sdk-docs-distribution', {
    buckets: [
        typescriptBucketCdnArgs,
        pythonBucketCdnArgs,
        rubyBucketCdnArgs,
    ],
}, {
    dependsOn: [
        typescriptSync,
        pythonSync,
        rubySync,
    ]
})

export const endpoint = cdn.cloudfrontDomain
