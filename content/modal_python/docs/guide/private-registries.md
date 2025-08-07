* * *

Copy page

# Private registries

Modal provides the [`Image.from_registry`](/docs/guide/images#use-an-existing-
container-image-with-from_registry) function, which can pull public images
available from registries such as Docker Hub and GitHub Container Registry, as
well as private images from registries such as [AWS Elastic Container Registry
(ECR)](https://aws.amazon.com/ecr/), [GCP Artifact
Registry](https://cloud.google.com/artifact-registry), and Docker Hub.

## Docker Hub (Private)

To pull container images from private Docker Hub repositories, [create an
access token](https://docs.docker.com/security/for-developers/access-tokens/)
with “Read-Only” permissions and use this token value and your Docker Hub
username to create a Modal [Secret](/docs/guide/secrets).

    REGISTRY_USERNAME=my-dockerhub-username
    REGISTRY_PASSWORD=<your-docker-access-token>

Copy

Use this Secret with the
[`modal.Image.from_registry`](/docs/reference/modal.Image#from_registry)
method.

## Elastic Container Registry (ECR)

You can pull images from your AWS ECR account by specifying the full image URI
as follows:

    import modal

    aws_secret = modal.Secret.from_name("my-aws-secret")
    image = (
        modal.Image.from_aws_ecr(
            "000000000000.dkr.ecr.us-east-1.amazonaws.com/my-private-registry:latest",
            secret=aws_secret,
        )
        .pip_install("torch", "huggingface")
    )

    app = modal.App(image=image)

Copy

As shown above, you also need to use a [Modal Secret](/docs/guide/secrets)
containing the environment variables `AWS_ACCESS_KEY_ID`,
`AWS_SECRET_ACCESS_KEY`, and `AWS_REGION`. The AWS IAM user account associated
with those keys must have access to the private registry you want to access.

The user needs to have the following read-only policies:

    {
      "Version": "2012-10-17",
      "Statement": [
        {
          "Action": ["ecr:GetAuthorizationToken"],
          "Effect": "Allow",
          "Resource": "*"
        },
        {
          "Effect": "Allow",
          "Action": [
            "ecr:BatchCheckLayerAvailability",
            "ecr:GetDownloadUrlForLayer",
            "ecr:GetRepositoryPolicy",
            "ecr:DescribeRepositories",
            "ecr:ListImages",
            "ecr:DescribeImages",
            "ecr:BatchGetImage",
            "ecr:GetLifecyclePolicy",
            "ecr:GetLifecyclePolicyPreview",
            "ecr:ListTagsForResource",
            "ecr:DescribeImageScanFindings"
          ],
          "Resource": "<MY-REGISTRY-ARN>"
        }
      ]
    }

Copy

You can use the IAM configuration above as a template for creating an IAM
user. You can then [generate an access
key](https://aws.amazon.com/premiumsupport/knowledge-center/create-access-
key/) and create a Modal Secret using the AWS integration option. Modal will
use your access keys to generate an ephemeral ECR token. That token is only
used to pull image layers at the time a new image is built. We don’t store
this token but will cache the image once it has been pulled.

Images on ECR must be private and follow [image configuration
requirements](/docs/reference/modal.Image#from_aws_ecr).

## Google Artifact Registry and Google Container Registry

For further detail on how to pull images from Google’s image registries, see
[`modal.Image.from_gcp_artifact_registry`](/docs/reference/modal.Image#from_gcp_artifact_registry).

Private registriesDocker Hub (Private)Elastic Container Registry (ECR)Google
Artifact Registry and Google Container Registry

See it in action

[Registry image for Algolia indexing](/docs/examples/algolia_indexer)
