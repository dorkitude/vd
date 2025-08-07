* * *

Copy page

# modal.forward

    @contextmanager
    def forward(port: int, *, unencrypted: bool = False, client: Optional[_Client] = None) -> Iterator[Tunnel]:

Copy

Expose a port publicly from inside a running Modal container, with TLS.

If `unencrypted` is set, this also exposes the TCP socket without encryption
on a random port number. This can be used to SSH into a container (see example
below). Note that it is on the public Internet, so make sure you are using a
secure protocol over TCP.

**Important:** This is an experimental API which may change in the future.

**Usage:**

    import modal
    from flask import Flask

    app = modal.App(image=modal.Image.debian_slim().pip_install("Flask"))
    flask_app = Flask(__name__)

    @flask_app.route("/")
    def hello_world():
        return "Hello, World!"

    @app.function()
    def run_app():
        # Start a web server inside the container at port 8000. `modal.forward(8000)` lets us
        # expose that port to the world at a random HTTPS URL.
        with modal.forward(8000) as tunnel:
            print("Server listening at", tunnel.url)
            flask_app.run("0.0.0.0", 8000)

        # When the context manager exits, the port is no longer exposed.

Copy

**Raw TCP usage:**

    import socket
    import threading

    import modal

    def run_echo_server(port: int):
        """Run a TCP echo server listening on the given port."""
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.bind(("0.0.0.0", port))
        sock.listen(1)

        while True:
            conn, addr = sock.accept()
            print("Connection from:", addr)

            # Start a new thread to handle the connection
            def handle(conn):
                with conn:
                    while True:
                        data = conn.recv(1024)
                        if not data:
                            break
                        conn.sendall(data)

            threading.Thread(target=handle, args=(conn,)).start()

    app = modal.App()

    @app.function()
    def tcp_tunnel():
        # This exposes port 8000 to public Internet traffic over TCP.
        with modal.forward(8000, unencrypted=True) as tunnel:
            # You can connect to this TCP socket from outside the container, for example, using `nc`:
            #  nc <HOST> <PORT>
            print("TCP tunnel listening at:", tunnel.tcp_socket)
            run_echo_server(8000)

Copy

**SSH example:** This assumes you have a rsa keypair in `~/.ssh/id_rsa{.pub}`,
this is a bare-bones example letting you SSH into a Modal container.

    import subprocess
    import time

    import modal

    app = modal.App()
    image = (
        modal.Image.debian_slim()
        .apt_install("openssh-server")
        .run_commands("mkdir /run/sshd")
        .add_local_file("~/.ssh/id_rsa.pub", "/root/.ssh/authorized_keys", copy=True)
    )

    @app.function(image=image, timeout=3600)
    def some_function():
        subprocess.Popen(["/usr/sbin/sshd", "-D", "-e"])
        with modal.forward(port=22, unencrypted=True) as tunnel:
            hostname, port = tunnel.tcp_socket
            connection_cmd = f'ssh -p {port} root@{hostname}'
            print(f"ssh into container using: {connection_cmd}")
            time.sleep(3600)  # keep alive for 1 hour or until killed

Copy

If you intend to use this more generally, a suggestion is to put the
subprocess and port forwarding code in an `@enter` lifecycle method of an
@app.cls, to only make a single ssh server and port for each container (and
not one for each input to the function).

modal.forward
