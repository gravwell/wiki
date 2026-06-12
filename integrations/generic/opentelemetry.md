# OpenTelemetry

:::{csv-table}
:align: left
:width: 45%
:widths: 15, 25
**Integration Details**
    Ingester, [HTTP Ingester](/ingesters/http)
         Kit, [OpenTelemetry Kit](https://github.com/gravwell/kits/tree/main/opentelemetry)
:::

## OpenTelemetry Configuration

OpenTelemetry uses Collectors to forward data from your application or infrastructure directly into Gravwell. While this guide demonstrates using Kubernetes it could be adapted to route logs from any OpenTelemetry compatible service.

### Example: Simple OpenTelemetry Example

To send data to Gravwell, modify your OpenTelemetry Collector configuration file to include the Gravwell HTTP ingester into the `exporters` section.

```YAML
exporters:
  otlphttp:
    endpoint: http://192.168.3.50:8080 # Change this line to point to your gravwell instance
    encoding: json
    tls:
      insecure: true
```

Next, update your active data pipelines under the `service` section to route your data through the new `otlphttp` exporer:

```YAML
service:
  pipelines:
    traces:
      receivers: [otlp, jaeger, zipkin]
      processors: [batch]
      exporters: [otlp_http, debug]

    metrics:
      receivers: [otlp, prometheus]
      processors: [batch]
      exporters: [otlp_http, debug]

    logs:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp_http, debug]

```

```{note}
You can validate your Collector config with:
`otelcol validate --config=customconfig.yaml`
```

### Example: Kubernetes Manifest File

This example provides a native Kubernetes manifest file that sets up an OpenTelemetry Collector to automatically ingest logs and performance metrics from your Kubernetes cluster. To route this data to your environment, you must modify the following line in the manifest to point to your Gravwell HTTP Ingester:
```
endpoint: http://192.168.3.50:8080
```

This manifest utilizes Custom Resource Definitions from the OpenTelemetry Operator. The following commands provide an example that could be ran on your master node to install the operator via Helm to deploy the collector.

```
# 1. Add the OpenTelemetry Helm repository
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update

# 2. Install the OpenTelemetry Operator chart
helm install my-otel-operator open-telemetry/opentelemetry-operator --set admissionWebhooks.certManager.enabled=false

# 3. Deploy the custom Collector configuration
kubectl apply -f otel_gravwell_collector.yml
```

```YAML
apiVersion: v1
kind: ServiceAccount
metadata:
  name: otelcol
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: otelcol
rules:
  - apiGroups: [""]
    resources:
      - nodes
      - nodes/stats
      - nodes/proxy
      - pods
      - namespaces
      - services
      - endpoints
    verbs: ["get", "list", "watch"]
  - apiGroups: ["apps"]
    resources:
      - deployments
      - daemonsets
      - replicasets
      - statefulsets
    verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: otelcol
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: otelcol
subjects:
  - kind: ServiceAccount
    name: otelcol
    namespace: default
---
apiVersion: opentelemetry.io/v1beta1
kind: OpenTelemetryCollector
metadata:
  name: otel
  namespace: default
spec:
  image: otel/opentelemetry-collector-contrib:0.112.0
  mode: daemonset
  serviceAccount: otelcol
 
  # hostNetwork is required for the `from: connection` pod_association in
  # k8sattributes to resolve the real client pod IP. Without it, all incoming
  # OTLP traffic looks like it came from the CNI's masqueraded address.
  hostNetwork: true
  dnsPolicy: ClusterFirstWithHostNet
 
  securityContext:
    runAsUser: 0
 
  env:
    - name: K8S_NODE_NAME
      valueFrom:
        fieldRef:
          fieldPath: spec.nodeName
 
  volumes:
    - name: varlogpods
      hostPath:
        path: /var/log/pods
    - name: varlibdockercontainers
      hostPath:
        path: /var/lib/docker/containers
  volumeMounts:
    - name: varlogpods
      mountPath: /var/log/pods
      readOnly: true
    - name: varlibdockercontainers
      mountPath: /var/lib/docker/containers
      readOnly: true
 
  config:
    receivers:
      otlp:
        protocols:
          grpc:
            endpoint: 0.0.0.0:4317
          http:
            endpoint: 0.0.0.0:4318
 
      kubeletstats:
        collection_interval: 20s
        auth_type: serviceAccount
        # Explicit https:// scheme is required by recent versions
        endpoint: https://${env:K8S_NODE_NAME}:10250
        insecure_skip_verify: true
        metric_groups: [node, pod, container]
 
      filelog:
        include: [/var/log/pods/*/*/*.log]
        # Exclude this collector's own pods to avoid a feedback loop.
        # The operator names DaemonSet pods `<cr-name>-collector-<hash>`,
        # so for a CR named `otel` in namespace `default` the pod dirs are:
        #   /var/log/pods/default_otel-collector-<hash>_<uid>/...
        exclude:
          - /var/log/pods/default_otel-collector-*/*/*.log
        include_file_path: true
        start_at: end
        operators:
          - type: container
            id: container-parser
 
    processors:
      batch: {}
 
      k8sattributes:
        auth_type: serviceAccount
        passthrough: false
        extract:
          metadata:
            - k8s.pod.name
            - k8s.pod.uid
            - k8s.namespace.name
            - k8s.node.name
            - k8s.deployment.name
        pod_association:
          - sources:
              - from: resource_attribute
                name: k8s.pod.uid
          - sources:
              - from: connection
 
    exporters:
      otlphttp:
        endpoint: http://192.168.3.50:8080 # Change this line to point to your gravwell instance
        encoding: json
        tls:
          insecure: true
 
      debug:
        verbosity: basic
 
    service:
      pipelines:
        metrics:
          receivers: [otlp, kubeletstats]
          processors:
            - k8sattributes
            - batch
          exporters: [otlphttp, debug]
        logs:
          receivers: [otlp, filelog]
          processors: [k8sattributes, batch]
          # NOTE: do NOT add `debug` here unless your filelog `exclude`
          # pattern definitely matches this collector's pod path, or you
          # will create a feedback loop.
          exporters: [otlphttp]
```

## Gravwell Configuration

### Gravwell Storage Well Configuration

Setup the well configuration in your Gravwell indexers.

**Sample well config:**  
Create or edit: `/opt/gravwell/etc/gravwell.conf.d/opentelemetry-well.conf`
```ini
[Storage-Well "opentelemetry"]
    Location=/opt/gravwell/storage/opentelemetry
    Tags=opentelemetry*
```

### Gravwell Ingester Configuration: HTTP
**Sample OpenTelemetry config:**  
Create or edit: `/opt/gravwell/etc/INGESTER_OpenTelemetry/opentelemetry.conf`
```ini
[OpenTelemetry-Logs-Listener "kubelogs"]
    URL="/v1/logs"
    Tag-Name=k8s-logs

[OpenTelemetry-Metrics-Listener "kubemetrics"]
    URL="/v1/metrics"
    Tag-Name=k8s-metricsQ
```

```{note}
Remember to restart the service to apply the new config:
`sudo systemctl restart INGESTER_OpenTelemetry.service`
```