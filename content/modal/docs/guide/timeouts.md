* * *

Copy page

# Timeouts

All Modal [Function](/docs/reference/modal.Function) executions have a default
execution timeout of 300 seconds (5 minutes), but users may specify timeout
durations between 1 second and 24 hours.

    import time

    @app.function()
    def f():
        time.sleep(599)  # Timeout!

    @app.function(timeout=600)
    def g():
        time.sleep(599)
        print("*Just* made it!")

Copy

The timeout duration is a measure of a Function’s _execution_ time. It does
not include scheduling time or any other period besides the time your code is
executing in Modal. This duration is also per execution attempt, meaning
Functions configured with [`modal.Retries`](/docs/reference/modal.Retries)
will start new execution timeouts on each retry. For example, an infinite-
looping Function with a 100 second timeout and 3 allowed retries will run for
least 400 seconds within Modal.

### Container startup timeout

Currently `timeout` applies also to the container’s _startup_ time as well as
its execution time. If you’re container is failing for spending too long
initializing, extend the `timeout` of your Function. In the future we will
decouple startup timeouts from execution timeouts.

## Handling timeouts

After exhausting any specified retries, a timeout in a Function will produce a
`modal.exception.FunctionTimeoutError` which you may catch in your code.

    import modal.exception

    @app.function(timeout=100)
    def f():
        time.sleep(200)  # Timeout!

    @app.local_entrypoint()
    def main():
        try:
            f.remote()
        except modal.exception.FunctionTimeoutError:
            ... # Handle the timeout.

Copy

## Timeout accuracy

Functions will run for _at least_ as long as their timeout allows, but they
may run a handful of seconds longer. If you require accurate and precise
timeout durations on your Function executions, it is recommended that you
implement timeout logic in your user code.

TimeoutsContainer startup timeoutHandling timeoutsTimeout accuracy
