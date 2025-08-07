# Run Flux fast on H100s with torch.compile Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/flux
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/flux

## Content Preview

PyTorch added faster attention kernels for Hopper GPUs in version 2.5.

Next, we map the model’s setup and inference code onto Modal.

By default, we do some basic optimizations, like adjusting memory layout
and re-expressing the attention head projections as a single matrix multiplication.
But there are additional speedups to be had!

PyTorch 2 added a compiler that optimizes the
compute graphs created dynamically during PyTorch execution.
This feature helps close the gap with the performance of static graph frameworks
like TensorRT and TensorFlow.

You can run this example on Modal in 60 seconds.

