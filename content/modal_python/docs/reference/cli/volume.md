* * *

Copy page

# `modal volume`

Read and edit `modal.Volume` volumes.

Note: users of `modal.NetworkFileSystem` should use the `modal nfs` command
instead.

**Usage** :

    modal volume [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `create`: Create a named, persistent modal.Volume.
  * `get`: Download files from a modal.Volume object.
  * `list`: List the details of all modal.Volume volumes in an Environment.
  * `ls`: List files and directories in a modal.Volume volume.
  * `put`: Upload a file or directory to a modal.Volume.
  * `rm`: Delete a file or directory from a modal.Volume.
  * `cp`: Copy within a modal.Volume.
  * `delete`: Delete a named, persistent modal.Volume.
  * `rename`: Rename a modal.Volume.

## `modal volume create`

Create a named, persistent modal.Volume.

**Usage** :

    modal volume create [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: [required]

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--version INTEGER`: VolumeFS version. (Experimental)
  * `--help`: Show this message and exit.

## `modal volume get`

Download files from a modal.Volume object.

If a folder is passed for REMOTE_PATH, the contents of the folder will be
downloaded recursively, including all subdirectories.

**Example**

    modal volume get <volume_name> logs/april-12-1.txt
    modal volume get <volume_name> / volume_data_dump

Copy

Use ”-” as LOCAL_DESTINATION to write file contents to standard output.

**Usage** :

    modal volume get [OPTIONS] VOLUME_NAME REMOTE_PATH [LOCAL_DESTINATION]

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

## `modal volume list`

List the details of all modal.Volume volumes in an Environment.

**Usage** :

    modal volume list [OPTIONS]

Copy

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--json / --no-json`: [default: no-json]
  * `--help`: Show this message and exit.

## `modal volume ls`

List files and directories in a modal.Volume volume.

**Usage** :

    modal volume ls [OPTIONS] VOLUME_NAME [PATH]

Copy

**Arguments** :

  * `VOLUME_NAME`: [required]
  * `[PATH]`: [default: /]

**Options** :

  * `--json / --no-json`: [default: no-json]
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal volume put`

Upload a file or directory to a modal.Volume.

Remote parent directories will be created as needed.

Ending the REMOTE_PATH with a forward slash (/), it’s assumed to be a
directory and the file will be uploaded with its current name under that
directory.

**Usage** :

    modal volume put [OPTIONS] VOLUME_NAME LOCAL_PATH [REMOTE_PATH]

Copy

**Arguments** :

  * `VOLUME_NAME`: [required]
  * `LOCAL_PATH`: [required]
  * `[REMOTE_PATH]`: [default: /]

**Options** :

  * `-f, --force`: Overwrite existing files.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal volume rm`

Delete a file or directory from a modal.Volume.

**Usage** :

    modal volume rm [OPTIONS] VOLUME_NAME REMOTE_PATH

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

## `modal volume cp`

Copy within a modal.Volume. Copy source file to destination file or multiple
source files to destination directory.

**Usage** :

    modal volume cp [OPTIONS] VOLUME_NAME PATHS...

Copy

**Arguments** :

  * `VOLUME_NAME`: [required]
  * `PATHS...`: [required]

**Options** :

  * `-r, --recursive`: Copy directories recursively
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal volume delete`

Delete a named, persistent modal.Volume.

**Usage** :

    modal volume delete [OPTIONS] VOLUME_NAME

Copy

**Arguments** :

  * `VOLUME_NAME`: Name of the modal.Volume to be deleted. Case sensitive [required]

**Options** :

  * `-y, --yes`: Run without pausing for confirmation.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal volume rename`

Rename a modal.Volume.

**Usage** :

    modal volume rename [OPTIONS] OLD_NAME NEW_NAME

Copy

**Arguments** :

  * `OLD_NAME`: [required]
  * `NEW_NAME`: [required]

**Options** :

  * `-y, --yes`: Run without pausing for confirmation.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

modal volumemodal volume createmodal volume getmodal volume listmodal volume
lsmodal volume putmodal volume rmmodal volume cpmodal volume deletemodal
volume rename
