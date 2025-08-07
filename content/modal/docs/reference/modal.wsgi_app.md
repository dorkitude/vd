* * *

Copy page

# modal.wsgi_app

    def wsgi_app(
        _warn_parentheses_missing=None,
        *,
        label: Optional[str] = None,  # Label for created endpoint. Final subdomain will be <workspace>--<label>.modal.run.
        custom_domains: Optional[Iterable[str]] = None,  # Deploy this endpoint on a custom domain.
        requires_proxy_auth: bool = False,  # Require Modal-Key and Modal-Secret HTTP Headers on requests.
    ) -> Callable[[Union[_PartialFunction, NullaryFuncOrMethod]], _PartialFunction]:

Copy

Decorator for registering a WSGI app with a Modal function.

Web Server Gateway Interface (WSGI) is a standard for synchronous Python web
apps. It has been [succeeded by the ASGI
interface](https://asgi.readthedocs.io/en/latest/introduction.html#wsgi-
compatibility) which is compatible with ASGI and supports additional
functionality such as web sockets. Modal supports ASGI via
[`asgi_app`](https://modal.com/docs/reference/modal.asgi_app).

**Usage:**

    from typing import Callable

    @app.function()
    @modal.wsgi_app()
    def create_wsgi() -> Callable:
        ...

Copy

To learn how to use this decorator with popular web frameworks, see the [guide
on web endpoints](https://modal.com/docs/guide/webhooks).

modal.wsgi_app
