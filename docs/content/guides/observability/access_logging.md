---
title: Trace Access Logging
menuTitle: Trace Access Logging
description: Guide on Gloo Mesh's trace logging features.
weight: 30
---

{{% notice note %}} Gloo Mesh Enterprise is required for this feature. {{% /notice %}}

## Background

Distributed microservices can be difficult to debug because there are many hops in a potential graph of calls. A service mesh can be used to help observe overall health of a system and can indicate when things are degraded or unhealthy. Things like metric collection, tracing, and alerting can be used to identify problem areas. Unfortunately there are times when an SRE or developer must dig into the details of a service graph to understand more deeploy where things are going wrong.

With Gloo Mesh, we can use trace logging to see all of the access logging through a particular call graph to further pinpoint problem areas. With trace logging, we can enable detailed access logging for specific workloads and specify pattern matching to collect access logs generated from specific requests and centralize them for further evaluation. 

## How does this relate to platform logging?

Gloo Mesh trace logging is not a replacement for proper platform logging with something like Fluentd/Elastic or Splunk. You should continue to use those tools for general platform logging. 

Access logging in the mesh is typically not enabled for all workloads. Generally it's enabled for edge gateways or on-demand for services in a graph. Gloo Mesh trace logging can be enabled on demand for certain workloads, used to debug an issue, and then turned off. It's intended to be a convenience feature and NOT to replace platform logging. 

## Before you begin

This guide assumes the following:

  * Gloo Mesh Enterprise is [installed in relay mode and running on the `cluster-1`]({{% versioned_link_path fromRoot="/setup/installation/enterprise_installation/" %}})
  * `gloo-mesh` is the installation namespace for Gloo Mesh
  * `enterprise-networking` is deployed on `cluster-1` in the `gloo-mesh` namespace and exposes its gRPC server on port 9900
  * `enterprise-agent` is deployed on both clusters and exposes its gRPC server on port 9977
  * Both `cluster-1` and `cluster-2` are [registered with Gloo Mesh]({{% versioned_link_path fromRoot="/guides/#two-registered-clusters" %}})
  * Istio is [installed on both clusters]({{% versioned_link_path fromRoot="/guides/installing_istio" %}}) clusters
  * `istio-system` is the root namespace for both Istio deployments
  * The `bookinfo` app is [installed into the two clusters]({{% versioned_link_path fromRoot="/guides/#bookinfo-deployed-on-two-clusters" %}}) under the `bookinfo` namespace
  * the following environment variables are set:
    ```shell
    CONTEXT_1=cluster_1's_context
    CONTEXT_2=cluster_2's_context
    ```

## Istio Configuration

Before we begin, we need to ensure that our Istio deployments in
both `cluster-1` and `cluster-2` have the necessary configuration for
Gloo Mesh access logging. View the Istio ConfigMap with the following command:

```shell
kubectl --context $CONTEXT_1 -n istio-system get configmap istio -oyaml
```

Ensure that the following highlighted entries exist in the ConfigMap:

{{< highlight yaml "hl_lines=5-7" >}}
data:
  mesh:
    defaultConfig:
      envoyAccessLogService:
        address: enterprise-agent.gloo-mesh:9977
      proxyMetadata:
        GLOO_MESH_CLUSTER_NAME: your-gloo-mesh-registered-cluster-name
{{< /highlight >}}

If you updated the ConfigMap, you must restart existing Istio injected workloads in order
for their sidecars to pick up the new config.

The `GLOO_MESH_CLUSTER_NAME` value is used to annotate the Gloo Mesh cluster name when emitting
access logs, which is then used by Gloo Mesh to correlate the envoy proxy to a discovered workload.

## AccessLogRecord CRD

Gloo Mesh uses the [AccessLogRecord]({{% versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.enterprise.observability.v1alpha1.access_logging/" %}})
custom resource to configure access log collection. The API allows specifying the workloads
for which to enable collection as well as request/response level filter criteria (for only emitting a filtered subset of all access logs).

For demonstration purposes let's create the following object:

{{< tabs >}}
{{< tab name="YAML file" codelang="yaml">}}
apiVersion: observability.enterprise.mesh.gloo.solo.io/v1
kind: AccessLogRecord
metadata:
  name: access-log-all
  namespace: gloo-mesh
spec:
  filters:
    - headerMatcher:
        name: foo
        value: bar
{{< /tab >}}
{{< tab name="CLI inline" codelang="shell" >}}
kubectl apply --context $CONTEXT_1 -f - << EOF
apiVersion: observability.enterprise.mesh.gloo.solo.io/v1
kind: AccessLogRecord
metadata:
  name: access-log-all
  namespace: gloo-mesh
spec:
  filters:
  - headerMatcher:
      name: foo
      value: bar
EOF
{{< /tab >}}
{{< /tabs >}}

This will enable access log collection for all workloads in both clusters, but only
for requests containing the header `"foo": "bar"`.

## Retrieving Access Logs

Let's first generate some access logs by making requests in both clusters:

```shell
kubectl --context $CONTEXT_1 -n bookinfo exec -it deploy/ratings-v1 -c ratings --  curl -H "foo: bar" -v reviews:9080/reviews/1
```

```shell
kubectl --context $CONTEXT_2 -n bookinfo exec -it deploy/ratings-v1 -c ratings --  curl -H "foo: bar" -v reviews:9080/reviews/1
```

Assuming the access logs were collected successfully, we can now view them via the Gloo Mesh UI.

### Gloo Mesh UI

{{% notice note %}} Gloo Mesh Enterprise is required for this feature. {{% /notice %}}

In the Gloo Mesh UI, the access logs will be displayed when clicking into
a workload and clicking its Access Logs tab.

## Debugging

Because access logs provide detailed contextual information at the granularity of 
individual networking requests and responses, they are a valuable tool for debugging.
To showcase this, we will contrive a network error and see how access logs can help
in diagnosing the problem.

First ensure that the Gloo Mesh settings object disables Istio mTLS. This will allow
us to modify mTLS settings for specific Destinations.

{{< highlight yaml "hl_lines=10" >}}
apiVersion: settings.mesh.gloo.solo.io/v1
kind: Settings
metadata:
  name: settings
  namespace: gloo-mesh
spec:
  ...
  mtls:
    istio:
      tlsMode: DISABLE
{{< /highlight >}}

Next, create the following Istio DestinationRule which is intentionally erroroneous,
the referenced TLS secret data does not exist.

{{< tabs >}}
{{< tab name="YAML file" codelang="yaml">}}
apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: ratings
  namespace: bookinfo
spec:
  host: ratings.bookinfo.svc.cluster.local
  trafficPolicy:
    tls:
      mode: MUTUAL
      # these files do not exist
      clientCertificate: /etc/certs/myclientcert.pem
      privateKey: /etc/certs/client_private_key.pem
      caCertificates: /etc/certs/rootcacerts.pem
{{< /tab >}}
{{< tab name="CLI inline" codelang="shell" >}}
kubectl apply --context $CONTEXT_2 -f - << EOF
apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: ratings
  namespace: bookinfo
spec:
  host: ratings.bookinfo.svc.cluster.local
    trafficPolicy:
      tls:
        mode: MUTUAL
        # these files do not exist
        clientCertificate: /etc/certs/myclientcert.pem
        privateKey: /etc/certs/client_private_key.pem
        caCertificates: /etc/certs/rootcacerts.pem
EOF
{{< /tab >}}
{{< /tabs >}}

Sending a request from the `productpage` pod to the ratings Destination should yield 
the following access log:

{{< highlight json "hl_lines=3" >}}
{
      ...
        "upstreamTransportFailureReason": "TLS error: Secret is not supplied by SDS",
        "routeName": "default",
        "downstreamDirectRemoteAddress": {
          "socketAddress": {
            "address": "192.168.2.14",
            "portValue": 52836
          }
        }
      ...
    }
  }
}
{{< /highlight >}}

Envoy access logs contain a highly detailed information, the details of which can be found
in the [envoy access log documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage).
