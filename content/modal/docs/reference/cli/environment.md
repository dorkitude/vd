* * *

Copy page

# `modal environment`

Create and interact with Environments

Environments are sub-divisons of workspaces, allowing you to deploy the same
app in different namespaces. Each environment has their own set of Secrets and
any lookups performed from an app in an environment will by default look for
entities in the same environment.

Typical use cases for environments include having one for development and one
for production, to prevent overwriting production apps when developing new
features while still being able to deploy changes to a live environment.

**Usage** :

    modal environment [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `list`: List all environments in the current workspace
  * `create`: Create a new environment in the current workspace
  * `delete`: Delete an environment in the current workspace
  * `update`: Update the name or web suffix of an environment

## `modal environment list`

List all environments in the current workspace

**Usage** :

    modal environment list [OPTIONS]

Copy

**Options** :

  * `--json / --no-json`: [default: no-json]
  * `--help`: Show this message and exit.

## `modal environment create`

Create a new environment in the current workspace

**Usage** :

    modal environment create [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: Name of the new environment. Must be unique. Case sensitive [required]

**Options** :

  * `--help`: Show this message and exit.

## `modal environment delete`

Delete an environment in the current workspace

Deletes all apps in the selected environment and deletes the environment
irrevocably.

**Usage** :

    modal environment delete [OPTIONS] NAME

Copy

**Arguments** :

  * `NAME`: Name of the environment to be deleted. Case sensitive [required]

**Options** :

  * `--confirm / --no-confirm`: Set this flag to delete without prompting for confirmation [default: no-confirm]
  * `--help`: Show this message and exit.

## `modal environment update`

Update the name or web suffix of an environment

**Usage** :

    modal environment update [OPTIONS] CURRENT_NAME

Copy

**Arguments** :

  * `CURRENT_NAME`: [required]

**Options** :

  * `--set-name TEXT`: New name of the environment
  * `--set-web-suffix TEXT`: New web suffix of environment (empty string is no suffix)
  * `--help`: Show this message and exit.

modal environmentmodal environment listmodal environment createmodal
environment deletemodal environment update
