* * *

Copy page

# modal.Client

    class Client(object)

Copy

## is_closed

    def is_closed(self) -> bool:

Copy

## hello

    def hello(self):

Copy

Connect to server and retrieve version information; raise appropriate error
for various failures.

## from_credentials

    @classmethod
    def from_credentials(cls, token_id: str, token_secret: str) -> "_Client":

Copy

Constructor based on token credentials; useful for managing Modal on behalf of
third-party users.

**Usage:**

    client = modal.Client.from_credentials("my_token_id", "my_token_secret")

    modal.Sandbox.create("echo", "hi", client=client, app=app)

Copy

## get_input_plane_metadata

    def get_input_plane_metadata(self, input_plane_region: str) -> list[tuple[str, str]]:

Copy

modal.Clientis_closedhellofrom_credentialsget_input_plane_metadata
