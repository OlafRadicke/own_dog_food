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




INSTALL OPERATOR
----------------

- [Installation docu](https://smallstep.com/docs/certificate-manager/kubernetes-autocert/#deploy-autocert)
- [Git-Repo](https://github.com/smallstep/autocert/tree/master)
