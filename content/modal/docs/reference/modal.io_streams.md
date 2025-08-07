* * *

Copy page

# modal.io_streams

## modal.io_streams.StreamReader

    class StreamReader(typing.Generic)

Copy

Retrieve logs from a stream (`stdout` or `stderr`).

As an asynchronous iterable, the object supports the `for` and `async for`
statements. Just loop over the object to read in chunks.

**Usage**

    from modal import Sandbox

    sandbox = Sandbox.create(
        "bash",
        "-c",
        "for i in $(seq 1 10); do echo foo; sleep 0.1; done",
        app=running_app,
    )
    for message in sandbox.stdout:
        print(f"Message: {message}")

Copy

### file_descriptor

    @property
    def file_descriptor(self) -> int:

Copy

Possible values are `1` for stdout and `2` for stderr.

### read

    def read(self) -> T:

Copy

Fetch the entire contents of the stream until EOF.

**Usage**

    from modal import Sandbox

    sandbox = Sandbox.create("echo", "hello", app=running_app)
    sandbox.wait()

    print(sandbox.stdout.read())

Copy

## modal.io_streams.StreamWriter

    class StreamWriter(object)

Copy

Provides an interface to buffer and write logs to a sandbox or container
process stream (`stdin`).

### write

    def write(self, data: Union[bytes, bytearray, memoryview, str]) -> None:

Copy

Write data to the stream but does not send it immediately.

This is non-blocking and queues the data to an internal buffer. Must be used
along with the `drain()` method, which flushes the buffer.

**Usage**

    from modal import Sandbox

    sandbox = Sandbox.create(
        "bash",
        "-c",
        "while read line; do echo $line; done",
        app=running_app,
    )
    sandbox.stdin.write(b"foo\n")
    sandbox.stdin.write(b"bar\n")
    sandbox.stdin.write_eof()

    sandbox.stdin.drain()
    sandbox.wait()

Copy

### write_eof

    def write_eof(self) -> None:

Copy

Close the write end of the stream after the buffered data is drained.

If the process was blocked on input, it will become unblocked after
`write_eof()`. This method needs to be used along with the `drain()` method,
which flushes the EOF to the process.

### drain

    def drain(self) -> None:

Copy

Flush the write buffer and send data to the running process.

This is a flow control method that blocks until data is sent. It returns when
it is appropriate to continue writing data to the stream.

**Usage**

    writer.write(data)
    writer.drain()

Copy

Async usage:

    writer.write(data)  # not a blocking operation
    await writer.drain.aio()

Copy

modal.io_streamsmodal.io_streams.StreamReaderfile_descriptorreadmodal.io_streams.StreamWriterwritewrite_eofdrain
