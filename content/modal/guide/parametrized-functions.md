# Parametrized functions Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/parametrized-functions
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/parametrized-functions

## Content Preview

A single Modal Function can be parametrized by a set of arguments, so that each unique combination of arguments will behave like an individual
Modal Function with its own auto-scaling and lifecycle logic.

For example, you might want to have a separate pool of containers for each unique user that invokes your Function. In this scenario, you would
parametrize your Function by a user ID.

The parameters create a keyword-only constructor for your class, and the methods can be called as follows:

Some things to note:

Parameters are specified in the URL as query parameter values.

