* * *

Copy page

# Networking and security

Sandboxes are built to be secure-by-default, meaning that a default Sandbox
has no ability to accept incoming network connections or access your Modal
resources.

## Networking

Since Sandboxes may run untrusted code, they have options to restrict their
network access. To block all network access, set `block_network=True` on
[`Sandbox.create`](/docs/reference/modal.Sandbox#create).

For more fine-grained networking control, a Sandbox’s outbound network access
can be restricted using the `cidr_allowlist` parameter. This parameter takes a
list of CIDR ranges that the Sandbox is allowed to access, blocking all other
outbound traffic.

### Forwarding ports

Sandboxes can also expose TCP ports to the internet. This is useful if, for
example, you want to connect to a web server running inside a Sandbox.

Use the `encrypted_ports` and `unencrypted_ports` parameters of
`Sandbox.create` to specify which ports to forward. You can then access the
public URL of a tunnel using the
[`Sandbox.tunnels`](/docs/reference/modal.Sandbox#tunnels) method:

    import requests
    import time

    sb = modal.Sandbox.create(
        "python",
        "-m",
        "http.server",
        "12345",
        encrypted_ports=[12345],
        app=my_app,
    )

    tunnel = sb.tunnels()[12345]

    time.sleep(1)  # Wait for server to start.

    print(f"Connecting to {tunnel.url}...")
    print(requests.get(tunnel.url, timeout=5).text)

Copy

It is also possible to create an encrypted port that uses `HTTP/2` rather than
`HTTP/1.1` with the `h2_ports` option. This will return a URL that you can
make H2 (HTTP/2 + TLS) requests to. If you want to run an `HTTP/2` server
inside a sandbox, this feature may be useful. Here is an example:

    import time

    port = 4359
    sb = modal.Sandbox.create(
        app=my_app,
        image=my_image,
        h2_ports = [port],
    )
    p = sb.exec("python", "my_http2_server.py")

    tunnel = sb.tunnels()[port]
    time.sleep(1)
    print(f"Tunnel URL: {tunnel.url}")

Copy

For more details on how tunnels work, see the [tunnels
guide](/docs/guide/tunnels).

## Security model

In a typical Modal Function, the Function code can call other Modal APIs
allowing it to spawn containers, create and destroy Volumes, read from Dicts
and Queues, etc. Sandboxes, by contrast, are isolated from the main Modal
workspace. They have no API access, meaning the blast radius of any malicious
code is limited to the Sandbox environment.

Sandboxes are built on top of [gVisor](https://gvisor.dev/), a container
runtime by Google that provides strong isolation properties. gVisor has custom
logic to prevent Sandboxes from making malicious system calls, giving you
stronger isolation than standard
[runc](https://github.com/opencontainers/runc) containers.

Networking and securityNetworkingForwarding portsSecurity model

See it in action

[Running a Jupyter notebook](/docs/examples/jupyter_sandbox)

[Safe code execution](/docs/examples/safe_code_execution)
