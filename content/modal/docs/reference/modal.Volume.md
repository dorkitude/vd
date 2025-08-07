* * *

Copy page

# modal.Volume

    class Volume(modal.object.Object)

Copy

A writeable volume that can be used to share files between one or more Modal
functions.

The contents of a volume is exposed as a filesystem. You can use it to share
data between different functions, or to persist durable state across several
instances of the same function.

Unlike a networked filesystem, you need to explicitly reload the volume to see
changes made since it was mounted. Similarly, you need to explicitly commit
any changes you make to the volume for the changes to become visible outside
the current container.

Concurrent modification is supported, but concurrent modifications of the same
files should be avoided! Last write wins in case of concurrent modification of
the same file - any data the last writer didn’t have when committing changes
will be lost!

As a result, volumes are typically not a good fit for use cases where you need
to make concurrent modifications to the same file (nor is distributed file
locking supported).

Volumes can only be reloaded if there are no open files for the volume -
attempting to reload with open files will result in an error.

**Usage**

    import modal

    app = modal.App()
    volume = modal.Volume.from_name("my-persisted-volume", create_if_missing=True)

    @app.function(volumes={"/root/foo": volume})
    def f():
        with open("/root/foo/bar.txt", "w") as f:
            f.write("hello")
        volume.commit()  # Persist changes

    @app.function(volumes={"/root/foo": volume})
    def g():
        volume.reload()  # Fetch latest changes
        with open("/root/foo/bar.txt", "r") as f:
            print(f.read())

Copy

## hydrate

    def hydrate(self, client: Optional[_Client] = None) -> Self:

Copy

Synchronize the local object with its identity on the Modal server.

It is rarely necessary to call this method explicitly, as most operations will
lazily hydrate when needed. The main use case is when you need to access
object metadata, such as its ID.

_Added in v0.72.39_ : This method replaces the deprecated `.resolve()` method.

## read_only

    def read_only(self) -> "_Volume":

Copy

Configure Volume to mount as read-only.

**Example**

    import modal

    volume = modal.Volume.from_name("my-volume", create_if_missing=True)

    @app.function(volumes={"/mnt/items": volume.read_only()})
    def f():
        with open("/mnt/items/my-file.txt") as f:
            return f.read()

Copy

The Volume is mounted as a read-only volume in a function. Any file system
write operation into the mounted volume will result in an error.

Added in v1.0.5.

## name

    @property
    def name(self) -> Optional[str]:

Copy

## from_name

    @staticmethod
    def from_name(
        name: str,
        *,
        environment_name: Optional[str] = None,
        create_if_missing: bool = False,
        version: "typing.Optional[modal_proto.api_pb2.VolumeFsVersion.ValueType]" = None,
    ) -> "_Volume":

Copy

Reference a Volume by name, creating if necessary.

This is a lazy method that defers hydrating the local object with metadata
from Modal servers until the first time is is actually used.

    vol = modal.Volume.from_name("my-volume", create_if_missing=True)

    app = modal.App()

    # Volume refers to the same object, even across instances of `app`.
    @app.function(volumes={"/data": vol})
    def f():
        pass

Copy

## ephemeral

    @classmethod
    @contextmanager
    def ephemeral(
        cls: type["_Volume"],
        client: Optional[_Client] = None,
        environment_name: Optional[str] = None,
        version: "typing.Optional[modal_proto.api_pb2.VolumeFsVersion.ValueType]" = None,
        _heartbeat_sleep: float = EPHEMERAL_OBJECT_HEARTBEAT_SLEEP,
    ) -> AsyncGenerator["_Volume", None]:

Copy

Creates a new ephemeral volume within a context manager:

Usage:

    import modal
    with modal.Volume.ephemeral() as vol:
        assert vol.listdir("/") == []

Copy

    async with modal.Volume.ephemeral() as vol:
        assert await vol.listdir("/") == []

Copy

## info

    @live_method
    def info(self) -> VolumeInfo:

Copy

Return information about the Volume object.

## commit

    @live_method
    def commit(self):

Copy

Commit changes to the volume.

If successful, the changes made are now persisted in durable storage and
available to other containers accessing the volume.

## reload

    @live_method
    def reload(self):

Copy

Make latest committed state of volume available in the running container.

Any uncommitted changes to the volume, such as new or modified files, may
implicitly be committed when reloading.

Reloading will fail if there are open files for the volume.

## iterdir

    @live_method_gen
    def iterdir(self, path: str, *, recursive: bool = True) -> Iterator[FileEntry]:

Copy

Iterate over all files in a directory in the volume.

Passing a directory path lists all files in the directory. For a file path,
return only that file’s description. If `recursive` is set to True, list all
files and folders under the path recursively.

## listdir

    @live_method
    def listdir(self, path: str, *, recursive: bool = False) -> list[FileEntry]:

Copy

List all files under a path prefix in the modal.Volume.

Passing a directory path lists all files in the directory. For a file path,
return only that file’s description. If `recursive` is set to True, list all
files and folders under the path recursively.

## read_file

    @live_method_gen
    def read_file(self, path: str) -> Iterator[bytes]:

Copy

Read a file from the modal.Volume.

Note - this function is primarily intended to be used outside of a Modal App.
For more information on downloading files from a Modal Volume, see [the
guide](https://modal.com/docs/guide/volumes).

**Example:**

    vol = modal.Volume.from_name("my-modal-volume")
    data = b""
    for chunk in vol.read_file("1mb.csv"):
        data += chunk
    print(len(data))  # == 1024 * 1024

Copy

## remove_file

    @live_method
    def remove_file(self, path: str, recursive: bool = False) -> None:

Copy

Remove a file or directory from a volume.

## copy_files

    @live_method
    def copy_files(self, src_paths: Sequence[str], dst_path: str, recursive: bool = False) -> None:

Copy

Copy files within the volume from src_paths to dst_path. The semantics of the
copy operation follow those of the UNIX cp command.

The `src_paths` parameter is a list. If you want to copy a single file, you
should pass a list with a single element.

`src_paths` and `dst_path` should refer to the desired location _inside_ the
volume. You do not need to prepend the volume mount path.

**Usage**

    vol = modal.Volume.from_name("my-modal-volume")

    vol.copy_files(["bar/example.txt"], "bar2")  # Copy files to another directory
    vol.copy_files(["bar/example.txt"], "bar/example2.txt")  # Rename a file by copying

Copy

Note that if the volume is already mounted on the Modal function, you should
use normal filesystem operations like `os.rename()` and then `commit()` the
volume. The `copy_files()` method is useful when you don’t have the volume
mounted as a filesystem, e.g. when running a script on your local computer.

## batch_upload

    @live_method
    def batch_upload(self, force: bool = False) -> "_AbstractVolumeUploadContextManager":

Copy

Initiate a batched upload to a volume.

To allow overwriting existing files, set `force` to `True` (you cannot
overwrite existing directories with uploaded files regardless).

**Example:**

    vol = modal.Volume.from_name("my-modal-volume")

    with vol.batch_upload() as batch:
        batch.put_file("local-path.txt", "/remote-path.txt")
        batch.put_directory("/local/directory/", "/remote/directory")
        batch.put_file(io.BytesIO(b"some data"), "/foobar")

Copy

## delete

    @staticmethod
    def delete(name: str, client: Optional[_Client] = None, environment_name: Optional[str] = None):

Copy

## rename

    @staticmethod
    def rename(
        old_name: str,
        new_name: str,
        *,
        client: Optional[_Client] = None,
        environment_name: Optional[str] = None,
    ):

Copy

modal.Volumehydrateread_onlynamefrom_nameephemeralinfocommitreloaditerdirlistdirread_fileremove_filecopy_filesbatch_uploaddeleterename
