# DoppelBot: Fine-tune an LLM to replace your CEO Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/llm-finetuning
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/llm-finetuning

## Content Preview

Initial versions of the model were prone to generating short responses
— unsurprising, because a majority of Slack communication is pretty terse.
Adding a minimum character length for the target user’s messages fixed this.

If you’re following along at home, you can run the scraper with the following
command:

Because of the typically small sample sizes we’re working with, training for
longer than a couple hundred steps (with our batch size of 128) quickly led to
overfitting. Admittedly, we haven’t thoroughly evaluated the hyperparameter
space yet — do reach out to us if you’re interested in collaborating on this!

To try this step yourself, run:

With parametrized functions, every user model gets its own pool of containers
that scales up when there are incoming requests, and scales to 0 when there’s
none. Here’s what that looks like stripped down to the essentials:

