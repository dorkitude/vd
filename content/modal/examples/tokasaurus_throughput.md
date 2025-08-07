# High-throughput LLM inference with Tokasaurus (LLama 3.2 1B Instruct) Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/tokasaurus_throughput
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/tokasaurus_throughput

## Content Preview

In this example, we demonstrate how to use Tokasaurus, an LLM inference framework designed for maximum throughput.

On throughput-focused benchmarks with high prefix sharing workloads, Tokasaurus can outperform vLLM and SGLang nearly three-fold.

We start by maximizing the number of tokens processed per forward pass by adjusting the following parameters:

We could apply the Torch compiler to the model to make it faster and, via kernel fusion, reduce the amount of used activation memory,
leaving space for a larger KV cache. However, it dramatically increases the startup time of the server,
and we only see modest (20%, not 2x) improvements to throughput, so we don’t use it here.

Lastly, we need to set a few of the parameters for the client requests,
again based on the official benchmarking script.

