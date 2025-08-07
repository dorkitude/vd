# Images Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/images
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/images

## Content Preview

This guide walks you through how to define the environment your Modal Functions run in.

Note that because you can define a different environment for each and every
Modal Function if you so choose, you don’t need to worry about virtual
environment management. Containers make for much better separation of concerns!

You can also supply shell commands that should be executed when building the
container image.

You might use this to preload custom assets, like model parameters, so that they
don’t need to be retrieved when Functions start up:

Essentially, this is equivalent to running a Modal Function and snapshotting the
resulting filesystem as an image.

