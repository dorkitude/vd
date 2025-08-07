# GPU acceleration Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/gpu
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/gpu

## Content Preview

Here’s a simple example of a Function running on an A100 in Modal:

Currently B200, H200, H100, A100, L4, T4 and L40S instances support up to 8 GPUs (up to 1,536 GB GPU RAM),
and A10 instances support up to 4 GPUs (up to 96 GB GPU RAM). Note that requesting
more than 2 GPUs per container will usually result in larger wait times. These
GPUs are always attached to the same physical machine.

Modal allows specifying a list of possible GPU types, suitable for Functions that are
compatible with multiple options. Modal respects the ordering of this list and
will try to allocate the most preferred GPU type before falling back to less
preferred ones.

Or take a look some examples of Modal apps using GPUs:

