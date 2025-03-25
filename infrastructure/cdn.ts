import * as aws from '@pulumi/aws'
import * as awsx from '@pulumi/awsx'
import * as pulumi from '@pulumi/pulumi'

export interface CdnBucket {
    name: string
    arn: string
    path: string
    domain: string
}

export interface CdnArgs {
  buckets: pulumi.Output<CdnBucket>[]
}

// CachingDisabled sets min, max, and default cache TTLs to 0.
// https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html
const cachingDisabledId = '4135ea2d-6df8-44a3-9df3-4b5a84be39ad'

export class Cdn extends pulumi.ComponentResource {
  public readonly cloudfrontDomain: pulumi.Output<string>

  constructor(
    name: string,
    args: CdnArgs,
    opts?: pulumi.ComponentResourceOptions,
  ) {
    super('custom:resource:Cdn', name, args, opts)

    const defaultCacheBehavior: aws.cloudfront.DistributionArgs['defaultCacheBehavior'] =
      {
        targetOriginId: pulumi.all(args.buckets).apply(buckets => buckets[0].arn),
        compress: true,
        viewerProtocolPolicy: 'redirect-to-https',
        allowedMethods: ['GET', 'HEAD', 'OPTIONS'],
        cachedMethods: ['GET', 'HEAD', 'OPTIONS'],
        forwardedValues: {
          cookies: {
            forward: 'all',
          },
          queryString: true,
        },
        minTtl: 0,
        defaultTtl: 0,
        maxTtl: 0,
      }

    const origins: aws.cloudfront.DistributionArgs['origins'] = pulumi.all(args.buckets).apply(buckets => {
        return buckets.map(b => ({
            originId: b.arn,
            domainName: b.domain,
            customOriginConfig: {
                originProtocolPolicy: 'http-only',
                httpPort: 80,
                httpsPort: 443,
                originSslProtocols: ['TLSv1.2'],
            }
        }))
    })

    const orderedCacheBehaviors: aws.cloudfront.DistributionArgs['orderedCacheBehaviors'] = pulumi.all(args.buckets).apply(buckets => {
      const behaviors: aws.cloudfront.DistributionArgs['orderedCacheBehaviors'] = []
      for (const b of buckets) {
        behaviors.push({ ...defaultCacheBehavior, targetOriginId: b.arn, pathPattern: b.path })
        behaviors.push({ ...defaultCacheBehavior, targetOriginId: b.arn, pathPattern: `${b.path}/*` })
      }

      return behaviors
    })

    const distributionArgs: aws.cloudfront.DistributionArgs = {
      origins,
      enabled: true,

      defaultRootObject: 'index.html',

      defaultCacheBehavior,
      orderedCacheBehaviors,

      priceClass: 'PriceClass_All',

      restrictions: {
        geoRestriction: {
          restrictionType: 'none',
        },
      },

      viewerCertificate: {
        cloudfrontDefaultCertificate: true,
        sslSupportMethod: 'sni-only',
        minimumProtocolVersion: 'TLSv1.2_2018',
      },

      customErrorResponses: [],
    }

    const cdn = new aws.cloudfront.Distribution(
      'buildite-sdk-docs-cdn',
      distributionArgs,
      { parent: this },
    )

    this.cloudfrontDomain = cdn.domainName
    this.registerOutputs({})
  }
}
