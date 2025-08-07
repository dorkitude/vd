# Volumes Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/volumes
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/volumes

## Content Preview

Modal Volumes provide a high-performance distributed file system for your Modal applications.
They are designed for write-once, read-many I/O workloads, like creating machine learning model
weights and distributing them for inference.

You can also browse and manipulate Volumes from an ad hoc Modal Shell:

While there’s no file size limit for individual files in a volume, the frontend only supports downloading files up to 16 MB. For larger files, please use the CLI:

You can also create Volumes lazily from code using:

This will create the Volume if it doesn’t exist.

