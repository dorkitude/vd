* * *

Copy page

# `modal token`

Manage tokens.

**Usage** :

    modal token [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `set`: Set account credentials for connecting to Modal.
  * `new`: Create a new token by using an authenticated web session.

## `modal token set`

Set account credentials for connecting to Modal.

If the credentials are not provided on the command line, you will be prompted
to enter them.

**Usage** :

    modal token set [OPTIONS]

Copy

**Options** :

  * `--token-id TEXT`: Account token ID.
  * `--token-secret TEXT`: Account token secret.
  * `--profile TEXT`: Modal profile to set credentials for. If unspecified (and MODAL_PROFILE environment variable is not set), uses the workspace name associated with the credentials.
  * `--activate / --no-activate`: Activate the profile containing this token after creation. [default: activate]
  * `--verify / --no-verify`: Make a test request to verify the new credentials. [default: verify]
  * `--help`: Show this message and exit.

## `modal token new`

Create a new token by using an authenticated web session.

**Usage** :

    modal token new [OPTIONS]

Copy

**Options** :

  * `--profile TEXT`: Modal profile to set credentials for. If unspecified (and MODAL_PROFILE environment variable is not set), uses the workspace name associated with the credentials.
  * `--activate / --no-activate`: Activate the profile containing this token after creation. [default: activate]
  * `--verify / --no-verify`: Make a test request to verify the new credentials. [default: verify]
  * `--source TEXT`
  * `--help`: Show this message and exit.

modal tokenmodal token setmodal token new
