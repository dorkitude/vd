# Add Modal Apps to Tailscale Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/modal_tailscale
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/modal_tailscale

## Content Preview

We use a custom entrypoint to automatically add containers to a Tailscale network (tailnet).
This configuration enables the containers to interact with one another and with
additional applications within the same tailnet.

Packages might not be installed locally. This catches import errors and
only attempts imports in the container.

Configure Python to use the SOCKS5 proxy globally.

You can run this example on Modal in 60 seconds.

After creating a free account, install the Modal Python package, and
      create an API token.

