Playbook for hosteurope server setup
====================================

Pre install
-----------

Install needed collections and python libs, see:

- [k3s_deployment](ansible/roles/k3s_deployment/README.md)


Run setup
---------


Enter:

```
ansible-playbook -i ./hosts.yaml ./run_all.yaml
```

Todos
-----

- Remove password login
- hardening [README.md](ansible/roles/k3s_install/README.md)