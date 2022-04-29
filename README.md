Playbook for hosteurope server setup
====================================

Pre install requirements
------------------------

Enter to install needed collections:


```bash
ansible-galaxy collection install kubernetes.core
```


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