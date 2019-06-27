# Kubernetes Secret Decode

### Description
Be able to easily see the values of a secret.

YAML and JSON are both supported and detection of the input type is performed automatically.

**Before:**
```yaml
$ kubectl get secret my-secret -o yaml
apiVersion: v1
data:
  password: cGFzc3dvcmQ=
  username: dXNlcm5hbWU=
kind: Secret
metadata:
  creationTimestamp: 2018-05-09T21:01:37Z
  name: my-secret
  namespace: default
  resourceVersion: "20229"
  selfLink: /api/v1/namespaces/default/secrets/my-secret
  uid: 29ef8024-53cc-11e8-967d-080027cd91ae
type: Opaque
```

**After:**
```yaml
$ kubectl ksd get secret my-secret -o yaml
apiVersion: v1
stringData:
  password: password
  username: username
kind: Secret
metadata:
  creationTimestamp: "2018-05-09T21:01:37Z"
  name: my-secret
  namespace: default
  resourceVersion: "20229"
  selfLink: /api/v1/namespaces/default/secrets/my-secret
  uid: 29ef8024-53cc-11e8-967d-080027cd91ae
type: Opaque
```

### Installation

These instructions assume you have go installed and a `$GOPATH` set.
The binary needs to be installed somewhere in your `$PATH`.
The binary also needs to be named either `kubectl-ksd` or `kubectl-kubernetes-secret-decode`

For easy install running the following:
```
make install
```

Another option is to download the binary and add it to your path
```
curl -LO https://github.com/ashleyschuett/kubernetes-secret-decode/releases/download/v3.0.0/kubectl-ksd && chmod +x kubectl-ksd && sudo mv kubectl-ksd /usr/local/bin
```

### Usage
`kubectl ksd get secret my-secret -o yaml`

`kubectl ksd get secret my-secret -o json`
