* * *

Copy page

# modal.current_input_id

    def current_input_id() -> Optional[str]:

Copy

Returns the input ID for the current input.

Can only be called from Modal function (i.e. in a container context).

    from modal import current_input_id

    @app.function()
    def process_stuff():
        print(f"Starting to process {current_input_id()}")

Copy

modal.current_input_id
