# Render a video with Blender on many GPUs or CPUs in parallel Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/blender_video
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/blender_video

## Content Preview

You can run it on CPUs to scale out on one hundred containers
or run it on GPUs to get higher throughput per node.
Even for this simple scene, GPUs render >10x faster than CPUs.

The final render looks something like this:

Modal runs your Python functions for you in the cloud.
You organize your code into apps, collections of functions that work together.

We define a function that renders a single frame. We’ll scale this function out on Modal later.

Functions in Modal are defined along with their hardware and their dependencies.
This function can be run with GPU acceleration or without it, and we’ll use a global flag in the code to switch between the two.

