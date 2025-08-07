* * *

Copy page

# `modal app`

Manage deployed and running apps.

**Usage** :

    modal app [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `list`: List Modal apps that are currently deployed/running or recently stopped.
  * `logs`: Show App logs, streaming while active.
  * `rollback`: Redeploy a previous version of an App.
  * `stop`: Stop an app.
  * `history`: Show App deployment history, for a currently deployed app

## `modal app list`

List Modal apps that are currently deployed/running or recently stopped.

**Usage** :

    modal app list [OPTIONS]

Copy

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--json / --no-json`: [default: no-json]
  * `--help`: Show this message and exit.

## `modal app logs`

Show App logs, streaming while active.

**Examples:**

Get the logs based on an app ID:

    modal app logs ap-123456

Copy

Get the logs for a currently deployed App based on its name:

    modal app logs my-app

Copy

**Usage** :

    modal app logs [OPTIONS] [APP_IDENTIFIER]

Copy

**Arguments** :

  * `[APP_IDENTIFIER]`: App name or ID

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--timestamps`: Show timestamps for each log line
  * `--help`: Show this message and exit.

## `modal app rollback`

Redeploy a previous version of an App.

Note that the App must currently be in a “deployed” state. Rollbacks will
appear as a new deployment in the App history, although the App state will be
reset to the state at the time of the previous deployment.

**Examples:**

Rollback an App to its previous version:

    modal app rollback my-app

Copy

Rollback an App to a specific version:

    modal app rollback my-app v3

Copy

Rollback an App using its App ID instead of its name:

    modal app rollback ap-abcdefghABCDEFGH123456

Copy

**Usage** :

    modal app rollback [OPTIONS] [APP_IDENTIFIER] [VERSION]

Copy

**Arguments** :

  * `[APP_IDENTIFIER]`: App name or ID
  * `[VERSION]`: Target version for rollback.

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal app stop`

Stop an app.

**Usage** :

    modal app stop [OPTIONS] [APP_IDENTIFIER]

Copy

**Arguments** :

  * `[APP_IDENTIFIER]`: App name or ID

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

## `modal app history`

Show App deployment history, for a currently deployed app

**Examples:**

Get the history based on an app ID:

    modal app history ap-123456

Copy

Get the history for a currently deployed App based on its name:

    modal app history my-app

Copy

**Usage** :

    modal app history [OPTIONS] [APP_IDENTIFIER]

Copy

**Arguments** :

  * `[APP_IDENTIFIER]`: App name or ID

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--json / --no-json`: [default: no-json]
  * `--help`: Show this message and exit.

modal appmodal app listmodal app logsmodal app rollbackmodal app stopmodal app
history
