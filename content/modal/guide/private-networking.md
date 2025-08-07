# Cluster networking Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/guide/private-networking
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/guide/private-networking

## Content Preview

i6pn (IPv6 private networking) is Modal’s private container-to-container networking solution. It allows users to create clusters of Modal containers which can send network traffic to each other with low latency and high bandwidth (≥ 50Gbps).

The upshot of this is that only containers in the same workspace can see each other and send each other network packets. i6pn networking is secure by default.

Consider having a container setup a Tunnel and act as the gateway to the private cluster networking.

