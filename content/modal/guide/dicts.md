# Dicts Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/dicts
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/dicts

## Content Preview

Modal Dicts provide distributed key-value storage to your Modal Apps.

Dicts provide distributed key-value storage to your Modal Apps.
Much like a standard Python dictionary, a Dict lets you store and retrieve
values using keys. However, unlike a regular dictionary, a Dict in Modal is
accessible from anywhere, concurrently and in parallel.

Dicts are persisted, which means that the data in the dictionary is
stored and can be retrieved even after the application is redeployed.

Python dicts can have keys of any hashable type and values of any type.

You can store Python objects of any serializable type within Dicts as keys or values.

