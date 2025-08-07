* * *

Copy page

# modal.Retries

    class Retries(object)

Copy

Adds a retry policy to a Modal function.

**Usage**

    import modal
    app = modal.App()

    # Basic configuration.
    # This sets a policy of max 4 retries with 1-second delay between failures.
    @app.function(retries=4)
    def f():
        pass

    # Fixed-interval retries with 3-second delay between failures.
    @app.function(
        retries=modal.Retries(
            max_retries=2,
            backoff_coefficient=1.0,
            initial_delay=3.0,
        )
    )
    def g():
        pass

    # Exponential backoff, with retry delay doubling after each failure.
    @app.function(
        retries=modal.Retries(
            max_retries=4,
            backoff_coefficient=2.0,
            initial_delay=1.0,
        )
    )
    def h():
        pass

Copy

    def __init__(
        self,
        *,
        # The maximum number of retries that can be made in the presence of failures.
        max_retries: int,
        # Coefficent controlling how much the retry delay increases each retry attempt.
        # A backoff coefficient of 1.0 creates fixed-delay where the delay period always equals the initial delay.
        backoff_coefficient: float = 2.0,
        # Number of seconds that must elapse before the first retry occurs.
        initial_delay: float = 1.0,
        # Maximum length of retry delay in seconds, preventing the delay from growing infinitely.
        max_delay: float = 60.0,
    ):

Copy

Construct a new retries policy, supporting exponential and fixed-interval
delays via a backoff coefficient.

modal.Retries
