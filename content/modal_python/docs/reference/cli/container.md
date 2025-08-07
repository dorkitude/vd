* * *

Copy page

# `modal container`

Manage and connect to running containers.

**Usage** :

    modal container [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `list`: List all containers that are currently running.
  * `logs`: Show logs for a specific container, streaming while active.
  * `exec`: Execute a command in a container.
  * `stop`: Stop a currently-running container and reassign its in-progress inputs.

## `modal container list`

List all containers that are currently running.

**Usage** :

    modal container list [OPTIONS]

Copy

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--json / --no-json`: [default: no-json]
  * `--help`: Show this message and exit.

## `modal container logs`

Show logs for a specific container, streaming while active.

**Usage** :

    modal container logs [OPTIONS] CONTAINER_ID

Copy

**Arguments** :

  * `CONTAINER_ID`: Container ID [required]

**Options** :

  * `--help`: Show this message and exit.

## `modal container exec`

Execute a command in a container.

**Usage** :

    modal container exec [OPTIONS] CONTAINER_ID COMMAND...

Copy

**Arguments** :

  * `CONTAINER_ID`: Container ID [required]
  * `COMMAND...`: A command to run inside the container.

To pass command-line flags or options, add `--` before the start of your
commands. For example: `modal container exec <id> -- /bin/bash -c 'echo hi'`
[required]

**Options** :

  * `--pty / --no-pty`: Run the command using a PTY.
  * `--help`: Show this message and exit.

## `modal container stop`

Stop a currently-running container and reassign its in-progress inputs.

This will send the container a SIGINT signal that Modal will handle.

**Usage** :

    modal container stop [OPTIONS] CONTAINER_ID

Copy

**Arguments** :

  * `CONTAINER_ID`: Container ID [required]

**Options** :

  * `--help`: Show this message and exit.

modal containermodal container listmodal container logsmodal container
execmodal container stop
