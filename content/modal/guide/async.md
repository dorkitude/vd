# Asynchronous API usage Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/async
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/async

## Content Preview

This is an advanced feature. If you are comfortable with asynchronous
programming, you can use this to create arbitrary parallel execution patterns,
with the added benefit that any Modal functions will be executed remotely.

An async function can call a blocking function, and vice versa.

If a function is configured to support multiple concurrent inputs per container,
the behavior varies slightly between blocking and async contexts:

