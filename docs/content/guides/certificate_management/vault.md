---
title: Vault PKI Integration
menuTitle: Vault PKI Integration
description: Guide to using Vault with Gloo Mesh Enterprise to manage Istio certificates
weight: 70
---

{{% notice note %}} Gloo Mesh Enterprise is required for this feature. {{% /notice %}}

[Vault](https://github.com/hashicorp/vault) is a popular open source secret management tool. One of its many use cases is PKI (Private Key Infrastructure). Vault allows for easy and secure storage of our private keys, as well as generation of new leaf/intermediary certificates. 

For multi-cluster traffic routing, we need to establish a shared root of trust using a single root CA and intermediate CA that is signed by the same root CA for each cluster. This guide will show you how to configure Istio and Gloo Mesh to use Vault to store our root CA and generate intermediate CA to be used by Istio on each cluster to sign its workload certificates.

In addition to using Vault for the intermediate CA, this guide will also explore the added security benefits of using Gloo Mesh Enterprise with Vault. The Gloo Mesh Enterprise integration with Vault uses the `istiod-agent`, which runs as a sidecar to the `istiod` pod, and communicates with Vault to request private keys and sign certificates. This allows Gloo Mesh to load the private key directly into the pod filesystem, thereby allowing for an added layer of security by not saving the key to `etcd` (or any permanent storage). 

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
CONTEXT_1="cluster_1's_context"
CONTEXT_2="cluster_2's_context"
```

## Installing Vault

{{% notice note %}} Installing Vault is an optional step. An existing Vault deployment may be used if you already have access to one. {{% /notice %}}

1. Install Vault using helm

```shell
helm repo add hashicorp https://helm.releases.hashicorp.com
```

2. Generate root-cert and key for Vault

```shell
openssl req -new -newkey rsa:4096 -x509 -sha256 \
    -days 3650 -nodes -out root-cert.pem -keyout root-key.pem \
    -subj "/O=my-org"
```

3. Let's install vault, and add our root-ca to each deployment

```shell
for cluster in ${CONTEXT_1} ${CONTEXT_2}; do

  # For more info on vault in kubernetes, please see: https://learn.hashicorp.com/tutorials/vault/kubernetes-cert-manager

  # install vault in dev mode
  helm install -n vault  vault hashicorp/vault --set "injector.enabled=false" --set "server.dev.enabled=true" --kube-context="${cluster}" --create-namespace

  # Wait for vault to come up, can't use kubectl rollout because it's a stateful set without rolling deployment
  kubectl --context="${cluster}" wait --for=condition=Ready -n vault pod/vault-0

  # Enable vault kubernetes Auth
  kubectl --context="${cluster}" exec -n vault vault-0 -- /bin/sh -c 'vault auth enable kubernetes'

  # Set the kubernetes Auth config for vault to the mounted token
  kubectl --context="${cluster}" exec -n vault vault-0 -- /bin/sh -c 'vault write auth/kubernetes/config \
    token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
    kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443" \
    kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt'

  # Bind the istiod service account to the pki policy below
  kubectl --context="${cluster}" exec -n vault vault-0 -- /bin/sh -c 'vault write auth/kubernetes/role/gen-int-ca-istio \
    bound_service_account_names=istiod-service-account \
    bound_service_account_namespaces=istio-system \
    policies=gen-int-ca-istio \
    ttl=2400h'

  # Initialize vault PKI
  kubectl --context="${cluster}" exec -n vault vault-0 -- /bin/sh -c 'vault secrets enable pki'

  # set the vault CA to our pem_bundle
  kubectl --context="${cluster}" exec -n vault vault-0 -- /bin/sh -c "vault write -format=json pki/config/ca pem_bundle=\"$(cat root-key.pem root-cert.pem)\""

  # Initialize vault intermediate path
  kubectl --context="${cluster}" exec -n vault vault-0 -- /bin/sh -c 'vault secrets enable -path pki_int pki'

  # Set the policy for intermediate cert path
  kubectl --context="${cluster}" exec -n vault vault-0 -- /bin/sh -c 'vault policy write gen-int-ca-istio - <<EOF
path "pki_int/*" {
capabilities = ["create", "read", "update", "delete", "list"]
}
path "pki/cert/ca" {
capabilities = ["read"]
}
path "pki/root/sign-intermediate" {
capabilities = ["create", "read", "update", "list"]
}
EOF'

done
```

## Enabling Vault as an intermediate CA provider

Now we need to federate our 2 meshes together using Vault to federate identity. To do this, we will need to create/edit a `VirtualMesh` with the new Vault shared mTLS config. The highlighted lines are the new config which the remote clusters will use to authenticate, and communicate with vault. There are brief explanations as comments below, but for further information see the [API docs]({{% versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.certificates.v1.vault_ca/" %}})

{{< highlight yaml "hl_lines=10-20" >}}
apiVersion: networking.mesh.gloo.solo.io/v1
kind: VirtualMesh
metadata:
  name: virtual-mesh
  namespace: gloo-mesh
spec:
  mtlsConfig:
    # Note: Do NOT use this autoRestartPods setting in production!!
    autoRestartPods: true
    shared:
      intermediateCertificateAuthority:
        vault:
          # Vault path to the CA endpoint
          caPath: "pki/root/sign-intermediate"
          # Vault path to the CSR endpoint
          csrPath: "pki_int/intermediate/generate/exported"
          # Vault server endpoint
          server: "http://vault.vault:8200"
          # Auth mechanism to use
          kubernetesAuth:
            role: "gen-int-ca-istio"
  federation:
    # federate all Destinations to all external meshes
    selectors:
    - {}
  meshes:
  - name: istiod-istio-system-cluster-1
    namespace: gloo-mesh
  - name: istiod-istio-system-cluster-2
    namespace: gloo-mesh
{{< /highlight >}}

Run the following command to apply this `VirtualMesh` to the mgmt-cluster.

```shell
cat << EOF | kubectl apply --context=${CONTEXT_1} -f -
apiVersion: networking.mesh.gloo.solo.io/v1
kind: VirtualMesh
metadata:
  name: virtual-mesh
  namespace: gloo-mesh
spec:
  mtlsConfig:
    # Note: Do NOT use this autoRestartPods setting in production!!
    autoRestartPods: true
    shared:
      intermediateCertificateAuthority:
        vault:
          # Vault path to the CA endpoint
          caPath: "pki/root/sign-intermediate"
          # Vault path to the CSR endpoint
          csrPath: "pki_int/intermediate/generate/exported"
          # Vault server endpoint
          server: "http://vault.vault:8200"
          # Auth mechanism to use
          kubernetesAuth:
            role: "gen-int-ca-istio"
  federation:
    # federate all Destinations to all external meshes
    selectors:
    - {}
  meshes:
  - name: istiod-istio-system-cluster-1
    namespace: gloo-mesh
  - name: istiod-istio-system-cluster-2
    namespace: gloo-mesh
EOF
```

## Updating RBAC

The new istio-agent sidecar (we'll install in the next step) will need to read and modify Gloo Mesh resources. To enable the necessary RBAC permissions, we are going to update our `enterprise-agent` helm release on both clusters.

If your  Enterprise Agents were installed via helm, and those manifests are applied using GitOps then add the following values to your `values.yaml` file.

```yaml
istiodSidecar:
  createRoleBinding: true
```

Otherwise, run the following to update the currently deployed agent helm releases.

```shell
for cluster in ${CONTEXT_1} ${CONTEXT_2}; do
  helm get values -n gloo-mesh enterprise-agent --kube-context="${cluster}" > $cluster-values.yaml
  echo "istiodSidecar:
  createRoleBinding: true" > $cluster-values.yaml
  helm upgrade -n gloo-mesh enterprise-agent --kube-context="${cluster}" -f $cluster-values.yaml
  rm $cluster-values.yaml
done
```


## Modifying Istiod

Now that we have created our VirtualMesh to use Vault for the intermediate CA, we need to go ahead and modify our Istio installation to support fetching and dynamically reloading the intermediate CA from Vault.

First things first, we need to get the verison of our components running in cluster:

```shell
export MGMT_PLANE_VERSION=$(meshctl version | jq '.server[].components[] | select(.componentName == "enterprise-networking") | .images[] | select(.name == "enterprise-networking") | .version')
```

Then we need to update our `istiod` deployment with the sidecar to load and store the certificates. Most installations can use the IstioOperator API and `istioctl` or operator for this. However, when running in Kind, a manual json patch is necessary. This operation should be performed on both clusters.

{{< tabs >}}
{{< tab name="Standard" codelang="shell" >}}
cat << EOF | istioctl manifest install -y -f -
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  name: example-istiooperator
  namespace: istio-system
spec:
  profile: minimal
  meshConfig:
    enableAutoMtls: true
    defaultConfig:
      proxyMetadata:
        # Enable Istio agent to handle DNS requests for known hosts
        # Unknown hosts will automatically be resolved using upstream dns servers in resolv.conf
        ISTIO_META_DNS_CAPTURE: "true"
  components:
    pilot:
      k8s:
        overlays:
        - apiVersion: apps/v1
          kind: Deployment
          name: istiod
          patches:
          - path: spec.template.spec.volumes[name:cacerts]
            value: 
              name: cacerts
              secret: null
              emptyDir:
                medium: Memory
          - path: spec.template.spec.containers[1]
            value: 
              name: istiod-agent
              image: gcr.io/gloo-mesh/istiod-agent:$MGMT_PLANE_VERSION
              imagePullPolicy: IfNotPresent
              volumeMounts:
              - mountPath: /etc/cacerts
                name: cacerts
              args: 
              - sidecar
              env:
              - name: PILOT_CERT_PROVIDER
                value: istiod
              - name: POD_NAME
                valueFrom:
                  fieldRef:
                    apiVersion: v1
                    fieldPath: metadata.name
              - name: POD_NAMESPACE
                valueFrom:
                  fieldRef:
                    apiVersion: v1
                    fieldPath: metadata.namespace
              - name: SERVICE_ACCOUNT
                valueFrom:
                  fieldRef:
                    apiVersion: v1
                    fieldPath: spec.serviceAccountName
          - path: spec.template.spec.initContainers
            value: 
            - name: istiod-agent-init
              image: gcr.io/gloo-mesh/istiod-agent:$MGMT_PLANE_VERSION
              imagePullPolicy: IfNotPresent
              volumeMounts:
              - mountPath: /etc/cacerts
                name: cacerts
              args: 
              - init-container
              env:
              - name: PILOT_CERT_PROVIDER
                value: istiod
              - name: POD_NAME
                valueFrom:
                  fieldRef:
                    apiVersion: v1
                    fieldPath: metadata.name
              - name: POD_NAMESPACE
                valueFrom:
                  fieldRef:
                    apiVersion: v1
                    fieldPath: metadata.namespace
              - name: SERVICE_ACCOUNT
                valueFrom:
                  fieldRef:
                    apiVersion: v1
                    fieldPath: spec.serviceAccountName
    # Istio Gateway feature
    ingressGateways:
    - name: istio-ingressgateway
      enabled: true
      k8s:
        env:
          - name: ISTIO_META_ROUTER_MODE
            value: "sni-dnat"
        service:
          type: NodePort
          ports:
            - port: 80
              targetPort: 8080
              name: http2
            - port: 443
              targetPort: 8443
              name: https
            - port: 15443
              targetPort: 15443
              name: tls
              nodePort: 32000
  values:
    global:
      pilotCertProvider: istiod
EOF
{{< /tab >}}
{{< tab name="Kind" codelang="shell" >}}
kubectl patch -n istio-system istiod --patch '{
	"spec": {
			"template": {
				"spec": {
						"initContainers": [
							{
									"args": [
										"init-container"
									],
									"env": [
										{
												"name": "PILOT_CERT_PROVIDER",
												"value": "istiod"
										},
										{
												"name": "POD_NAME",
												"valueFrom": {
													"fieldRef": {
															"apiVersion": "v1",
															"fieldPath": "metadata.name"
													}
												}
										},
										{
												"name": "POD_NAMESPACE",
												"valueFrom": {
													"fieldRef": {
															"apiVersion": "v1",
															"fieldPath": "metadata.namespace"
													}
												}
										},
										{
												"name": "SERVICE_ACCOUNT",
												"valueFrom": {
													"fieldRef": {
															"apiVersion": "v1",
															"fieldPath": "spec.serviceAccountName"
													}
												}
										}
									],
									"volumeMounts": [
										{
												"mountPath": "/etc/cacerts",
												"name": "cacerts"
										}
									],
									"imagePullPolicy": "IfNotPresent",
									"image": "gcr.io/gloo-mesh/istiod-agent:$MGMT_PLANE_VERSION",
									"name": "istiod-agent-init"
							}
						],
						"containers": [
							{
									"args": [
										"sidecar"
									],
									"env": [
										{
												"name": "PILOT_CERT_PROVIDER",
												"value": "istiod"
										},
										{
												"name": "POD_NAME",
												"valueFrom": {
													"fieldRef": {
															"apiVersion": "v1",
															"fieldPath": "metadata.name"
													}
												}
										},
										{
												"name": "POD_NAMESPACE",
												"valueFrom": {
													"fieldRef": {
															"apiVersion": "v1",
															"fieldPath": "metadata.namespace"
													}
												}
										},
										{
												"name": "SERVICE_ACCOUNT",
												"valueFrom": {
													"fieldRef": {
															"apiVersion": "v1",
															"fieldPath": "spec.serviceAccountName"
													}
												}
										}
									],
									"volumeMounts": [
										{
												"mountPath": "/etc/cacerts",
												"name": "cacerts"
										}
									],
									"imagePullPolicy": "IfNotPresent",
									"image": "gcr.io/gloo-mesh/istiod-agent:$MGMT_PLANE_VERSION",
									"name": "istiod-agent"
							}
						],
						"volumes": [
							{
									"name": "cacerts",
									"secret": null,
									"emptyDir": {
										"medium": "Memory"
									}
							}
						]
				}
			}
	}
}'
{{< /tab >}}
{{< /tabs >}}

## Final Steps

Now that Istio has been patched with the Gloo Mesh `istiod-agent` sidecar, we can verify that all of our traffic is being secured using the root CA we generated for Vault in the previous steps.

The easiest way to do this will be to check `root-cert.pem` in the`istio-ca-root-cert` config map which Istio propagates for initial TLS connection. This command will check the propagated root-cert against the local cert which we supplied to vault in an earlier step. If Vault was not setup using the earlier part of the tutorial, the Vault root-cert should instead be fetched and saved to the file `root-cert.pem` in the current directory.

If installed correctly, the output from the following command should be empty.

```shell
for cluster in ${CONTEXT_1} ${CONTEXT_2}; do
  kubectl --context="${cluster}" get cm -n bookinfo istio-ca-root-cert -ojson | jq -r  '.data["root-cert.pem"]' | diff root-cert.pem -
done
```
