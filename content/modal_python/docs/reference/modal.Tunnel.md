* * *

Copy page

# modal.Tunnel

    class Tunnel(object)

Copy

A port forwarded from within a running Modal container. Created by
`modal.forward()`.

**Important:** This is an experimental API which may change in the future.

    def __init__(self, host: str, port: int, unencrypted_host: str, unencrypted_port: int) -> None

Copy

## url

    @property
    def url(self) -> str:

Copy

Get the public HTTPS URL of the forwarded port.

## tls_socket

    @property
    def tls_socket(self) -> tuple[str, int]:

Copy

Get the public TLS socket as a (host, port) tuple.

## tcp_socket

    @property
    def tcp_socket(self) -> tuple[str, int]:

Copy

Get the public TCP socket as a (host, port) tuple.

modal.Tunnelurltls_sockettcp_socket
