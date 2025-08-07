# Stream transcriptions with Kyutai STT Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/streaming_kyutai_stt
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/streaming_kyutai_stt

## Content Preview

This example demonstrates the deployment of a streaming audio transcription service with Kyutai STT on Modal.

We start by importing some basic packages and the Modal SDK.

One dependency is missing: the model weights.

Now we’re ready to add the code that runs the speech-to-text model.

That plus the code for manipulating the streams of audio bytes and output text
leads to a pretty big class! But there’s not anything too complex here.

