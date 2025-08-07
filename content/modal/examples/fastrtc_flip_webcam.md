# Run a FastRTC app on Modal Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/fastrtc_flip_webcam
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/fastrtc_flip_webcam

## Content Preview

WebRTC provides low latency (“real-time”) peer-to-peer communication
for Web applications, focusing on audio and video.
Considering that the Web is a platform originally designed
for high-latency, client-server communication of text and images,
that’s no mean feat!

But before we do that, we need to consider limits:
on how many peers can connect to one instance on Modal
and on how long they can stay connected.
We picked some sensible defaults to show how they interact
with the deployment parameters of the Modal Function.
You’ll want to tune these for your application!

To try this out for yourself, run

This temporary deployment is tied to your terminal session.
To deploy permanently, run

This FastRTC app is very much the “hello world” or “echo server”
of FastRTC: it just flips the incoming webcam stream and adds a “hello” message.
That logic appears below.

