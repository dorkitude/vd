* * *

Copy page

# `modal config`

Manage client configuration for the current profile.

Refer to <https://modal.com/docs/reference/modal.config> for a full
explanation of what these options mean, and how to set them.

**Usage** :

    modal config [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `show`: Show current configuration values (debugging command).
  * `set-environment`: Set the default Modal environment for the active profile

## `modal config show`

Show current configuration values (debugging command).

**Usage** :

    modal config show [OPTIONS]

Copy

**Options** :

  * `--redact / --no-redact`: Redact the `token_secret` value. [default: redact]
  * `--help`: Show this message and exit.

## `modal config set-environment`

Set the default Modal environment for the active profile

The default environment of a profile is used when no —env flag is passed to
`modal run`, `modal deploy` etc.

If no default environment is set, and there exists multiple environments in a
workspace, an error will be raised when running a command that requires an
environment.

**Usage** :

    modal config set-environment [OPTIONS] ENVIRONMENT_NAME

Copy

**Arguments** :

  * `ENVIRONMENT_NAME`: [required]

**Options** :

  * `--help`: Show this message and exit.

modal configmodal config showmodal config set-environment
