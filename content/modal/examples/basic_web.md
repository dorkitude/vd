# Hello world wide web!

* * *

[ View on GitHub](https://github.com/modal-labs/modal-
examples/blob/main/07_web_endpoints/basic_web.py)

Copy page

# Hello world wide web!

Modal makes it easy to turn your Python functions into serverless web
services: access them via a browser or call them from any client that speaks
HTTP, all without having to worry about setting up servers or managing
infrastructure.

This tutorial shows the path with the shortest [“time to
200”](https://shkspr.mobi/blog/2021/05/whats-your-apis-time-to-200/):
[`modal.fastapi_endpoint`](https://modal.com/docs/reference/modal.fastapi_endpoint).

On Modal, web endpoints have all the superpowers of Modal Functions: they can
be [accelerated with GPUs](https://modal.com/docs/guide/gpu), they can access
[Secrets](https://modal.com/docs/guide/secrets) or
[Volumes](https://modal.com/docs/guide/volumes), and they [automatically
scale](https://modal.com/docs/guide/cold-start) to handle more traffic.

Under the hood, we use the [FastAPI library](https://fastapi.tiangolo.com/),
which has [high-quality
documentation](https://fastapi.tiangolo.com/tutorial/), linked throughout this
tutorial.

## Turn a Modal Function into an API endpoint with a single decorator

Modal Functions are already accessible remotely — when you add the
`@app.function` decorator to a Python function and run `modal deploy`, you
make it possible for your [other Python functions to call
it](https://modal.com/docs/guide/trigger-deployed-functions).

That’s great, but it’s not much help if you want to share what you’ve written
with someone running code in a different language — or not running code at
all!

And that’s where most of the power of the Internet comes from: sharing
information and functionality across different computer systems.

So we provide the `fastapi_endpoint` decorator to wrap your Modal Functions in
the lingua franca of the web: HTTP. Here’s what that looks like:

    import modal

    image = modal.Image.debian_slim().pip_install("fastapi[standard]")
    app = modal.App(name="example-basic-web", image=image)

    @app.function()
    @modal.fastapi_endpoint(
        docs=True  # adds interactive documentation in the browser
    )
    def hello():
        return "Hello world!"

Copy

You can turn this function into a web endpoint by running `modal serve
basic_web.py`. In the output, you should see a URL that ends with `hello-
dev.modal.run`. If you navigate to this URL, you should see the `"Hello
world!"` message appear in your browser.

You can also find interactive documentation, powered by OpenAPI and Swagger,
if you add `/docs` to the end of the URL. From this documentation, you can
interact with your endpoint, sending HTTP requests and receiving HTTP
responses. For more details, see the [FastAPI
documentation](https://fastapi.tiangolo.com/features/#automatic-docs).

By running the endpoint with `modal serve`, you created a temporary endpoint
that will disappear if you interrupt your terminal. These temporary endpoints
are great for debugging — when you save a change to any of your dependent
files, the endpoint will redeploy. Try changing the message to something else,
hitting save, and then hitting refresh in your browser or re-sending the
request from `/docs` or the command line. You should see the new message,
along with logs in your terminal showing the redeploy and the request.

When you’re ready to deploy this endpoint permanently, run `modal deploy
basic_web.py`. Now, your function will be available even when you’ve closed
your terminal or turned off your computer.

## Send data to a web endpoint

The web endpoint above was a bit silly: it always returns the same message.

Most endpoints need an input to be useful. There are two ways to send data to
a web endpoint:

  * in the URL as a query parameter
  * in the body of the request as JSON

### Sending data in query parameters

By default, your function’s arguments are treated as query parameters: they
are extracted from the end of the URL, where they should be added in the form
`?arg1=foo&arg2=bar`.

From the Python side, there’s hardly anything to do:

    @app.function()
    @modal.fastapi_endpoint(docs=True)
    def greet(user: str) -> str:
        return f"Hello {user}!"

Copy

If you are already running `modal serve basic_web.py`, this endpoint will be
available at a URL, printed in your terminal, that ends with `greet-
dev.modal.run`.

We provide Python type-hints to get type information in the docs and
[automatic validation](https://fastapi.tiangolo.com/tutorial/query-params-str-
validations/). For example, if you navigate directly to the URL for `greet`,
you will get a detailed error message indicating that the `user` parameter is
missing. Navigate instead to `/docs` to see how to invoke the endpoint
properly.

You can read more about query parameters in the [FastAPI
documentation](https://fastapi.tiangolo.com/tutorial/query-params/).

### Sending data in the request body

For larger and more complex data, it is generally preferrable to send data in
the body of the HTTP request. This body is formatted as
[JSON](https://developer.mozilla.org/en-
US/docs/Learn/JavaScript/Objects/JSON), the most common data interchange
format on the web.

To set up an endpoint that accepts JSON data, add an argument with a `dict`
type-hint to your function. This argument will be populated with the data sent
in the request body.

    @app.function()
    @modal.fastapi_endpoint(method="POST", docs=True)
    def goodbye(data: dict) -> str:
        name = data.get("name") or "world"
        return f"Goodbye {name}!"

Copy

Note that we gave a value of `"POST"` for the `method` argument here. This
argument defines the HTTP request method that the endpoint will respond to,
and it defaults to `"GET"`. If you head to the URL for the `goodbye` endpoint
in your browser, you will get a 405 Method Not Allowed error, because browsers
only send GET requests by default. While this is technically a separate
concern from query parameters versus request bodies and you can define an
endpoint that accepts GET requests and uses data from the body, it is
[considered bad form](https://stackoverflow.com/a/983458).

Navigate to `/docs` for more on how to invoke the endpoint properly. You will
need to send a POST request with a JSON body containing a `name` key. To get
the same typing and validation benefits as with query parameters, use a
[Pydantic model](https://fastapi.tiangolo.com/tutorial/body/) for this
argument.

You can read more about request bodies in the [FastAPI
documentation](https://fastapi.tiangolo.com/tutorial/body/).

## Handle expensive startup with `modal.Cls`

Sometimes your endpoint needs to do something before it can handle its first
request, like get a value from a database or set the value of a variable. If
that step is expensive, like [loading a large ML
model](https://modal.com/docs/guide/model-weights), it’d be a shame to have to
do it every time a request comes in!

Web endpoints can be methods on a
[`modal.Cls`](https://modal.com/docs/guide/lifecycle-functions#container-
lifecycle-functions-and-parameters), which allows you to manage the
container’s lifecycle independently from processing individual requests.

This example will only set the `start_time` instance variable once, on
container startup.

    @app.cls()
    class WebApp:
        @modal.enter()
        def startup(self):
            from datetime import datetime, timezone

            print("🏁 Starting up!")
            self.start_time = datetime.now(timezone.utc)

        @modal.fastapi_endpoint(docs=True)
        def web(self):
            from datetime import datetime, timezone

            current_time = datetime.now(timezone.utc)
            return {"start_time": self.start_time, "current_time": current_time}

Copy

## Protect web endpoints with proxy authentication

Sharing your Python functions on the web is great, but it’s not always a good
idea to make those functions available to just anyone.

For example, you might have a function like the one below that is more
expensive to run than to call (and so might be abused by your enemies) or
reveals information that you would rather keep secret.

To protect your Modal web endpoints so that they can’t be triggered except by
members of your [Modal workspace](https://modal.com/docs/guide/workspaces),
add the `requires_proxy_auth=True` flag to the `fastapi_endpoint` decorator.

    @app.function(gpu="h100")
    @modal.fastapi_endpoint(requires_proxy_auth=True, docs=False)
    def expensive_secret():
        return "I didn't care for 'The Godfather'. It insists upon itself."

Copy

The `expensive-secret` endpoint URL will still be printed to the output when
you `modal serve` or `modal deploy`, along with a ”🔑” emoji indicating that it
is secured with proxy authentication. If you head to that URL via the browser,
you will get a [`401 Unauthorized`](https://developer.mozilla.org/en-
US/docs/Web/HTTP/Status/401) error code in response. You should also check the
dashboard page for this app (at the URL printed at the very top of the `modal`
command output) so you can see that no containers were spun up to handle the
request — this authorization is handled entirely inside Modal’s
infrastructure.

You can trigger the web endpoint by [creating a Proxy Auth
Token](https://modal.com/settings/proxy-auth-tokens) and then including the
token ID and secret in the `Modal-Key` and `Modal-Secret` headers.

From the command line, that might look like

    export TOKEN_ID=wk-1234abcd
    export TOKEN_SECRET=ws-1234abcd
    curl -H "Modal-Key: $TOKEN_ID" \
         -H "Modal-Secret: $TOKEN_SECRET" \
         https://your-workspace-name--expensive-secret.modal.run

Copy

For more details, see the [guide to proxy
authentication](https://modal.com/docs/guide/webhook-proxy-auth).

## What next?

Modal’s `fastapi_endpoint` decorator is opinionated and designed for
relatively simple web applications — one or a few independent Python functions
that you want to expose to the web.

Three additional decorators allow you to serve more complex web applications
with greater control:

  * [`asgi_app`](https://modal.com/docs/guide/webhooks#asgi) to serve applications compliant with the ASGI standard, like [FastAPI](https://fastapi.tiangolo.com/)
  * [`wsgi_app`](https://modal.com/docs/guide/webhooks#wsgi) to serve applications compliant with the WSGI standard, like [Flask](https://flask.palletsprojects.com/)
  * [`web_server`](https://modal.com/docs/guide/webhooks#non-asgi-web-servers) to serve any application that listens on a port

Hello world wide web!Turn a Modal Function into an API endpoint with a single
decoratorSend data to a web endpointSending data in query parametersSending
data in the request bodyHandle expensive startup with modal.ClsProtect web
endpoints with proxy authenticationWhat next?

## Try this on Modal!

You can run this example on Modal in 60 seconds.

[Create account to run](/signup)

After creating a free account, install the Modal Python package, and create an
API token.

$

    pip install modal

$

    modal setup

Copy

Clone the [modal-examples](https://github.com/modal-labs/modal-examples)
repository and run:

$

    git clone https://github.com/modal-labs/modal-examples

$

    cd modal-examples

$

    modal serve 07_web_endpoints/basic_web.py

Copy
