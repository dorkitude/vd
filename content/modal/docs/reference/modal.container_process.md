* * *

Copy page

# modal.container_process

## modal.container_process.ContainerProcess

    class ContainerProcess(typing.Generic)

Copy

    def __init__(
        self,
        process_id: str,
        client: _Client,
        stdout: StreamType = StreamType.PIPE,
        stderr: StreamType = StreamType.PIPE,
        exec_deadline: Optional[float] = None,
        text: bool = True,
        by_line: bool = False,
    ) -> None:

Copy

### stdout

    @property
    def stdout(self) -> _StreamReader[T]:

Copy

StreamReader for the container process’s stdout stream.

### stderr

    @property
    def stderr(self) -> _StreamReader[T]:

Copy

StreamReader for the container process’s stderr stream.

### stdin

    @property
    def stdin(self) -> _StreamWriter:

Copy

StreamWriter for the container process’s stdin stream.

### returncode

    @property
    def returncode(self) -> int:

Copy

### poll

    def poll(self) -> Optional[int]:

Copy

Check if the container process has finished running.

Returns `None` if the process is still running, else returns the exit code.

### wait

    def wait(self) -> int:

Copy

Wait for the container process to finish running. Returns the exit code.

modal.container_processmodal.container_process.ContainerProcessstdoutstderrstdinreturncodepollwait
