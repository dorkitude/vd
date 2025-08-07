* * *

Copy page

# modal.asgi_app

    def asgi_app(
        _warn_parentheses_missing=None,
        *,
        label: Optional[str] = None,  # Label for created endpoint. Final subdomain will be <workspace>--<label>.modal.run.
        custom_domains: Optional[Iterable[str]] = None,  # Deploy this endpoint on a custom domain.
        requires_proxy_auth: bool = False,  # Require Modal-Key and Modal-Secret HTTP Headers on requests.
    ) -> Callable[[Union[_PartialFunction, NullaryFuncOrMethod]], _PartialFunction]:

Copy

Decorator for registering an ASGI app with a Modal function.

Asynchronous Server Gateway Interface (ASGI) is a standard for Python
synchronous and asynchronous apps, supported by all popular Python web
libraries. This is an advanced decorator that gives full flexibility in
defining one or more web endpoints on Modal.

**Usage:**

    from typing import Callable

    @app.function()
    @modal.asgi_app()
    def create_asgi() -> Callable:
        ...

Copy

To learn how to use Modal with popular web frameworks, see the [guide on web
endpoints](https://modal.com/docs/guide/webhooks).

modal.asgi_app
