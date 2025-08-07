* * *

Copy page

# modal.enable_output

    @contextlib.contextmanager
    def enable_output(show_progress: bool = True) -> Generator[None, None, None]:

Copy

Context manager that enable output when using the Python SDK.

This will print to stdout and stderr things such as

  1. Logs from running functions
  2. Status of creating objects
  3. Map progress

Example:

    app = modal.App()
    with modal.enable_output():
        with app.run():
            ...

Copy

modal.enable_output
