ANSIBLE
=======

PRE INSTALL REQUIREMENTS
------------------------

Enter to install needed collections:


```bash
ansible-galaxy collection install kubernetes.core
```

RUN PLAYBOOK
------------

Enter:

```
ansible-playbook \
  -i ./ansible/hosts.yaml \
  --vault-password-file ~/.ssh/ansible_vault \
  ./ansible/install_and_update.yaml \
  --check
```

or

```bash
ansible-playbook \
 -i ./ansible/hosts.yaml \
 --vault-password-file ~/.ssh/ansible_vault \
 ./ansible/install_and_update.yaml \
 -l baltic-sea \
  --check
```

and

```bash
ansible-playbook \
 -i ./ansible/hosts.yaml \
 --vault-password-file ~/.ssh/ansible_vault \
 ./ansible/install_and_update.yaml \
 -l atlantic-ocean \
  --check
```
