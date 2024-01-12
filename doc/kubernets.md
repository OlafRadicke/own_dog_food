KUBERNETES
==========


LOCAL STORAGE
-------------

The path /srv/k8s-pvc is reserve for local storage class.


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