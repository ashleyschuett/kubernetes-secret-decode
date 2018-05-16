# Kubernetes Secret Decode

### Description
Be able to easily see the values of a secret.

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
data:
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
mkdir -p $GOPATH/github.com/ashleyschuett && \
cd $GOPATH/github.com/ashleyschuett && \
git clone https://github.com/ashleyschuett/kubernetes-secret-decode.git && \
cd kubernetes-secret-decode && \
go build -o $GOPATH/bin/ksd
```

```
go get github.com/ashleyschuett/kubernetes-secret-decode && \
cd $GOPATH/github.com/ashleyschuett && \
go build -o $GOPATH/bin/ksd
```

Another option is to download the binary and add it to your path
```
curl -LO https://github.com/ashleyschuett/kubernetes-secret-decode/releases/download/v1.0/ksd && chmod +x ksd && sudo mv ksd /usr/local/bin 
```

### Usage
`kubectl get secret my-secret -o yaml | ksd`
