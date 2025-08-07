* * *

Copy page

# `modal run`

Run a Modal function or local entrypoint.

`FUNC_REF` should be of the format `{file or module}::{function name}`.
Alternatively, you can refer to the function via the app:

`{file or module}::{app variable name}.{function name}`

**Examples:**

To run the hello_world function (or local entrypoint) in my_app.py:

    modal run my_app.py::hello_world

Copy

If your module only has a single app and your app has a single local
entrypoint (or single function), you can omit the app and function parts:

    modal run my_app.py

Copy

Instead of pointing to a file, you can also use the Python module path, which
by default will ensure that your remote functions will use the same module
names as they do locally.

    modal run -m my_project.my_app

Copy

**Usage** :

    modal run [OPTIONS] FUNC_REF

Copy

**Options** :

  * `-w, --write-result TEXT`: Write return value (which must be str or bytes) to this local path.
  * `-q, --quiet`: Don’t show Modal progress indicators.
  * `-d, --detach`: Don’t stop the app if the local process dies or disconnects.
  * `-i, --interactive`: Run the app in interactive mode.
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `-m`: Interpret argument as a Python module path instead of a file/script path
  * `--help`: Show this message and exit.

modal run
