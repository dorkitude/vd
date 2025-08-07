# Fold proteins with Chai-1 Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/chai1
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/chai1

## Content Preview

In biology, function follows form quite literally:
the physical shapes of proteins dictate their behavior.
Measuring those shapes directly is difficult
and first-principles physical simulation prohibitively expensive.

This simple script is meant as a starting point showing how to handle fiddly bits
like installing dependencies, loading weights, and formatting outputs so that you can get on with the fun stuff.
To experience the full power of Modal, try scaling inference up and running on hundreds or thousands of structures!

The logic for running Chai-1 is encapsulated in the function below,
which you can trigger from the command line by running

To learn how it works, read on!

Not all “dependencies” belong in a container image. Chai-1, for example, depends on
the weights of several models.

