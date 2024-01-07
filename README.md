Playbook for hosteurope server setup
====================================

Mission
-------

*"Eating your (my) own dog food"* :sweat_smile:


Pre install requirements
------------------------

Enter to install needed collections:


```bash
ansible-galaxy collection install kubernetes.core
```


ANSIBLE
-------

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

PULUMI
------

### RUN AND UPDATE

To run Pulumi enter:

```bash
pulumi/baltic-sea/run.sh
```

or

```bash
pulumi/atlantic-ocean/run.sh
```

### BACKEND

As backend is S3 from Google Cloud in use: https://cloud.google.com/

G.g.f:

```bash
gcloud auth application-default login
```


Change workflow
---------------

- Change helm chart (if necessary )
- Rebuild chart and reposytory (see [helm/README.md](helm/README.md))
- Commit changes (include add new helm tgz file)
- Push changes
- Change playbook
  - Change helm chart version (if necessary )
  - Change helm values (if necessary )
- Run playbook

***Note:*** the helm chart repository is maped onliy on the develop branch!

Other READMEs
-------------

- [K3s / podman topics](doc/K3s-podman-topics.md)
- [Troubleshooting](doc/Troubleshooting.md)
- [Known issue](doc/Known-issue.md)
- ***Ansible roles***
  - [os_config](ansible/roles/os_configs/README.md)
  - [k3s_cert_manager](ansible/roles/k3s_cert_manager/README.md)
  - [k3s_install](ansible/roles/k3s_install/README.md)
  - [k8s_ansible_requirements](ansible/roles/k8s_ansible_requirements/README.md)
  - [argocd](ansible/roles/argocd/README.md)
  - [k3s_mtls](ansible/roles/k3s_mtls/README.md)
  - [grafana](ansible/roles/grafana/README.md)
  - [Prometheus](ansible/roles/prometheus/README.md)
- [pulumi]()

Links
-----

- [bitnami/argo-workflows](https://github.com/bitnami/charts/tree/master/bitnami/argo-workflows)
- [helm repos on github](https://medium.com/@mattiaperi/create-a-public-helm-chart-repository-with-github-pages-49b180dbb417)
- [concourse chart](https://github.com/concourse/concourse-chart)
- [tekton-pipeline](https://github.com/cdfoundation/tekton-helm-chart