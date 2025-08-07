* * *

Copy page

# modal.Error

    class Error(Exception)

Copy

Base class for all Modal errors. See
[`modal.exception`](https://modal.com/docs/reference/modal.exception) for the
specialized error classes.

**Usage**

    import modal

    try:
        ...
    except modal.Error:
        # Catch any exception raised by Modal's systems.
        print("Responding to error...")

Copy

modal.Error
