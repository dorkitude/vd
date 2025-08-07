# Embed 30 million Amazon reviews at 575k tokens per second with Qwen2-7B Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/amazon_embeddings
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/amazon_embeddings

## Content Preview

This example demonstrates how to create embeddings for a large text dataset. This is
often necessary to enable semantic search, translation, and other language
processing tasks. Modal makes it easy to deploy large, capable embedding models and handles
all of the scaling to process very large datasets in parallel on many cloud GPUs.

We create a Modal Function that will handle all of the data loading and submit inputs to an
inference Cls that will automatically scale up to handle hundreds of large
batches in parallel.

You can run it with the command

In it, we download the data we need and cache it to the container’s local disk,
which will disappear when the job is finished. We will be saving the review data
along with the embeddings, so we don’t need to keep the dataset around.

Once all of the batches have been sent for inference, we can return the function IDs
to the local client to save.

