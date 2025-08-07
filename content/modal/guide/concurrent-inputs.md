# Input concurrency Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/concurrent-inputs
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/concurrent-inputs

## Content Preview

As traffic to your application increases, Modal will automatically scale up the
number of containers running your Function:

By default, each container will be assigned one input at a time. Autoscaling
across containers allows your Function to process inputs in parallel. This is
ideal when the operations performed by your Function are CPU-bound.

When used effectively, input concurrency can reduce latency and lower costs.

Input concurrency can be especially effective for workloads that are primarily
I/O-bound, e.g.:

For such workloads, individual containers may be able to concurrently process
large numbers of inputs with minimal additional latency. This means that your
Modal application will be more efficient overall, as it won’t need to scale
containers up and down as traffic ebbs and flows.

