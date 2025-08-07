* * *

Copy page

# `modal dict`

Manage `modal.Dict` objects and inspect their contents.

**Usage** :

    modal dict [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `create`: Create a named Dict object.
  * `list`: List all named Dicts.
  * `clear`: Clear the contents of a named Dict by deleting all of its data.
  * `delete`: Delete a named Dict and all of its data.
  * `get`: Print the value for a specific key.
  * `items`: Print the contents of a Dict.

## `modal dict create`

Create a named Dict object.

Note: This is a no-op when the Dict already exists.

**Usage** :

    modal dict create [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: [required]

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal dict list`

List all named Dicts.

**Usage** :

    modal dict list [OPTIONS]

Copy

**Options** :

  * `--json / --no-json`: [default: no-json]
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal dict clear`

Clear the contents of a named Dict by deleting all of its data.

**Usage** :

    modal dict clear [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: [required]

**Options** :

  * `-y, --yes`: Run without pausing for confirmation.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal dict delete`

Delete a named Dict and all of its data.

**Usage** :

    modal dict delete [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: [required]

**Options** :

  * `-y, --yes`: Run without pausing for confirmation.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal dict get`

Print the value for a specific key.

Note: When using the CLI, keys are always interpreted as having a string type.

**Usage** :

    modal dict get [OPTIONS] NAME KEY

Copy

**Arguments** :

  * `NAME`: [required]
  * `KEY`: [required]

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal dict items`

Print the contents of a Dict.

Note: By default, this command truncates the contents. Use the `N` argument to
control the amount of data shown or the `--all` option to retrieve the entire
Dict, which may be slow.

**Usage** :

    modal dict items [OPTIONS] NAME [N]

Copy

**Arguments** :

  * `NAME`: [required]
  * `[N]`: Limit the number of entries shown [default: 20]

**Options** :

  * `-a, --all`: Ignore N and print all entries in the Dict (may be slow)
  * `-r, --repr`: Display items using `repr()` to see more details
  * `--json / --no-json`: [default: no-json]
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

modal dictmodal dict createmodal dict listmodal dict clearmodal dict
deletemodal dict getmodal dict items
