# Dynamic batching (beta) Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/dynamic-batching
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/dynamic-batching

## Content Preview

Batching increases throughput at a potential cost to latency.
Batched requests can share resources and reuse work, reducing the time and cost per request.
Batching is particularly useful for GPU-accelerated machine learning workloads,
as GPUs are designed to maximize throughput and are frequently bottlenecked on shareable resources,
like weights stored in memory.

Here’s what that looks like:

One additional rule applies to classes with Batched Methods:

To optimize the batching configurations for your application, consider the following heuristics:

Now, you can submit requests to the web endpoint and process them in batches. For instance, the three requests
in the following example, which might be requests from concurrent clients in a real deployment,
will be batched into two executions:

