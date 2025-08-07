* * *

Copy page

# `modal deploy`

Deploy a Modal application.

**Usage:** modal deploy my_script.py modal deploy -m my_package.my_mod

**Usage** :

    modal deploy [OPTIONS] APP_REF

Copy

**Arguments** :

  * `APP_REF`: Path to a Python file with an app to deploy [required]

**Options** :

  * `--name TEXT`: Name of the deployment.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--stream-logs / --no-stream-logs`: Stream logs from the app upon deployment. [default: no-stream-logs]
  * `--tag TEXT`: Tag the deployment with a version.
  * `-m`: Interpret argument as a Python module path instead of a file/script path
  * `--help`: Show this message and exit.

modal deploy
