* * *

Copy page

# `modal nfs`

Read and edit `modal.NetworkFileSystem` file systems.

**Usage** :

    modal nfs [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `list`: List the names of all network file systems.
  * `create`: Create a named network file system.
  * `ls`: List files and directories in a network file system.
  * `put`: Upload a file or directory to a network file system.
  * `get`: Download a file from a network file system.
  * `rm`: Delete a file or directory from a network file system.
  * `delete`: Delete a named, persistent modal.NetworkFileSystem.

## `modal nfs list`

List the names of all network file systems.

**Usage** :

    modal nfs list [OPTIONS]

Copy

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--json / --no-json`: [default: no-json]
  * `--help`: Show this message and exit.

## `modal nfs create`

Create a named network file system.

**Usage** :

    modal nfs create [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: [required]

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal nfs ls`

List files and directories in a network file system.

**Usage** :

    modal nfs ls [OPTIONS] VOLUME_NAME [PATH]

Copy

**Arguments** :

  * `VOLUME_NAME`: [required]
  * `[PATH]`: [default: /]

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal nfs put`

Upload a file or directory to a network file system.

Remote parent directories will be created as needed.

Ending the REMOTE_PATH with a forward slash (/), it’s assumed to be a
directory and the file will be uploaded with its current name under that
directory.

**Usage** :

    modal nfs put [OPTIONS] VOLUME_NAME LOCAL_PATH [REMOTE_PATH]

Copy

**Arguments** :

  * `VOLUME_NAME`: [required]
  * `LOCAL_PATH`: [required]
  * `[REMOTE_PATH]`: [default: /]

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal nfs get`

Download a file from a network file system.

Specifying a glob pattern (using any `*` or `**` patterns) as the
`remote_path` will download all matching files, preserving their directory
structure.

For example, to download an entire network file system into `dump_volume`:

    modal nfs get <volume-name> "**" dump_volume

Copy

Use ”-” as LOCAL_DESTINATION to write file contents to standard output.

**Usage** :

    modal nfs get [OPTIONS] VOLUME_NAME REMOTE_PATH [LOCAL_DESTINATION]

Copy

**Arguments** :

  * `VOLUME_NAME`: [required]
  * `REMOTE_PATH`: [required]
  * `[LOCAL_DESTINATION]`: [default: .]

**Options** :

  * `--force / --no-force`: [default: no-force]
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal nfs rm`

Delete a file or directory from a network file system.

**Usage** :

    modal nfs rm [OPTIONS] VOLUME_NAME REMOTE_PATH

Copy

**Arguments** :

  * `VOLUME_NAME`: [required]
  * `REMOTE_PATH`: [required]

**Options** :

  * `-r, --recursive`: Delete directory recursively
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal nfs delete`

Delete a named, persistent modal.NetworkFileSystem.

**Usage** :

    modal nfs delete [OPTIONS] NFS_NAME

Copy

**Arguments** :

  * `NFS_NAME`: Name of the modal.NetworkFileSystem to be deleted. Case sensitive [required]

**Options** :

  * `-y, --yes`: Run without pausing for confirmation.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

modal nfsmodal nfs listmodal nfs createmodal nfs lsmodal nfs putmodal nfs
getmodal nfs rmmodal nfs delete
