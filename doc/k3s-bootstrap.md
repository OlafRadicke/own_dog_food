K3s bootstraping
================


- [K3s bootstraping](#k3s-bootstraping)
	- [ANSIBLE ROLES](#ansible-roles)
	- [STORAGE](#storage)
		- [Configuration of driver](#configuration-of-driver)
		- [Default storage class](#default-storage-class)
		- [Extra storage class](#extra-storage-class)



ANSIBLE ROLES
-------------

The kubernetes ist installed by an Ansible playbook. The
relevant roles and ths READMEs are to find in...

- [The K3s Installation](../ansible/roles/k3s_install/README.md)
- [The K3s namespaces](../ansible/roles/k3s_namespaces)
- [The let's encrypt support](../ansible/roles/k3s_cert_manager/README.md)
- [The mutual TLS support (PKI)](../ansible/roles/k3s_mtls/README.md)
- [The mutual TLS routing](../ansible/roles/k3s_mtls_routes)
- [The K3s unsinstall](../ansible/roles/k8s_ansible_requirements/README.md)


STORAGE
-------

### Configuration of driver

See `/var/lib/rancher/k3s/server/manifests/local-storage.yaml`

### Default storage class

local path: `/var/lib/rancher/k3s/storage`

### Extra storage class

Name: `manual-pv`

Local path: `/srv/k3s-pv`
