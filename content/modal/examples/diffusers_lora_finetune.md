# Fine-tune Flux on your pet using LoRA Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/diffusers_lora_finetune
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/diffusers_lora_finetune

## Content Preview

And with some light customization, you can use it to generate images of your pet!

We start by importing the necessary libraries and setting up the environment.

We start from a base image and specify all of our dependencies.
We’ll call out the interesting ones as they come up below.
Note that these dependencies are not installed locally
— they are only installed in the remote environment where our Modal App runs.

Machine learning apps often have a lot of configuration information.
We collect up all of our configuration into dataclasses to avoid scattering special/magic values throughout code.

Part of the magic of the low-rank fine-tuning is that we only need 3-10 images for fine-tuning.
So we can fetch just a few images, stored on consumer platforms like Imgur or Google Drive,
whenever we need them — no need for expensive, hard-to-maintain data pipelines.

