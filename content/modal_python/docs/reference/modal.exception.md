* * *

Copy page

# modal.exception

## modal.exception.AlreadyExistsError

    class AlreadyExistsError(modal.exception.Error)

Copy

Raised when a resource creation conflicts with an existing resource.

## modal.exception.AuthError

    class AuthError(modal.exception.Error)

Copy

Raised when a client has missing or invalid authentication.

## modal.exception.ClientClosed

    class ClientClosed(modal.exception.Error)

Copy

## modal.exception.ConnectionError

    class ConnectionError(modal.exception.Error)

Copy

Raised when an issue occurs while connecting to the Modal servers.

## modal.exception.DeprecationError

    class DeprecationError(UserWarning)

Copy

UserWarning category emitted when a deprecated Modal feature or API is used.

## modal.exception.DeserializationError

    class DeserializationError(modal.exception.Error)

Copy

Raised to provide more context when an error is encountered during
deserialization.

## modal.exception.ExecutionError

    class ExecutionError(modal.exception.Error)

Copy

Raised when something unexpected happened during runtime.

## modal.exception.FilesystemExecutionError

    class FilesystemExecutionError(modal.exception.Error)

Copy

Raised when an unknown error is thrown during a container filesystem
operation.

## modal.exception.FunctionTimeoutError

    class FunctionTimeoutError(modal.exception.TimeoutError)

Copy

Raised when a Function exceeds its execution duration limit and times out.

## modal.exception.InputCancellation

    class InputCancellation(BaseException)

Copy

Raised when the current input is cancelled by the task

Intentionally a BaseException instead of an Exception, so it won’t get caught
by unspecified user exception clauses that might be used for retries and other
control flow.

## modal.exception.InteractiveTimeoutError

    class InteractiveTimeoutError(modal.exception.TimeoutError)

Copy

Raised when interactive frontends time out while trying to connect to a
container.

## modal.exception.InternalFailure

    class InternalFailure(modal.exception.Error)

Copy

Retriable internal error.

## modal.exception.InvalidError

    class InvalidError(modal.exception.Error)

Copy

Raised when user does something invalid.

## modal.exception.ModuleNotMountable

    class ModuleNotMountable(Exception)

Copy

## modal.exception.MountUploadTimeoutError

    class MountUploadTimeoutError(modal.exception.TimeoutError)

Copy

Raised when a Mount upload times out.

## modal.exception.NotFoundError

    class NotFoundError(modal.exception.Error)

Copy

Raised when a requested resource was not found.

## modal.exception.OutputExpiredError

    class OutputExpiredError(modal.exception.TimeoutError)

Copy

Raised when the Output exceeds expiration and times out.

## modal.exception.PendingDeprecationError

    class PendingDeprecationError(UserWarning)

Copy

Soon to be deprecated feature. Only used intermittently because of multi-repo
concerns.

## modal.exception.RemoteError

    class RemoteError(modal.exception.Error)

Copy

Raised when an error occurs on the Modal server.

## modal.exception.RequestSizeError

    class RequestSizeError(modal.exception.Error)

Copy

Raised when an operation produces a gRPC request that is rejected by the
server for being too large.

## modal.exception.SandboxTerminatedError

    class SandboxTerminatedError(modal.exception.Error)

Copy

Raised when a Sandbox is terminated for an internal reason.

## modal.exception.SandboxTimeoutError

    class SandboxTimeoutError(modal.exception.TimeoutError)

Copy

Raised when a Sandbox exceeds its execution duration limit and times out.

## modal.exception.SerializationError

    class SerializationError(modal.exception.Error)

Copy

Raised to provide more context when an error is encountered during
serialization.

## modal.exception.ServerWarning

    class ServerWarning(UserWarning)

Copy

Warning originating from the Modal server and re-issued in client code.

## modal.exception.TimeoutError

    class TimeoutError(modal.exception.Error)

Copy

Base class for Modal timeouts.

## modal.exception.VersionError

    class VersionError(modal.exception.Error)

Copy

Raised when the current client version of Modal is unsupported.

## modal.exception.VolumeUploadTimeoutError

    class VolumeUploadTimeoutError(modal.exception.TimeoutError)

Copy

Raised when a Volume upload times out.

## modal.exception.simulate_preemption

    def simulate_preemption(wait_seconds: int, jitter_seconds: int = 0):

Copy

Utility for simulating a preemption interrupt after `wait_seconds` seconds.
The first interrupt is the SIGINT signal. After 30 seconds, a second interrupt
will trigger.

This second interrupt simulates SIGKILL, and should not be caught. Optionally
add between zero and `jitter_seconds` seconds of additional waiting before
first interrupt.

**Usage:**

    import time
    from modal.exception import simulate_preemption

    simulate_preemption(3)

    try:
        time.sleep(4)
    except KeyboardInterrupt:
        print("got preempted") # Handle interrupt
        raise

Copy

See <https://modal.com/docs/guide/preemption> for more details on preemption
handling.

modal.exceptionmodal.exception.AlreadyExistsErrormodal.exception.AuthErrormodal.exception.ClientClosedmodal.exception.ConnectionErrormodal.exception.DeprecationErrormodal.exception.DeserializationErrormodal.exception.ExecutionErrormodal.exception.FilesystemExecutionErrormodal.exception.FunctionTimeoutErrormodal.exception.InputCancellationmodal.exception.InteractiveTimeoutErrormodal.exception.InternalFailuremodal.exception.InvalidErrormodal.exception.ModuleNotMountablemodal.exception.MountUploadTimeoutErrormodal.exception.NotFoundErrormodal.exception.OutputExpiredErrormodal.exception.PendingDeprecationErrormodal.exception.RemoteErrormodal.exception.RequestSizeErrormodal.exception.SandboxTerminatedErrormodal.exception.SandboxTimeoutErrormodal.exception.SerializationErrormodal.exception.ServerWarningmodal.exception.TimeoutErrormodal.exception.VersionErrormodal.exception.VolumeUploadTimeoutErrormodal.exception.simulate_preemption
