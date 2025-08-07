* * *

Copy page

# `modal launch`

Open a serverless app instance on Modal.

This command is in preview and may change in the future.

**Usage** :

    modal launch [OPTIONS] COMMAND [ARGS]...

Copy

**Options** :

  * `--help`: Show this message and exit.

**Commands** :

  * `jupyter`: Start Jupyter Lab on Modal.
  * `vscode`: Start Visual Studio Code on Modal.

## `modal launch jupyter`

Start Jupyter Lab on Modal.

**Usage** :

    modal launch jupyter [OPTIONS]

Copy

**Options** :

  * `--cpu INTEGER`: [default: 8]
  * `--memory INTEGER`: [default: 32768]
  * `--gpu TEXT`
  * `--timeout INTEGER`: [default: 3600]
  * `--image TEXT`: [default: ubuntu:22.04]
  * `--add-python TEXT`: [default: 3.11]
  * `--mount TEXT`
  * `--volume TEXT`
  * `--detach / --no-detach`: [default: no-detach]
  * `--help`: Show this message and exit.

## `modal launch vscode`

Start Visual Studio Code on Modal.

**Usage** :

    modal launch vscode [OPTIONS]

Copy

**Options** :

  * `--cpu INTEGER`: [default: 8]
  * `--memory INTEGER`: [default: 32768]
  * `--gpu TEXT`
  * `--image TEXT`: [default: debian:12]
  * `--timeout INTEGER`: [default: 3600]
  * `--mount TEXT`
  * `--volume TEXT`
  * `--detach / --no-detach`: [default: no-detach]
  * `--help`: Show this message and exit.

modal launchmodal launch jupytermodal launch vscode
