* * *

Copy page

# modal.Secret

    class Secret(modal.object.Object)

Copy

Secrets provide a dictionary of environment variables for images.

Secrets are a secure way to add credentials and other sensitive information to
the containers your functions run in. You can create and edit secrets on [the
dashboard](https://modal.com/secrets), or programmatically from Python code.

See [the secrets guide page](https://modal.com/docs/guide/secrets) for more
information.

## hydrate

    def hydrate(self, client: Optional[_Client] = None) -> Self:

Copy

Synchronize the local object with its identity on the Modal server.

It is rarely necessary to call this method explicitly, as most operations will
lazily hydrate when needed. The main use case is when you need to access
object metadata, such as its ID.

_Added in v0.72.39_ : This method replaces the deprecated `.resolve()` method.

## name

    @property
    def name(self) -> Optional[str]:

Copy

## from_dict

    @staticmethod
    def from_dict(
        env_dict: dict[
            str, Union[str, None]
        ] = {},  # dict of entries to be inserted as environment variables in functions using the secret
    ) -> "_Secret":

Copy

Create a secret from a str-str dictionary. Values can also be `None`, which is
ignored.

Usage:

    @app.function(secrets=[modal.Secret.from_dict({"FOO": "bar"})])
    def run():
        print(os.environ["FOO"])

Copy

## from_local_environ

    @staticmethod
    def from_local_environ(
        env_keys: list[str],  # list of local env vars to be included for remote execution
    ) -> "_Secret":

Copy

Create secrets from local environment variables automatically.

## from_dotenv

    @staticmethod
    def from_dotenv(path=None, *, filename=".env") -> "_Secret":

Copy

Create secrets from a .env file automatically.

If no argument is provided, it will use the current working directory as the
starting point for finding a `.env` file. Note that it does not use the
location of the module calling `Secret.from_dotenv`.

If called with an argument, it will use that as a starting point for finding
`.env` files. In particular, you can call it like this:

    @app.function(secrets=[modal.Secret.from_dotenv(__file__)])
    def run():
        print(os.environ["USERNAME"])  # Assumes USERNAME is defined in your .env file

Copy

This will use the location of the script calling `modal.Secret.from_dotenv` as
a starting point for finding the `.env` file.

A file named `.env` is expected by default, but this can be overridden with
the `filename` keyword argument:

    @app.function(secrets=[modal.Secret.from_dotenv(filename=".env-dev")])
    def run():
        ...

Copy

## from_name

    @staticmethod
    def from_name(
        name: str,
        *,
        environment_name: Optional[str] = None,
        required_keys: list[
            str
        ] = [],  # Optionally, a list of required environment variables (will be asserted server-side)
    ) -> "_Secret":

Copy

Reference a Secret by its name.

In contrast to most other Modal objects, named Secrets must be provisioned
from the Dashboard. See other methods for alternate ways of creating a new
Secret from code.

    secret = modal.Secret.from_name("my-secret")

    @app.function(secrets=[secret])
    def run():
       ...

Copy

## info

    @live_method
    def info(self) -> SecretInfo:

Copy

Return information about the Secret object.

modal.Secrethydratenamefrom_dictfrom_local_environfrom_dotenvfrom_nameinfo
