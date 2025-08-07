# Train an SLM from scratch with early-stopping grid search over hyperparameters Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/hp_sweep_gpt
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/hp_sweep_gpt

## Content Preview

When you want a language model that performs well on your task, there are three options,
ordered by the degree of customization:

Each step adds additional engineering complexity, but also leads to a superior cost-performance Pareto frontier
for your tasks. Fine-tuned models at one-tenth the size regularly outperform more generic models,
and models trained from scratch outperform them.

In this example, we will explore training an SLM from scratch on Modal.

In fact, we’ll train 8 SLMs in parallel with different hyperparameters
and then select the best one for additional training.

We’ll monitor this training live and serve our training and trained models
as web endpoints and simple browser UIs.

