---
canonical: https://grafana.com/docs/alloy/latest/reference/components/discovery/discovery.dockerswarm/
aliases:
  - ../discovery.dockerswarm/ # /docs/alloy/latest/reference/components/discovery.dockerswarm/
description: Learn about discovery.dockerswarm
labels:
  stage: general-availability
  products:
    - oss
title: discovery.dockerswarm
---

# `discovery.dockerswarm`

`discovery.dockerswarm` allows you to retrieve scrape targets from [Docker Swarm](https://docs.docker.com/engine/swarm/key-concepts/).

## Usage

```alloy
discovery.dockerswarm "<LABEL>" {
  host = "<DOCKER_DAEMON_HOST>"
  role = "<SWARM_ROLE>"
}
```

## Arguments

You can use the following arguments with `discovery.dockerswarm`:

| Name                     | Type                | Description                                                                                                                   | Default | Required |
| ------------------------ | ------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ------- | -------- |
| `host`                   | `string`            | Address of the Docker daemon.                                                                                                 |         | yes      |
| `role`                   | `string`            | Role of the targets to retrieve. Must be `services`, `tasks`, or `nodes`.                                                     |         | yes      |
| `bearer_token_file`      | `string`            | File containing a bearer token to authenticate with.                                                                          |         | no       |
| `bearer_token`           | `secret`            | Bearer token to authenticate with.                                                                                            |         | no       |
| `enable_http2`           | `bool`              | Whether HTTP2 is supported for requests.                                                                                      | `true`  | no       |
| `follow_redirects`       | `bool`              | Whether redirects returned by the server should be followed.                                                                  | `true`  | no       |
| `http_headers`           | `map(list(secret))` | Custom HTTP headers to be sent along with each request. The map key is the header name.                                       |         | no       |
| `no_proxy`               | `string`            | Comma-separated list of IP addresses, CIDR notations, and domain names to exclude from proxying.                              |         | no       |
| `port`                   | `number`            | The port to scrape metrics from, when `role` is nodes, and for discovered tasks and services that don't have published ports. | `80`    | no       |
| `proxy_connect_header`   | `map(list(secret))` | Specifies headers to send to proxies during CONNECT requests.                                                                 |         | no       |
| `proxy_from_environment` | `bool`              | Use the proxy URL indicated by environment variables.                                                                         | `false` | no       |
| `proxy_url`              | `string`            | HTTP proxy to send requests through.                                                                                          |         | no       |
| `refresh_interval`       | `duration`          | Interval at which to refresh the list of targets.                                                                             | `"60s"` | no       |

At most, one of the following can be provided:

* [`authorization`][authorization] block
* [`basic_auth`][basic_auth] block
* [`bearer_token_file`][arguments] argument
* [`bearer_token`][arguments] argument
* [`oauth2`][oauth2] block

{{< docs/shared lookup="reference/components/http-client-proxy-config-description.md" source="alloy" version="<ALLOY_VERSION>" >}}

[arguments]: #arguments

## Blocks

You can use the following blocks with `discovery.dockerswarm`:

| Block                                 | Description                                                                        | Required |
| ------------------------------------- | ---------------------------------------------------------------------------------- | -------- |
| [`authorization`][authorization]      | Configure generic authorization to the endpoint.                                   | no       |
| [`basic_auth`][basic_auth]            | Configure `basic_auth` for authenticating to the endpoint.                         | no       |
| [`filter`][filter]                    | Optional filter to limit the discovery process to a subset of available resources. | no       |
| [`oauth2`][oauth2]                    | Configure OAuth 2.0 for authenticating to the endpoint.                            | no       |
| `oauth2` > [`tls_config`][tls_config] | Configure TLS settings for connecting to the endpoint.                             | no       |
| [`tls_config`][tls_config]            | Configure TLS settings for connecting to the endpoint.                             | no       |

The `>` symbol indicates deeper levels of nesting.
For example, `oauth2 > tls_config` refers to a `tls_config` block defined inside an `oauth2` block.

[filter]: #filter
[basic_auth]: #basic_auth
[authorization]: #authorization
[oauth2]: #oauth2
[tls_config]: #tls_config

### `authorization`

The `authorization` block configures generic authorization to the endpoint.

{{< docs/shared lookup="reference/components/authorization-block.md" source="alloy" version="<ALLOY_VERSION>" >}}

### `basic_auth`

The `basic_auth` block configures basic authentication to the endpoint.

{{< docs/shared lookup="reference/components/basic-auth-block.md" source="alloy" version="<ALLOY_VERSION>" >}}

### `filter`

The `filter` block limits the discovery process to a subset of available resources.
You can define multiple `filter` blocks within the `discovery.dockerswarm` block.
The list of available filters depends on the `role`:

* [nodes filters](https://docs.docker.com/engine/api/v1.40/#operation/NodeList)
* [services filters](https://docs.docker.com/engine/api/v1.40/#operation/ServiceList)
* [tasks filters](https://docs.docker.com/engine/api/v1.40/#operation/TaskList)

You can use the following arguments to configure a filter.

| Name     | Type           | Description                                | Default | Required |
| -------- | -------------- | ------------------------------------------ | ------- | -------- |
| `name`   | `string`       | Name of the filter.                        |         | yes      |
| `values` | `list(string)` | List of values associated with the filter. |         | yes      |

### `oauth2`

The `oauth2` block configures OAuth 2.0 authentication to the endpoint.

{{< docs/shared lookup="reference/components/oauth2-block.md" source="alloy" version="<ALLOY_VERSION>" >}}

### `tls_config`

The `tls_config` block configures TLS settings for connecting to the endpoint.

{{< docs/shared lookup="reference/components/tls-config-block.md" source="alloy" version="<ALLOY_VERSION>" >}}

## Exported fields

The following fields are exported and can be referenced by other components:

| Name      | Type                | Description                               |
| --------- | ------------------- | ----------------------------------------- |
| `targets` | `list(map(string))` | The set of targets discovered from Swarm. |

## Roles

The `role` attribute decides the role of the targets to retrieve.

### `services`

The `services` role discovers all [Swarm services](https://docs.docker.com/engine/swarm/key-concepts/#services-and-tasks) and exposes their ports as targets.
For each published port of a service, a single target is generated.
If a service has no published ports, a target per service is created using the `port` attribute defined in the arguments.

Available meta labels:

* `__meta_dockerswarm_network_id`: The ID of the network.
* `__meta_dockerswarm_network_ingress`: Whether the network is ingress.
* `__meta_dockerswarm_network_internal`: Whether the network is internal.
* `__meta_dockerswarm_network_label_<labelname>`: Each label of the network.
* `__meta_dockerswarm_network_name`: The name of the network.
* `__meta_dockerswarm_network_scope`: The scope of the network.
* `__meta_dockerswarm_service_endpoint_port_name`: The name of the endpoint port, if available.
* `__meta_dockerswarm_service_endpoint_port_publish_mode`: The publish mode of the endpoint port.
* `__meta_dockerswarm_service_id`: The ID of the service.
* `__meta_dockerswarm_service_label_<labelname>`: Each label of the service.
* `__meta_dockerswarm_service_mode`: The mode of the service.
* `__meta_dockerswarm_service_name`: The name of the service.
* `__meta_dockerswarm_service_task_container_hostname`: The container hostname of the target, if available.
* `__meta_dockerswarm_service_task_container_image`: The container image of the target.
* `__meta_dockerswarm_service_updating_status`: The status of the service, if available.

### `tasks`

The `tasks` role discovers all [Swarm tasks](https://docs.docker.com/engine/swarm/key-concepts/#services-and-tasks) and exposes their ports as targets.
For each published port of a task, a single target is generated.
If a task has no published ports, a target per task is created using the `port` attribute defined in the arguments.

Available meta labels:

* `__meta_dockerswarm_container_label_<labelname>`: Each label of the container.
* `__meta_dockerswarm_network_id`: The ID of the network.
* `__meta_dockerswarm_network_ingress`: Whether the network is ingress.
* `__meta_dockerswarm_network_internal`: Whether the network is internal.
* `__meta_dockerswarm_network_label_<labelname>`: Each label of the network.
* `__meta_dockerswarm_network_label`: Each label of the network.
* `__meta_dockerswarm_network_name`: The name of the network.
* `__meta_dockerswarm_network_scope`: The scope of the network.
* `__meta_dockerswarm_node_address`: The address of the node.
* `__meta_dockerswarm_node_availability`: The availability of the node.
* `__meta_dockerswarm_node_hostname`: The hostname of the node.
* `__meta_dockerswarm_node_id`: The ID of the node.
* `__meta_dockerswarm_node_label_<labelname>`: Each label of the node.
* `__meta_dockerswarm_node_platform_architecture`: The architecture of the node.
* `__meta_dockerswarm_node_platform_os`: The operating system of the node.
* `__meta_dockerswarm_node_role`: The role of the node.
* `__meta_dockerswarm_node_status`: The status of the node.
* `__meta_dockerswarm_service_id`: The ID of the service.
* `__meta_dockerswarm_service_label_<labelname>`: Each label of the service.
* `__meta_dockerswarm_service_mode`: The mode of the service.
* `__meta_dockerswarm_service_name`: The name of the service.
* `__meta_dockerswarm_task_container_id`: The container ID of the task.
* `__meta_dockerswarm_task_desired_state`: The desired state of the task.
* `__meta_dockerswarm_task_id`: The ID of the task.
* `__meta_dockerswarm_task_port_publish_mode`: The publish mode of the task port.
* `__meta_dockerswarm_task_slot`: The slot of the task.
* `__meta_dockerswarm_task_state`: The state of the task.

The `__meta_dockerswarm_network_*` meta labels aren't populated for ports which are published with mode=host.

### `nodes`

The `nodes` role is used to discover [Swarm nodes](https://docs.docker.com/engine/swarm/key-concepts/#nodes).

Available meta labels:

* `__meta_dockerswarm_node_address`: The address of the node.
* `__meta_dockerswarm_node_availability`: The availability of the node.
* `__meta_dockerswarm_node_engine_version`: The version of the node engine.
* `__meta_dockerswarm_node_hostname`: The hostname of the node.
* `__meta_dockerswarm_node_id`: The ID of the node.
* `__meta_dockerswarm_node_label_<labelname>`: Each label of the node.
* `__meta_dockerswarm_node_manager_address`: The address of the manager component of the node.
* `__meta_dockerswarm_node_manager_leader`: The leadership status of the manager component of the node (true or false).
* `__meta_dockerswarm_node_manager_reachability`: The reachability of the manager component of the node.
* `__meta_dockerswarm_node_platform_architecture`: The architecture of the node.
* `__meta_dockerswarm_node_platform_os`: The operating system of the node.
* `__meta_dockerswarm_node_role`: The role of the node.
* `__meta_dockerswarm_node_status`: The status of the node.

## Component health

`discovery.dockerswarm` is only reported as unhealthy when given an invalid configuration.
In those cases, exported fields retain their last healthy values.

## Debug information

`discovery.dockerswarm` doesn't expose any component-specific debug information.

## Debug metrics

`discovery.dockerswarm` doesn't expose any component-specific debug metrics.

## Example

This example discovers targets from Docker Swarm tasks:

```alloy
discovery.dockerswarm "example" {
  host = "unix:///var/run/docker.sock"
  role = "tasks"

  filter {
    name = "id"
    values = ["0kzzo1i0y4jz6027t0k7aezc7"]
  }

  filter {
    name = "desired-state"
    values = ["running", "accepted"]
  }
}

prometheus.scrape "demo" {
  targets    = discovery.dockerswarm.example.targets
  forward_to = [prometheus.remote_write.demo.receiver]
}

prometheus.remote_write "demo" {
  endpoint {
    url = "<PROMETHEUS_REMOTE_WRITE_URL>"

    basic_auth {
      username = "<USERNAME>"
      password = "<PASSWORD>"
    }
  }
}
```

Replace the following:

* _`<PROMETHEUS_REMOTE_WRITE_URL>`_: The URL of the Prometheus remote_write-compatible server to send metrics to.
* _`<USERNAME>`_: The username to use for authentication to the `remote_write` API.
* _`<PASSWORD>`_: The password to use for authentication to the `remote_write` API.

<!-- START GENERATED COMPATIBLE COMPONENTS -->

## Compatible components

`discovery.dockerswarm` has exports that can be consumed by the following components:

- Components that consume [Targets](../../../compatibility/#targets-consumers)

{{< admonition type="note" >}}
Connecting some components may not be sensible or components may require further configuration to make the connection work correctly.
Refer to the linked documentation for more details.
{{< /admonition >}}

<!-- END GENERATED COMPATIBLE COMPONENTS -->
