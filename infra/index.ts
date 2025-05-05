import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as synced from "@pulumi/synced-folder";

const config = new pulumi.Config();
const path = config.require("path");
const indexDocument = config.require("indexDocument");
const errorDocument = config.require("errorDocument");

// Create an S3 bucket to hold the docs.
const bucket = new aws.s3.BucketV2("bucket", {
    tags: {
        Name: 'bk-sdk-docs-bucket',
        CostBucket: 'marketing-site',
        VantaNoAlert: 'Bucket intentionally public',
    },
});


const bucketWebsite = new aws.s3.BucketWebsiteConfigurationV2(
    "bucket-website",
    {
        bucket: bucket.bucket,
        indexDocument: { suffix: indexDocument },
        errorDocument: { key: errorDocument },
    }
);

// Configure ownership controls for the bucket.
const ownershipControls = new aws.s3.BucketOwnershipControls(
    "ownership-controls",
    {
        bucket: bucket.bucket,
        rule: {
            objectOwnership: "ObjectWriter",
        },
    }
);

// Configure a public access block for the bucket.
const publicAccessBlock = new aws.s3.BucketPublicAccessBlock(
    "public-access-block",
    {
        bucket: bucket.bucket,
        blockPublicAcls: false,
    }
);

// Use a synced folder to manage the files of the website.
new synced.S3BucketFolder(
    "bucket-folder",
    {
        path: path,
        bucketName: bucket.bucket,
        acl: "public-read",
        managedObjects: false,
    },
    { dependsOn: [ownershipControls, publicAccessBlock] }
);

// Create a CloudFront CDN to distribute and cache the website.
const tenMinutes = 60 * 10;
const cdn = new aws.cloudfront.Distribution("cdn", {
    enabled: true,
    origins: [
        {
            originId: bucket.arn,
            domainName: bucketWebsite.websiteEndpoint,
            customOriginConfig: {
                originProtocolPolicy: "http-only",
                httpPort: 80,
                httpsPort: 443,
                originSslProtocols: ["TLSv1.2"],
            },
        },
    ],
    defaultCacheBehavior: {
        targetOriginId: bucket.arn,
        viewerProtocolPolicy: "redirect-to-https",
        allowedMethods: ["GET", "HEAD", "OPTIONS"],
        cachedMethods: ["GET", "HEAD", "OPTIONS"],
        defaultTtl: tenMinutes,
        maxTtl: tenMinutes,
        minTtl: tenMinutes,
        forwardedValues: {
            queryString: false,
            cookies: {
                forward: "none",
            },
        },
    },
    customErrorResponses: [
        {
            errorCode: 404,
            responseCode: 404,
            responsePagePath: `/${errorDocument}`,
        },
    ],
    restrictions: {
        geoRestriction: {
            restrictionType: "none",
        },
    },
    viewerCertificate: {
        cloudfrontDefaultCertificate: true,
    },
});

// Export relevant URLs and hostnames.
export const bucketName = bucket.bucket;
export const bucketURI = pulumi.interpolate`s3://${bucket.bucket}`;
export const originURL = pulumi.interpolate`http://${bucketWebsite.websiteEndpoint}`;
export const originHostname = bucketWebsite.websiteEndpoint;
export const cdnURL = pulumi.interpolate`https://${cdn.domainName}`;
export const cdnHostname = cdn.domainName;
export const typescriptDocsURL = pulumi.interpolate`https://${cdn.domainName}/docs/sdk/typescript`;
export const pythonDocsURL = pulumi.interpolate`https://${cdn.domainName}/docs/sdk/python`;
export const rubyDocsURL = pulumi.interpolate`https://${cdn.domainName}/docs/sdk/ruby`;
