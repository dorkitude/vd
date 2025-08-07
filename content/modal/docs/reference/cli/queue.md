* * *

Copy page

# `modal queue`

Manage `modal.Queue` objects and inspect their contents.

**Usage** :

    modal queue [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `create`: Create a named Queue.
  * `delete`: Delete a named Queue and all of its data.
  * `list`: List all named Queues.
  * `clear`: Clear the contents of a queue by removing all of its data.
  * `peek`: Print the next N items in the queue or queue partition (without removal).
  * `len`: Print the length of a queue partition or the total length of all partitions.

## `modal queue create`

Create a named Queue.

Note: This is a no-op when the Queue already exists.

**Usage** :

    modal queue create [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: [required]

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal queue delete`

Delete a named Queue and all of its data.

**Usage** :

    modal queue delete [OPTIONS] NAME

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

## `modal queue list`

List all named Queues.

**Usage** :

    modal queue list [OPTIONS]

Copy

**Options** :

  * `--json / --no-json`: [default: no-json]
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal queue clear`

Clear the contents of a queue by removing all of its data.

**Usage** :

    modal queue clear [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: [required]

**Options** :

  * `-p, --partition TEXT`: Name of the partition to use, otherwise use the default (anonymous) partition.
  * `-a, --all`: Clear the contents of all partitions.
  * `-y, --yes`: Run without pausing for confirmation.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal queue peek`

Print the next N items in the queue or queue partition (without removal).

**Usage** :

    modal queue peek [OPTIONS] NAME [N]

Copy

**Arguments** :

  * `NAME`: [required]
  * `[N]`: [default: 1]

**Options** :

  * `-p, --partition TEXT`: Name of the partition to use, otherwise use the default (anonymous) partition.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal queue len`

Print the length of a queue partition or the total length of all partitions.

**Usage** :

    modal queue len [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: [required]

**Options** :

  * `-p, --partition TEXT`: Name of the partition to use, otherwise use the default (anonymous) partition.
  * `-t, --total`: Compute the sum of the queue lengths across all partitions
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

modal queuemodal queue createmodal queue deletemodal queue listmodal queue
clearmodal queue peekmodal queue len
