Smallstep Autocert
==================

INSTALL CLIENT
--------------

```bash
$ sudo su
$ cat <<EOT > /etc/yum.repos.d/smallstep.repo
[smallstep]
name=Smallstep
baseurl=https://packages.smallstep.com/stable/fedora/
enabled=1
repo_gpgcheck=0
gpgcheck=1
gpgkey=https://packages.smallstep.com/keys/smallstep-0x889B19391F774443.gpg
EOT
dnf makecache && dnf install -y step-cli step-ca
```

INIT ROOT CA
------------

```
step ca bootstrap --ca-url small.irish-sea \
	--install
```


INIT AUTOCERT
-------------

Install the root CA certificat of the root CA:



```bash
$ kubectl -n smallstep-root-ca \
    get configmap \
    smallstep-root-ca-step-certificates-certs \
    -o jsonpath="{.data.root_ca\.crt}" > /tmp/step-root_ca.crt
```

```bash
$ AUTOCERT_NAMESPACE=smallstep-autocert
$ kubectl -n $AUTOCERT_NAMESPACE \
    create configmap certs \
    --from-file=/tmp/step-root_ca.crt \
    --dry-run=true \
    -o yaml
```

Create secret with password

```bash
$ kubectl -n $AUTOCERT_NAMESPACE \
    create secret generic \
    autocert-password \
    --from-file=password=autocert-password.txt \
    --dry-run=true \
    -o yaml
```

Create configmap config

```bash
$ kubectl -n $AUTOCERT_NAMESPACE \
    create configmap config \
    --from-file $(step path)/config \
    --dry-run=true \
    -o yaml

```

INSTALL OPERATOR
----------------

- [Installation docu](https://smallstep.com/docs/certificate-manager/kubernetes-autocert/#deploy-autocert)
- [Git-Repo](https://github.com/smallstep/autocert/tree/master)
