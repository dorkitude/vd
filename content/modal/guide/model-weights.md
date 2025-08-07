# Storing model weights on Modal Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/model-weights
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/model-weights

## Content Preview

Efficiently managing the weights of large models is crucial for optimizing the
build times and startup latency of many ML and AI applications.

To store your model weights in a Volume, you need to either
make the Volume available to a Modal Function that saves the model weights
or upload the model weights into the Volume from a client.

If you’re already generating the weights on Modal, you just need to
attach the Volume to your Modal Function, making it available for reading and writing:

If the model weights are generated outside of Modal and made available
over the Internet, for example by an open-weights model provider
or your own training job on a dedicated cluster,
you can also download them into a Volume from a Modal Function:

Instead of pulling weights into a Modal Volume from inside a Modal Function,
you might wish to push weights into Modal from a client,
like your laptop or a dedicated training cluster.

