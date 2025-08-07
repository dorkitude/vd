# Publish interactive datasets with Datasette Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/cron_datasette
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/cron_datasette

## Content Preview

Build and deploy an interactive movie database that automatically updates daily with the latest IMDb data.
This example shows how to serve a Datasette application on Modal with millions of movie and TV show records.

Along the way, we will learn how to use the following Modal features:

Let’s get started writing code.
For the Modal container image we need a few Python packages.

This dataset is no swamp, but a bit of data cleaning is still in order.
The following function reads a .tsv file, cleans the data and yields batches of records.

With the TSV processing out of the way, we’re ready to create a SQLite database and feed data into it.

