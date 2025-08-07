* * *

Copy page

# modal.call_graph

## modal.call_graph.InputInfo

    class InputInfo(object)

Copy

Simple data structure storing information about a function input.

    def __init__(self, input_id: str, function_call_id: str, task_id: str, status: modal.call_graph.InputStatus, function_name: str, module_name: str, children: list['InputInfo']) -> None

Copy

## modal.call_graph.InputStatus

    class InputStatus(enum.IntEnum)

Copy

Enum representing status of a function input.

The possible values are:

  * `PENDING`
  * `SUCCESS`
  * `FAILURE`
  * `INIT_FAILURE`
  * `TERMINATED`
  * `TIMEOUT`

modal.call_graphmodal.call_graph.InputInfomodal.call_graph.InputStatus
