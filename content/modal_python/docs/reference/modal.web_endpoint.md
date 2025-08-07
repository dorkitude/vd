* * *

Copy page

# modal.web_endpoint

    def web_endpoint(
        _warn_parentheses_missing=None,
        *,
        method: str = "GET",  # REST method for the created endpoint.
        label: Optional[str] = None,  # Label for created endpoint. Final subdomain will be <workspace>--<label>.modal.run.
        docs: bool = False,  # Whether to enable interactive documentation for this endpoint at /docs.
        custom_domains: Optional[
            Iterable[str]
        ] = None,  # Create an endpoint using a custom domain fully-qualified domain name (FQDN).
        requires_proxy_auth: bool = False,  # Require Modal-Key and Modal-Secret HTTP Headers on requests.
    ) -> Callable[
        [Union[_PartialFunction[P, ReturnType, ReturnType], Callable[P, ReturnType]]],
        _PartialFunction[P, ReturnType, ReturnType],
    ]:

Copy

Register a basic web endpoint with this application.

DEPRECATED: This decorator has been renamed to `@modal.fastapi_endpoint`.

This is the simple way to create a web endpoint on Modal. The function behaves
as a [FastAPI](https://fastapi.tiangolo.com/) handler and should return a
response object to the caller.

Endpoints created with `@modal.web_endpoint` are meant to be simple, single
request handlers and automatically have
[CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) enabled. For
more flexibility, use `@modal.asgi_app`.

To learn how to use Modal with popular web frameworks, see the [guide on web
endpoints](https://modal.com/docs/guide/webhooks).

modal.web_endpoint
