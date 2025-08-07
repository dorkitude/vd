# modal.Volume Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/reference/modal.Volume
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/reference/modal.Volume

## Content Preview

A writeable volume that can be used to share files between one or more Modal functions.

The contents of a volume is exposed as a filesystem. You can use it to share data between different functions, or
to persist durable state across several instances of the same function.

Unlike a networked filesystem, you need to explicitly reload the volume to see changes made since it was mounted.
Similarly, you need to explicitly commit any changes you make to the volume for the changes to become visible
outside the current container.

Concurrent modification is supported, but concurrent modifications of the same files should be avoided! Last write
wins in case of concurrent modification of the same file - any data the last writer didn’t have when committing
changes will be lost!

As a result, volumes are typically not a good fit for use cases where you need to make concurrent modifications to
the same file (nor is distributed file locking supported).

