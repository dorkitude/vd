* * *

Copy page

# Introduction

Modal is a cloud function platform that lets you:

  * Run any code remotely within seconds.
  * Define [container environments](/docs/guide/images) in code (or use one of our pre-built backends).
  * [Scale out horizontally](/docs/guide/scale) to thousands of containers.
  * Attach [GPUs](/docs/guide/gpu) with a single line of code.
  * Serve your functions as [web endpoints](/docs/guide/webhooks).
  * Deploy and monitor [persistent scheduled jobs](/docs/guide/cron).
  * Store data in distributed versions of [dictionaries](/docs/guide/dicts) and [queues](/docs/guide/queues).

You get [full serverless execution and pricing](/pricing), because we host
everything and charge per second of usage. Notably, there’s zero configuration
in Modal - everything is code. Take a breath of fresh air and feel how good it
tastes with no YAML in it.

## Getting started

The nicest thing about all of this is that **you don’t have to set up any
infrastructure.** Just:

  1. Create an account at [modal.com](https://modal.com)
  2. Run `pip install modal` to install the `modal` Python package
  3. Run `modal setup` to authenticate (if this doesn’t work, try `python -m modal setup`)

…and you can start running jobs right away. Check out some of our simple
getting started examples:

  * [Hello, world!](/docs/examples/hello_world)
  * [A simple web scraper](/docs/examples/webscraper)

You can also learn Modal interactively without installing anything through our
[code playground](/playground).

## How does it work?

Modal takes your code, puts it in a container, and executes it in the cloud.

Where does it run? Modal runs it in its own cloud environment. The benefit is
that we solve all the hard infrastructure problems for you, so you don’t have
to do anything. You don’t need to mess with Kubernetes, Docker or even an AWS
account.

Modal is currently Python-only, but we may support other languages in the
future.

IntroductionGetting startedHow does it work?

See it in action

[Hello, world!](/docs/examples/hello_world)

[A simple web scraper](/docs/examples/webscraper)
