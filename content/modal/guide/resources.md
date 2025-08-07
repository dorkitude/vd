# Reserving CPU and memory Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/resources
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/resources

## Content Preview

Each Modal container has a default reservation of 0.125 CPU cores and 128 MiB of memory.
Containers can exceed this minimum if the worker has available CPU or memory.
You can also guarantee access to more resources by requesting a higher reservation.

As the platform grows, we plan to support larger CPU and memory reservations.

For CPU and memory, you’ll be charged based on whichever is higher: your reservation or actual usage.

Disk requests are billed by increasing the memory request at a 20:1 ratio. For example, requesting 500 GiB of disk will increase the memory request to 25 GiB, if it is not already set higher.

Modal containers have a default soft CPU limit that is set at 16 physical cores above the CPU request.
Given that the default CPU request is 0.125 cores the default soft CPU limit is 16.125 cores.
Above this limit the host will begin to throttle the CPU usage of the container.

