# Queues Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/queues
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/queues

## Content Preview

Modal Queues provide distributed FIFO queues to your Modal Apps.

Queues are particularly useful when you want to handle tasks or process
data asynchronously, or when you need to pass messages between different
components of your distributed system.

Queues are split into separate FIFO partitions via a string key. By default, one
partition (corresponding to an empty key) is used.

Python Queues can have values of any type.

Modal Queues can store Python objects of any serializable type.

