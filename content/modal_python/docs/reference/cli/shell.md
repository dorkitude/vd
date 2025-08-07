* * *

Copy page

# `modal shell`

Run a command or interactive shell inside a Modal container.

**Examples:**

Start an interactive shell inside the default Debian-based image:

    modal shell

Copy

Start an interactive shell with the spec for `my_function` in your App (uses
the same image, volumes, mounts, etc.):

    modal shell hello_world.py::my_function

Copy

Or, if you’re using a [modal.Cls](https://modal.com/docs/reference/modal.Cls)
you can refer to a `@modal.method` directly:

    modal shell hello_world.py::MyClass.my_method

Copy

Start a `python` shell:

    modal shell hello_world.py --cmd=python

Copy

Run a command with your function’s spec and pipe the output to a file:

    modal shell hello_world.py -c 'uv pip list' > env.txt

Copy

**Usage** :

    modal shell [OPTIONS] REF

Copy

**Arguments** :

  * `REF`: ID of running container, or path to a Python file containing a Modal App. Can also include a function specifier, like `module.py::func`, if the file defines multiple functions.

**Options** :

  * `-c, --cmd TEXT`: Command to run inside the Modal image. [default: /bin/bash]
  * `-e, --env TEXT`: Environment to interact with.

If not specified, Modal will use the default environment of your current
profile, or the `MODAL_ENVIRONMENT` variable. Otherwise, raises an error if
the workspace has multiple environments.

  * `--image TEXT`: Container image tag for inside the shell (if not using REF).
  * `--add-python TEXT`: Add Python to the image (if not using REF).
  * `--volume TEXT`: Name of a `modal.Volume` to mount inside the shell at `/mnt/{name}` (if not using REF). Can be used multiple times.
  * `--secret TEXT`: Name of a `modal.Secret` to mount inside the shell (if not using REF). Can be used multiple times.
  * `--cpu INTEGER`: Number of CPUs to allocate to the shell (if not using REF).
  * `--memory INTEGER`: Memory to allocate for the shell, in MiB (if not using REF).
  * `--gpu TEXT`: GPUs to request for the shell, if any. Examples are `any`, `a10g`, `a100:4` (if not using REF).
  * `--cloud TEXT`: Cloud provider to run the shell on. Possible values are `aws`, `gcp`, `oci`, `auto` (if not using REF).
  * `--region TEXT`: Region(s) to run the container on. Can be a single region or a comma-separated list to choose from (if not using REF).
  * `--pty / --no-pty`: Run the command using a PTY.
  * `-m`: Interpret argument as a Python module path instead of a file/script path
  * `--help`: Show this message and exit.

modal shell
