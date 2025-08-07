* * *

Copy page

# modal.web_server

    def web_server(
        port: int,
        *,
        startup_timeout: float = 5.0,  # Maximum number of seconds to wait for the web server to start.
        label: Optional[str] = None,  # Label for created endpoint. Final subdomain will be <workspace>--<label>.modal.run.
        custom_domains: Optional[Iterable[str]] = None,  # Deploy this endpoint on a custom domain.
        requires_proxy_auth: bool = False,  # Require Modal-Key and Modal-Secret HTTP Headers on requests.
    ) -> Callable[[Union[_PartialFunction, NullaryFuncOrMethod]], _PartialFunction]:

Copy

Decorator that registers an HTTP web server inside the container.

This is similar to `@asgi_app` and `@wsgi_app`, but it allows you to expose a
full HTTP server listening on a container port. This is useful for servers
written in other languages like Rust, as well as integrating with non-ASGI
frameworks like aiohttp and Tornado.

**Usage:**

    import subprocess

    @app.function()
    @modal.web_server(8000)
    def my_file_server():
        subprocess.Popen("python -m http.server -d / 8000", shell=True)

Copy

The above example starts a simple file server, displaying the contents of the
root directory. Here, requests to the web endpoint will go to external port
8000 on the container. The `http.server` module is included with Python, but
you could run anything here.

Internally, the web server is transparently converted into a web endpoint by
Modal, so it has the same serverless autoscaling behavior as other web
endpoints.

For more info, see the [guide on web
endpoints](https://modal.com/docs/guide/webhooks).

modal.web_server
