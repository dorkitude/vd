# Filesystem Access Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/sandbox-files
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/sandbox-files

## Content Preview

There are multiple options for uploading files to a Sandbox and accessing them
from outside the Sandbox.

The caller also can access files created in the Volume from the Sandbox, even after the Sandbox is terminated:

File syncing behavior differs between Volumes and CloudBucketMounts. For
Volumes, files are only synced back to the Volume when the Sandbox terminates.
For CloudBucketMounts, files are synced automatically.

If you’re less concerned with efficiency of uploads and want a convenient way
to pass data in and out of the Sandbox during execution, you can use our
filesystem API to easily read and write files. The API supports reading
files up to 100 MiB and writes up to 1 GiB in size.

This API is currently in Alpha, and we don’t recommend using it for production
workloads.

