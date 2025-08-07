* * *

Copy page

# `modal secret`

Manage secrets.

**Usage** :

    modal secret [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `list`: List your published secrets.
  * `create`: Create a new secret.
  * `delete`: Delete a named secret.

## `modal secret list`

List your published secrets.

**Usage** :

    modal secret list [OPTIONS]

Copy

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--json / --no-json`: [default: no-json]
  * `--help`: Show this message and exit.

## `modal secret create`

Create a new secret.

**Usage** :

    modal secret create [OPTIONS] SECRET_NAME [KEYVALUES]...

Copy

**Arguments** :

  * `SECRET_NAME`: [required]
  * `[KEYVALUES]...`: Space-separated KEY=VALUE items.

**Options** :

  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--from-dotenv PATH`: Path to a .env file to load secrets from.
  * `--from-json PATH`: Path to a JSON file to load secrets from.
  * `--force`: Overwrite the secret if it already exists.
  * `--help`: Show this message and exit.

## `modal secret delete`

Delete a named secret.

**Usage** :

    modal secret delete [OPTIONS] SECRET_NAME

Copy

**Arguments** :

  * `SECRET_NAME`: Name of the modal.Secret to be deleted. Case sensitive [required]

**Options** :

  * `-y, --yes`: Run without pausing for confirmation.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--help`: Show this message and exit.

modal secretmodal secret listmodal secret createmodal secret delete
