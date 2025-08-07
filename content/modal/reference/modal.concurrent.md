# modal.concurrent Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/reference/modal.concurrent
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/reference/modal.concurrent

## Content Preview

Decorator that allows individual containers to handle multiple inputs concurrently.

The concurrency mechanism depends on whether the function is async or not:

Input concurrency will be most useful for workflows that are IO-bound
(e.g., making network requests) or when running an inference server that supports
dynamic batching.

