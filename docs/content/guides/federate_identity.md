---
title: Federated Trust and Identity
menuTitle: Federated Trust and Identity
weight: 25
---

Gloo Mesh can help unify the root identity between multiple service mesh installations so any intermediates are signed by the same Root CA and end-to-end mTLS between clusters and destinations can be established correctly.

Gloo Mesh will establish trust based on the [trust model](https://spiffe.io/spiffe/concepts/#trust-domain) defined by the user -- is there complete *shared trust* and a common root and identity? Or is there *limited trust* between clusters and traffic is gated by egress and ingress gateways? 

In this guide, we'll explore the *shared trust* model between two Istio clusters and how Gloo Mesh simplifies and orchestrates the processes needed for this to happen.

## Before you begin
To illustrate these concepts, we will assume that:

* There are two clusters managed by Gloo Mesh named `cluster-1` and `cluster-2`. 
* Istio is [installed on both client clusters]({{% versioned_link_path fromRoot="/guides/installing_istio" %}})
* The `bookinfo` app is [installed across the two clusters]({{% versioned_link_path fromRoot="/guides/#bookinfo-deployed-on-two-clusters" %}})

{{% notice note %}}
Be sure to review the assumptions and satisfy the pre-requisites from the [Guides]({{% versioned_link_path fromRoot="/guides" %}}) top-level document.
{{% /notice %}}

Ensure you have the correct context names set in your environment:

```shell
CONTEXT_1=your_first_context
CONTEXT_2=your_second_context
```

## Enforce mTLS

Apply the following yaml to both clusters,
assuming that istio-system is the root namespace for the istio deployment:

{{< tabs >}}
{{< tab name="YAML file" codelang="yaml">}}
apiVersion: "security.istio.io/v1beta1"
kind: "PeerAuthentication"
metadata:
  name: "default"
  namespace: "istio-system"
spec:
  mtls:
    mode: STRICT
{{< /tab >}}
{{< tab name="CLI inline" codelang="shell" >}}
kubectl apply --context $CONTEXT_1 -f - << EOF
apiVersion: "security.istio.io/v1beta1"
kind: "PeerAuthentication"
metadata:
  name: "default"
  namespace: "istio-system"
spec:
  mtls:
    mode: STRICT
EOF
kubectl apply --context $CONTEXT_2 -f - << EOF
apiVersion: "security.istio.io/v1beta1"
kind: "PeerAuthentication"
metadata:
  name: "default"
  namespace: "istio-system"
spec:
  mtls:
    mode: STRICT
EOF
{{< /tab >}}
{{< /tabs >}}

This is an Istio setting. For more, see: https://istio.io/latest/docs/concepts/security/

## Verify identity in two clusters is different

We can see the certificate chain used to establish mTLS between Istio destinations in `cluster-1` and `cluster-2` and can compare them to be different. One way to see the certificates, is to use the `openssl s_client` tool with the `-showcerts` param when calling between two destinations. Let's try it on `cluster-1`:

```shell
kubectl --context $CONTEXT_1 -n bookinfo exec -it deploy/reviews-v1 -c istio-proxy \
-- openssl s_client -showcerts -connect ratings.bookinfo:9080
```
You should see an output of the certificate chain among other handshake-related information. You can review the last certificate in the chain and that's the root cert:

{{< highlight shell "hl_lines=24-43" >}}
---
Certificate chain
 0 s:
   i:O = cluster.local
-----BEGIN CERTIFICATE-----
MIIDKTCCAhGgAwIBAgIRAKZzWK3r7aZqVd0pXUalzKIwDQYJKoZIhvcNAQELBQAw
GDEWMBQGA1UEChMNY2x1c3Rlci5sb2NhbDAeFw0yMDA0MjMxMjM2MDFaFw0yMDA0
MjQxMjM2MDFaMAAwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCwCZft
3uavGRCv+ooKVWUod7Z3PWukPGIR0icI12+ghFygT9ZKlyu+LQ8iN93A6/soIo8r
ccp5YCyW9O4JCJSPg+iFqGeg9yNDLCATb+6vwTsHx0rvdLdX4803bjF9evWkr5yZ
AlPath6S/Wxihue2xrnw9mSF1nKRQxxw8ypysKiqLfVNBhCnBsN28gppYnl1pIiv
YamBeSiNA887BDnuXIc6t6yNJudlvefuixUhzBeR9zYNlstWBLsdqSubbPdPVxfv
7H6hRjeAmu0VB2oDpsWJ0OYGu42ZavCSHRIL2nD3fqk9DyWtXKIklU4B7rE4mySe
InDLbkj1lLv1FB1VAgMBAAGjgYUwgYIwDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQW
MBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMEMGA1UdEQEB/wQ5
MDeGNXNwaWZmZTovL2NsdXN0ZXIubG9jYWwvbnMvZGVmYXVsdC9zYS9ib29raW5m
by1yYXRpbmdzMA0GCSqGSIb3DQEBCwUAA4IBAQBEUi1ge/M3NlQ6xuizY7o/mkFe
+PXKjKT/vf21d6N5FTAJT4fjL6nEsa4NqJC7Txiz9kEjlqLy/SywtB3qYGuG0/+d
QGgWmN1NVOMtl2Kq++LOQOIaEV24mjHb5r38DVk4YyVH2E/1QAWByONDB54Ovlyf
l3qiE3gEeegKsgtsLuhzQCReU5evdmPhnCAMiZvUhQKxHIoCJEx5A+eB4q2zBDN2
H2CNJyWLPulBNCsZvCYXGLDRIy+Sp9AsXhqMTAxvqNS2NaNQ9fh7gSqOORO3kIKz
axoFg6neo+LAaYwoyBtO7/V9OvShd9TbkyPP4amR7k3zkdulFo2o+jKAqzCq
-----END CERTIFICATE-----
 1 s:O = cluster.local
   i:O = cluster.local
-----BEGIN CERTIFICATE-----
MIIC3TCCAcWgAwIBAgIQDZ3lILg70fkKSuuBpj3O7TANBgkqhkiG9w0BAQsFADAY
MRYwFAYDVQQKEw1jbHVzdGVyLmxvY2FsMB4XDTIwMDQyMjIzMjM0NVoXDTMwMDQy
MDIzMjM0NVowGDEWMBQGA1UEChMNY2x1c3Rlci5sb2NhbDCCASIwDQYJKoZIhvcN
AQEBBQADggEPADCCAQoCggEBALX8anGTKtlpdbIlwGsxTW/ZJeqSM29eei5Lmsee
wll7xaNE4sNaj6HFyqAZomPDJm/4PYZ0fWmJ1FIXFqCXQ6PNf/J592D1x8oIHh50
88BbOkH7wYzEMymoP+2BqXQsY5kxjCg9N6xj4XygSunjXo3ctyVP11GhUew0j+Aw
U7dtZqWlpgMsZsPEn2V4JFid20q+0qz6iCzRh/a3iO98QSfvlpeQuVQkhLiPZOzA
q796C1HLWU7sefkXzVAsQGHA5FqSQLQbOqXBPWaf82Fw9cO4/skBH/qOIVtIh8Ks
rHMgrYkSXprev4bMafUAdfJ9GLity/4D2Mn0rK3k4GiLoL8CAwEAAaMjMCEwDgYD
VR0PAQH/BAQDAgIEMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEB
AGZVlJzyM9E14eucxf1UmJvt5NeGGmgf8YdUd9R693iffG8eAg44VJP6xZ8LTrj7
WoGoWC9SmclvAm2Lo1zh7JQWL2jca5X/aJSW4CZROGDMwWm9e+SaujsOKG3hhis6
iwTl1VsjV4o0UBRm5Z26T/gn1CoIPjQDJRb86RPr/6DHY8jFGvGjceEl+o6mf+gk
Q0xfk7VNxpxScJ/+lU5+IJrqQTBmrhk40eDe24D4zbtnk4YVRRbiMh4p9PIBySyp
gyMylEJ3SgwpVoWwV0e2UvNCG1AlZADiYPpgy2qANzJqtF/GYjfgcpR01r8LceIj
s2rL2u8nTerM5bjlurn1Z58=
-----END CERTIFICATE-----
{{< /highlight >}}

Run the same thing in `cluster-2` and explore the output and compare. For the `reviews` service running in `cluster-2`, we have to use `deploy/reviews-v3` as `reviews-v1` which we used in the previous command doesn't exist on that cluster:


```shell
kubectl --context $CONTEXT_2 -n bookinfo exec -it deploy/reviews-v3 -c istio-proxy \
-- openssl s_client -showcerts -connect ratings.bookinfo:9080
```

You should notice that the root certificates that signed the workload certificates are **different**. Let's unify those into a *shared trust* model of identity. 

## Creating a Virtual Mesh

Gloo Mesh uses the [Virtual Mesh]({{% versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh/" %}}) Custom Resource to configure a Virtual Mesh, which is a logical grouping of one or multiple service meshes for the purposes of federation according to some parameters. Let's take a look at a VirtualMesh configuration that can help unify our two service meshes and establish a *shared trust* model for identity:

{{< highlight yaml "hl_lines=8-15 17-21" >}}
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
      rootCertificateAuthority:
        generated: {}
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

We are creating the VirtualMesh with two different service meshes: `istiod-istio-system-cluster-1` and `istiod-istio-system-cluster-2`. We can have any meshes defined here that should be part of this virtual grouping and federation.

##### Establishing Trust

In the first highlighted section, we can see the parameters establishing shared identity and federation. In this case, we tell Gloo Mesh to create a Root CA using the parameters specified above (ttl, key size, org name, etc).

We could have also configured an existing Root CA by providing an existing secret:

```yaml
  mtlsConfig:
    # Note: Do NOT use this autoRestartPods setting in production!!
    autoRestartPods: true
    shared:
      rootCertificateAuthority:
        secret:
          name: root-ca-name
          namespace: root-ca-namespace
```

See the section on [User Provided Certificates]({{% versioned_link_path fromRoot="/guides/federate_identity/#user-provided-certificates" %}}) below for details on how to format the certificate as a Kubernetes Secret.

##### User Provided Certificates

A root certificate for a VirtualMesh must be supplied to Gloo Mesh 
as a Secret formatted as follows:

```yaml
kind: Secret
metadata:
  name: providedrootcert
  namespace: default
type: Opaque
data:
  key.pem: {private key file}
  root-cert.pem: {root CA certificate file}
```

Given a root certificate file `root-cert.pem` and its associated private key file `key.pem`,
this secret can be created by running:

`kubectl -n default create secret generic providedrootcert --from-file=root-cert.pem --from-file=key.pem`.

An example root certificate and private key file can be generated by following 
[this guide](https://github.com/istio/istio/tree/1.8.0/samples/certs) and running `make root-ca`.

Note that the name/namespace of the provided root cert cannot be `cacerts/istio-system` as that is used by Gloo Mesh for carrying out the CSR ([certificate signing request](https://en.wikipedia.org/wiki/Certificate_signing_request)) procedure
that unifies the trust root between Meshes in the VirtualMesh.

##### Destination Federation

We also specify which destinations to federate to remote service meshes through the `federation` field. In the example we provide a single empty selector,
which expresses permissive federation, i.e. exposes all destinations from each mesh to the other mesh.

Alternatively, we can fine tune which destinations are exposed to which service meshes by specifying selection criteria.
For instance, the following `federation` stanza selectively federates the `reviews` destination on `cluster-1` to `cluster-2`, and vice versa:

```yaml
federation:
  hostnameSuffix: "soloio"
  selectors:
  - destinationSelectors:
    - kubeServiceRefs:
        services:
        - name: reviews
          namespace: bookinfo
          clusterName: cluster-1
    meshes:
    - name: istiod-istio-system-cluster-2
      namespace: gloo-mesh
  - destinationSelectors:
    - kubeServiceRefs:
        services:
        - name: reviews
          namespace: bookinfo
          clusterName: cluster-2
    meshes:
    - name: istiod-istio-system-cluster-1
      namespace: gloo-mesh
```

##### Applying VirtualMesh

If you saved this VirtualMesh CR to a file named `demo-virtual-mesh.yaml`, you can apply it like this:

```shell
kubectl --context $CONTEXT_1 apply -f demo-virtual-mesh.yaml
```

Notice the `autoRestartPods: true` in the mtlsConfig stanza. This instructs Gloo Mesh to restart ALL of the Istio pods in the relevant clusters. DO NOT SET THIS `true` IN PRODUCTION. 

This is an optional convenience flag for testing in development to speed up workload certificate rotation. This is due to a limitation of Istio. The Istio control plane picks up the CA for Citadel and does not rotate it often enough. This is being [improved in future versions of Istio](https://github.com/istio/istio/issues/22993). 



{{% notice note %}}
Note, after you bounce the control plane, it may still take time for the workload certs to get re-issued with the new CA. You can force the workloads to re-load by bouncing them. For example, for the bookinfo sample running in the `bookinfo` namespace:

```shell
kubectl --context $CONTEXT_1 -n bookinfo delete po --all
kubectl --context $CONTEXT_2 -n bookinfo delete po --all
```
{{% /notice %}}

Creating this resource will instruct Service Mesh to establish a shared root identity across the clusters in the Virtual Mesh as well as federate the destinations. The next sections of this document help you understand some of the pieces of how this works.

## Understanding the Shared Root Process

When we create the VirtualMesh CR, set the trust model to `shared`, and configure the Root CA parameters, Gloo Mesh will kick off the process to unify the identity to a shared root. First, Gloo Mesh will either create the Root CA specified (if `generated` is used) or use the supplied CA information. 

Then Gloo Mesh will use a Certificate Request (CR) agent on each of the affected clusters to create a new key/cert pair that will form an intermediate CA used by the mesh on that cluster. It will then create a Certificate Request, represented by the [CertificateRequest]({{% versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.certificates.certificate_request/" %}}) CR.

 Gloo Mesh will sign the certificate with the Root CA specified in the VirtualMesh. At that point, we will want the mesh (Istio in this case) to pick up the new intermediate CA and start using that for its workloads.

![Gloo Mesh Architecture]({{% versioned_link_path fromRoot="/img/gloomesh-csr.png" %}})

To verify, let's check the `IssuedCertificates` CR in `cluster-2-context`:

```shell
kubectl --context $CONTEXT_2 \
get issuedcertificates -n gloo-mesh
```

We should see this on `cluster-2`:

```shell
NAME                                 AGE
istiod-istio-system-cluster-2   3m15s
```

If we do the same on `cluster-1`, we should also see an `IssuedCertificates` entry there as well.

Lastly, let's verify the correct `cacerts` was created in the `istio-system` namespace that can be used for Istio's Citadel:

```shell
kubectl --context $CONTEXT_1 get secret -n istio-system cacerts 

NAME      TYPE                                          DATA   AGE
cacerts   certificates.mesh.gloo.solo.io/issued_certificate   5      20s
```

```shell
kubectl --context $CONTEXT_2 get secret -n istio-system cacerts 

NAME      TYPE                                          DATA   AGE
cacerts   certificates.mesh.gloo.solo.io/issued_certificate   5      5m3s
```

In the previous section, we bounced the Istio control plane to pick up these intermediate certs. Again, this is being [improved in future versions of Istio](https://github.com/istio/istio/issues/22993).

##### Disabling Shared Trust Process

The automated trust process described above can be disabled by simply omitting the `mtlsConfig` field from the VirtualMesh. Doing so
will prevent Gloo Mesh from managing Istio's certificates. This is useful in scenarios where establishing shared trust between 
service mesh deployments is performed manually or by some external process.

##### Multi-cluster mesh federation

Once trust has been established, Gloo Mesh will start federating destinations so that they are accessible across clusters. Behind the scenes, Gloo Mesh will handle the networking -- possibly through egress and ingress gateways, and possibly affected by user-defined traffic and access policies -- and ensure requests to the service will resolve and be routed to the right destination. Users can fine-tune which destinations are federated where by editing the virtual mesh. 

For example, you can see what Istio `ServiceEntry` objects were created. On `cluster-1` you can see:

```shell
kubectl --context $CONTEXT_1 \
  get serviceentry -n istio-system
```

```shell
NAME                                                          HOSTS                                                           LOCATION        RESOLUTION   AGE
istio-ingressgateway.istio-system.svc.cluster-2.global        [istio-ingressgateway.istio-system.svc.cluster-2.global]        MESH_INTERNAL   DNS          6m2s
ratings.bookinfo.svc.cluster-2.global                         [ratings.bookinfo.svc.cluster-2.global]                         MESH_INTERNAL   DNS          6m2s
reviews.bookinfo.svc.cluster-2.global                         [reviews.bookinfo.svc.cluster-2.global]                         MESH_INTERNAL   DNS          6m2s
```

On the `cluster-2-context` cluster, you can see:

```shell
kubectl --context $CONTEXT_2 \
get serviceentry -n istio-system
```

```shell
NAME                                                            HOSTS                                                             LOCATION        RESOLUTION   AGE
details.bookinfo.svc.cluster-1.global                       [details.bookinfo.svc.cluster-1.global]                       MESH_INTERNAL   DNS          2m5s     
istio-ingressgateway.istio-system.svc.cluster-1.global      [istio-ingressgateway.istio-system.svc.cluster-1.global]      MESH_INTERNAL   DNS          5m18s    
productpage.bookinfo.svc.cluster-1.global                   [productpage.bookinfo.svc.cluster-1.global]                   MESH_INTERNAL   DNS          55s      
ratings.bookinfo.svc.cluster-1.global                       [ratings.bookinfo.svc.cluster-1.global]                       MESH_INTERNAL   DNS          7m2s     
reviews.bookinfo.svc.cluster-1.global                       [reviews.bookinfo.svc.cluster-1.global]                       MESH_INTERNAL   DNS          90s 
```

## See it in action

Check out "Part Two" of the ["Dive into Gloo Mesh" video series](https://www.youtube.com/watch?v=4sWikVELr5M&list=PLBOtlFtGznBjr4E9xYHH9eVyiOwnk1ciK)
(note that the video content reflects Gloo Mesh <b>v0.6.1</b>):

<iframe width="560" height="315" src="https://www.youtube.com/embed/djcDaIsqIl8" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>


## Next steps

At this point, you should be able to route traffic across your clusters with end-to-end mTLS. You can verify the certs following the same [approach we did earlier in this section]({{% versioned_link_path fromRoot="/guides/federate_identity/#verify-identity-in-two-clusters-is-different" %}}).

Now that you have a single logical "virtual mesh" you can begin configuring it with an API that is aware of this VirtualMesh concept. In the next sections, you can apply [access control]({{% versioned_link_path fromRoot="/guides/access_control_intro/" %}}) and [traffic policies]({{% versioned_link_path fromRoot="/guides/multicluster_communication/" %}}). 

