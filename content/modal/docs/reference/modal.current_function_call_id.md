* * *

Copy page

# modal.current_function_call_id

    def current_function_call_id() -> Optional[str]:

Copy

Returns the function call ID for the current input.

Can only be called from Modal function (i.e. in a container context).

    from modal import current_function_call_id

    @app.function()
    def process_stuff():
        print(f"Starting to process input from {current_function_call_id()}")

Copy

modal.current_function_call_id
