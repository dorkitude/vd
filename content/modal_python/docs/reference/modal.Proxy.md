* * *

Copy page

# modal.Proxy

    class Proxy(modal.object.Object)

Copy

Proxy objects give your Modal containers a static outbound IP address.

This can be used for connecting to a remote address with network whitelist,
for example a database. See [the guide](https://modal.com/docs/guide/proxy-
ips) for more information.

## hydrate

    def hydrate(self, client: Optional[_Client] = None) -> Self:

Copy

Synchronize the local object with its identity on the Modal server.

It is rarely necessary to call this method explicitly, as most operations will
lazily hydrate when needed. The main use case is when you need to access
object metadata, such as its ID.

_Added in v0.72.39_ : This method replaces the deprecated `.resolve()` method.

## from_name

    @staticmethod
    def from_name(
        name: str,
        *,
        environment_name: Optional[str] = None,
    ) -> "_Proxy":

Copy

Reference a Proxy by its name.

In contrast to most other Modal objects, new Proxy objects must be provisioned
via the Dashboard and cannot be created on the fly from code.

modal.Proxyhydratefrom_name
