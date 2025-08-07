* * *

Copy page

# `modal serve`

Run a web endpoint(s) associated with a Modal app and hot-reload code.

**Examples:**

    modal serve hello_world.py

Copy

**Usage** :

    modal serve [OPTIONS] APP_REF

Copy

**Arguments** :

  * `APP_REF`: Path to a Python file with an app. [required]

**Options** :

  * `--timeout FLOAT`
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `-m`: Interpret argument as a Python module path instead of a file/script path
  * `--help`: Show this message and exit.

modal serve
