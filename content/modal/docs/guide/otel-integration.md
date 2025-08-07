* * *

Copy page

# Connecting Modal to your OpenTelemetry Provider

You can export Modal logs to your
[OpenTelemetry](https://opentelemetry.io/docs/what-is-opentelemetry/) provider
using the Modal OpenTelemetry integration. This integration is compatible with
any observability provider that supports the OpenTelemetry HTTP APIs.

## What this integration does

This integration allows you to:

  1. Export Modal audit logs to your provider
  2. Export Modal function logs to your provider
  3. Export container metrics to your provider

## Metrics

The Modal OpenTelemetry Integration will forward the following metrics to your
provider:

  * `modal.cpu.utilization`
  * `modal.memory.utilization`
  * `modal.gpu.memory.utilization`
  * `modal.gpu.compute.utilization`

These metrics are tagged with `container_id`, `environment_name`, and
`workspace_name`.

## Installing the integration

  1. Find out the endpoint URL for your OpenTelemetry provider. This is the URL that the Modal integration will send logs to. Note that this should be the base URL of the OpenTelemetry provider, and not a specific endpoint. For example, for the [US New Relic instance](https://docs.newrelic.com/docs/opentelemetry/best-practices/opentelemetry-otlp/#configure-endpoint-port-protocol), the endpoint URL is `https://otlp.nr-data.net`, not `https://otlp.nr-data.net/v1/logs`.
  2. Find out the API key or other authentication method required to send logs to your OpenTelemetry provider. This is the key that the Modal integration will use to authenticate with your provider. Modal can provide any key/value HTTP header pairs. For example, for [New Relic](https://docs.newrelic.com/docs/opentelemetry/best-practices/opentelemetry-otlp/#api-key), the header is `api-key`.
  3. Create a new OpenTelemetry Secret in Modal with one key per header. These keys should be prefixed with `OTEL_HEADER_`, followed by the name of the header. The value of this key should be the value of the header. For example, for New Relic, an example Secret might look like `OTEL_HEADER_api-key: YOUR_API_KEY`. If you use the OpenTelemetry Secret template, this will be pre-filled for you.
  4. Navigate to the [Modal metrics settings page](http://modal.com/settings/metrics) and configure the OpenTelemetry push URL from step 1 and the Secret from step 3.
  5. Save your changes and use the test button to confirm that logs are being sent to your provider. If it’s all working, you should see a `Hello from Modal! 🚀` log from the `modal.test_logs` service.

## Allowlisting the integration’s IPs

The integration uses a set of static IP addresses (subject to change) to send
data to your OpenTelemetry provider:

    3.215.65.235
    3.219.40.38
    13.219.36.96
    18.206.3.184
    35.169.34.255
    44.208.153.216
    52.86.93.233
    54.198.254.114
    54.208.217.246
    204.236.196.209

Copy

## Uninstalling the integration

Once the integration is uninstalled, all logs will stop being sent to your
provider.

  1. Navigate to the [Modal metrics settings page](http://modal.com/settings/metrics) and disable the OpenTelemetry integration.

Connecting Modal to your OpenTelemetry ProviderWhat this integration
doesMetricsInstalling the integrationAllowlisting the integration’s
IPsUninstalling the integration
