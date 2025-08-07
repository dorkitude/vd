* * *

Copy page

# modal.file_io

## modal.file_io.FileIO

    class FileIO(typing.Generic)

Copy

[Alpha] FileIO handle, used in the Sandbox filesystem API.

The API is designed to mimic Python’s io.FileIO.

Currently this API is in Alpha and is subject to change. File I/O operations
may be limited in size to 100 MiB, and the throughput of requests is
restricted in the current implementation. For our recommendations on large
file transfers see the Sandbox [filesystem access
guide](https://modal.com/docs/guide/sandbox-files).

**Usage**

    import modal

    app = modal.App.lookup("my-app", create_if_missing=True)

    sb = modal.Sandbox.create(app=app)
    f = sb.open("/tmp/foo.txt", "w")
    f.write("hello")
    f.close()

Copy

    def __init__(self, client: _Client, task_id: str) -> None:

Copy

### create

    @classmethod
    def create(
        cls, path: str, mode: Union["_typeshed.OpenTextMode", "_typeshed.OpenBinaryMode"], client: _Client, task_id: str
    ) -> "_FileIO":

Copy

Create a new FileIO handle.

### read

    def read(self, n: Optional[int] = None) -> T:

Copy

Read n bytes from the current position, or the entire remaining file if n is
None.

### readline

    def readline(self) -> T:

Copy

Read a single line from the current position.

### readlines

    def readlines(self) -> Sequence[T]:

Copy

Read all lines from the current position.

### write

    def write(self, data: Union[bytes, str]) -> None:

Copy

Write data to the current position.

Writes may not appear until the entire buffer is flushed, which can be done
manually with `flush()` or automatically when the file is closed.

### flush

    def flush(self) -> None:

Copy

Flush the buffer to disk.

### seek

    def seek(self, offset: int, whence: int = 0) -> None:

Copy

Move to a new position in the file.

`whence` defaults to 0 (absolute file positioning); other values are 1
(relative to the current position) and 2 (relative to the file’s end).

### ls

    @classmethod
    def ls(cls, path: str, client: _Client, task_id: str) -> list[str]:

Copy

List the contents of the provided directory.

### mkdir

    @classmethod
    def mkdir(cls, path: str, client: _Client, task_id: str, parents: bool = False) -> None:

Copy

Create a new directory.

### rm

    @classmethod
    def rm(cls, path: str, client: _Client, task_id: str, recursive: bool = False) -> None:

Copy

Remove a file or directory in the Sandbox.

### watch

    @classmethod
    def watch(
        cls,
        path: str,
        client: _Client,
        task_id: str,
        filter: Optional[list[FileWatchEventType]] = None,
        recursive: bool = False,
        timeout: Optional[int] = None,
    ) -> Iterator[FileWatchEvent]:

Copy

### close

    def close(self) -> None:

Copy

Flush the buffer and close the file.

## modal.file_io.FileWatchEvent

    class FileWatchEvent(object)

Copy

FileWatchEvent(paths: list[str], type: modal.file_io.FileWatchEventType)

    def __init__(self, paths: list[str], type: modal.file_io.FileWatchEventType) -> None

Copy

## modal.file_io.FileWatchEventType

    class FileWatchEventType(enum.Enum)

Copy

An enumeration.

The possible values are:

  * `Unknown`
  * `Access`
  * `Create`
  * `Modify`
  * `Remove`

## modal.file_io.delete_bytes

    async def delete_bytes(file: "_FileIO", start: Optional[int] = None, end: Optional[int] = None) -> None:

Copy

Delete a range of bytes from the file.

`start` and `end` are byte offsets. `start` is inclusive, `end` is exclusive.
If either is None, the start or end of the file is used, respectively.

## modal.file_io.replace_bytes

    async def replace_bytes(file: "_FileIO", data: bytes, start: Optional[int] = None, end: Optional[int] = None) -> None:

Copy

Replace a range of bytes in the file with new data. The length of the data
does not have to be the same as the length of the range being replaced.

`start` and `end` are byte offsets. `start` is inclusive, `end` is exclusive.
If either is None, the start or end of the file is used, respectively.

modal.file_iomodal.file_io.FileIOcreatereadreadlinereadlineswriteflushseeklsmkdirrmwatchclosemodal.file_io.FileWatchEventmodal.file_io.FileWatchEventTypemodal.file_io.delete_bytesmodal.file_io.replace_bytes
