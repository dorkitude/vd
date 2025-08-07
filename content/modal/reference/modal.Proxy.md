# modal.Proxy Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/reference/modal.Proxy
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/reference/modal.Proxy

## Content Preview

Proxy objects give your Modal containers a static outbound IP address.

Synchronize the local object with its identity on the Modal server.

It is rarely necessary to call this method explicitly, as most operations
will lazily hydrate when needed. The main use case is when you need to
access object metadata, such as its ID.

Reference a Proxy by its name.

In contrast to most other Modal objects, new Proxy objects must be
provisioned via the Dashboard and cannot be created on the fly from code.

