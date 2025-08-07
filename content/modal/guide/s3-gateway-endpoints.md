# S3 Gateway endpoints Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/s3-gateway-endpoints
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/s3-gateway-endpoints

## Content Preview

Workloads running on Modal should not incur egress or ingress fees associated
with S3 operations. No configuration is needed in order for your app to use S3 Gateway endpoints.
S3 Gateway endpoints are automatically used when your app runs on AWS.

Avoid specifying regional endpoints manually, as this can lead to unexpected cost
or performance degradation.

S3 Gateway endpoints guarantee no costs for network traffic within the same AWS region.
However, if your Modal Function runs in one region but your bucket resides in a
different region you will be billed for inter-region traffic.

