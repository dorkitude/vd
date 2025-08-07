# Create your own music samples with MusicGen Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/musicgen
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/musicgen

## Content Preview

MusicGen is a popular open-source music-generation model family from Meta.
In this example, we show you how you can run MusicGen models on Modal GPUs,
along with a Gradio UI for playing around with the model.

We start by defining the environment our generation runs in.
This takes some explaining since, like most cutting-edge ML environments, it is a bit fiddly.

In addition to source code, we’ll also need the model weights.

But Modal Functions are serverless: instances spin down when they aren’t being used.
If we want to avoid downloading the weights every time we start a new instance,
we need to store the weights somewhere besides our local filesystem.

We don’t need to change any of the model loading code —
we just need to make sure the model gets stored in the right directory.

