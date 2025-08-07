# Run Qwen2-VL on SGLang for Visual QA Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/sgl_vlm
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/sgl_vlm

## Content Preview

Vision-Language Models (VLMs) are like LLMs with eyes:
they can generate text based not just on other text,
but on images as well.

Here’s a sample inference, with the image rendered directly (and at low resolution) in the terminal:

First, we’ll import the libraries we need locally
and define some constants.

VLMs are generally larger than LLMs with the same cognitive capability.
LLMs are already hard to run effectively on CPUs, so we’ll use a GPU here.
We find that inference for a single input takes about 3-4 seconds on an A10G.

Running an inference service on Modal is as easy as writing inference in Python.

