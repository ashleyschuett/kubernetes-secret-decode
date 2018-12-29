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
$ kubectl get secret my-secret -o yaml | ksd
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

These instructions assume you have go installed and a $GOPATH set.

```
make install
```

If you do not have go installed locally, but have Docker:

```bash
docker run -u $(id -u):$(id -g) --rm -v "$PWD":/go/src/github.com/ashleyschuett/kubernetes-secret-decode -w /go/src/github.com/ashleyschuett/kubernetes-secret-decode billyteves/alpine-golang-glide:1.2.0 bash -c 'glide update && go build -v -o ksd'
```

Another option is to download the binary and add it to your path
```
curl -LO https://github.com/ashleyschuett/kubernetes-secret-decode/releases/download/v2.0.0/ksd && chmod +x ksd && sudo mv ksd /usr/local/bin
```

### Usage
`kubectl get secret my-secret -o yaml | ksd`

`kubectl get secret my-secret -o json | ksd`
