# Build a protein folding dashboard with ESM3, Molstar, and Gradio Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/esm3
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/esm3

## Content Preview

In this example, we’ll show how you can use Modal to not
just run the latest protein-folding model but also build tools around it for
you and your team of scientists to understand and analyze the results.

Next, we map the model’s setup and inference code onto Modal.

In this section we’ll create a web interface around the ESM3 model
that can help scientists and stakeholders understand and interrogate the results of the model.

You can deploy this UI, along with the backing inference endpoint,
with the following command:

The inference runs in a GPU-accelerated container with all of ESM3’s
dependencies, while this code executes in a CPU-only container
with only our web dependencies.

