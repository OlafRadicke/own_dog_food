OpenBao
=======

- [OpenBao](#openbao)
	- [K8s-secret with OpenBao-Token](#k8s-secret-with-openbao-token)
	- [Port forwart](#port-forwart)
	- [Pre step OpenBao / vault cli](#pre-step-openbao--vault-cli)
		- [/etc/hosts](#etchosts)
		- [Add openbao env variables](#add-openbao-env-variables)
		- [Use vault cli](#use-vault-cli)
	- [Use OpenBao / Vault cli](#use-openbao--vault-cli)
		- [Status](#status)


K8s-secret with OpenBao-Token
-----------------------------

Create k8s-secret with OpenBao-Token

```yaml
apiVersion:   v1
kind:         Secret
metadata:
  name:       vault-token
type:         Opaque
data:
  token:      Q2hhTkdfbUUscGxFQXNlISE=
```

And enter

```bash
$ kubectl -n opentofu apply -f ./workspace/vault-token.yaml
```

Port forwart
------------

Use script:

```bash
scrips/port-forward/openbao.sh
```


Pre step OpenBao / vault cli
----------------------------


### /etc/hosts

Add OpenBao-Server in line...

```bash
127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4 openbao.openbao
```

### Add openbao env variables

Add file `${HOME}/.ssh/openbao.env` with content:

```bash
export VAULT_TOKEN=XXXXXXX
export VAULT_ADDR='http://openbao.openbao:8083'
```

Add this line in file `~/.bashrc`

```bash
source ${HOME}/.ssh/openbao.env
```

### Use vault cli


Use OpenBao / Vault cli
-----------------------

### Status

```bash
$ vault status

Key             Value
---             -----
Seal Type       shamir
Initialized     true
Sealed          false
Total Shares    1
Threshold       1
Version         2.2.0
Build Date      2025-03-05T13:07:08Z
Storage Type    file
Cluster Name    vault-cluster-248e257a
Cluster ID      681288cb-727b-b80d-4b64-f131b46c71c1
HA Enabled      false
```