* * *

Copy page

# API Reference

This is the API reference for the [`modal`](https://pypi.org/project/modal/)
Python package, which allows you to run distributed applications on Modal.

The reference is intended to be limited to low-level descriptions of various
programmatic functionality. If you’re just getting started with Modal, we
would instead recommend looking at the [guide](/docs/guide) first or to get
started quickly with an [example](/docs/examples).

## Application construction

|
---|---
[`App`](/docs/reference/modal.App)| The main unit of deployment for code on
Modal
[`App.function`](/docs/reference/modal.App#function)| Decorator for
registering a function with an App
[`App.cls`](/docs/reference/modal.App#cls)| Decorator for registering a class
with an App

## Serverless execution

|
---|---
[`Function`](/docs/reference/modal.Function)| A serverless function backed by
an autoscaling container pool
[`Cls`](/docs/reference/modal.Cls)| A serverless class supporting
parametrization and lifecycle hooks

## Extended Function configuration

### Class parametrization

|
---|---
[`parameter`](/docs/reference/modal.parameter)| Used to define class
parameters, akin to a Dataclass field

### Lifecycle hooks

|
---|---
[`enter`](/docs/reference/modal.enter)| Decorator for a method that will be
executed during container startup
[`exit`](/docs/reference/modal.exit)| Decorator for a method that will be
executed during container shutdown
[`method`](/docs/reference/modal.method)| Decorator for exposing a method as
an invokable function

### Web integrations

|
---|---
[`fastapi_endpoint`](/docs/reference/modal.fastapi_endpoint)| Decorator for
exposing a simple FastAPI-based endpoint
[`asgi_app`](/docs/reference/modal.asgi_app)| Decorator for functions that
construct an ASGI web application
[`wsgi_app`](/docs/reference/modal.wsgi_app)| Decorator for functions that
construct a WSGI web application
[`web_server`](/docs/reference/modal.web_server)| Decorator for functions that
construct an HTTP web server

### Function semantics

|
---|---
[`batched`](/docs/reference/modal.batched)| Decorator that enables [dynamic
input batching](/docs/guide/dynamic-batching)
[`concurrent`](/docs/reference/modal.concurrent)| Decorator that enables
[input concurrency](/docs/guide/concurrent-inputs)

### Scheduling

|
---|---
[`Cron`](/docs/reference/modal.Cron)| A schedule that runs based on cron
syntax
[`Period`](/docs/reference/modal.Period)| A schedule that runs at a fixed
interval

### Exception handling

|
---|---
[`Retries`](/docs/reference/modal.Retries)| Function retry policy for input
failures

## Sandboxed execution

|
---|---
[`Sandbox`](/docs/reference/modal.Sandbox)| An interface for restricted code
execution
[`ContainerProcess`](/docs/reference/modal.container_process#modalcontainer_processcontainerprocess)|
An object representing a sandboxed process
[`FileIO`](/docs/reference/modal.file_io#modalfile_iofileio)| A handle for a
file in the Sandbox filesystem

## Container configuration

|
---|---
[`Image`](/docs/reference/modal.Image)| An API for specifying container images
[`Secret`](/docs/reference/modal.Secret)| A pointer to secrets that will be
exposed as environment variables

## Data primitives

### Persistent storage

|
---|---
[`Volume`](/docs/reference/modal.Volume)| Distributed storage supporting
highly performant parallel reads
[`CloudBucketMount`](/docs/reference/modal.CloudBucketMount)| Storage backed
by a third-party cloud bucket (S3, etc.)
[`NetworkFileSystem`](/docs/reference/modal.NetworkFileSystem)| Shared,
writeable cloud storage (superseded by `modal.Volume`)

### In-memory storage

|
---|---
[`Dict`](/docs/reference/modal.Dict)| A distributed key-value store
[`Queue`](/docs/reference/modal.Queue)| A distributed FIFO queue

## Networking

|
---|---
[`Proxy`](/docs/reference/modal.Proxy)| An object that provides a static
outbound IP address for containers
[`forward`](/docs/reference/modal.forward)| A context manager for publicly
exposing a port from a container

API ReferenceApplication constructionServerless executionExtended Function
configurationClass parametrizationLifecycle hooksWeb integrationsFunction
semanticsSchedulingException handlingSandboxed executionContainer
configurationData primitivesPersistent storageIn-memory storageNetworking
