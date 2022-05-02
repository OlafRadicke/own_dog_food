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

K3s topics
----------


Gerate config file (for generating image pull secret):

```bash
podman login --authfile ~/.docker/config.json docker.io
```

Troubleshooting
---------------


```bash
sudo kubectl get ingress -A
```

```bash
sudo kubectl logs -f traefik-75b67cbc98-bvzzx -n kube-system
```

Todos
-----

- Remove ssh password login
- tlsoptions [letsencrypt](ansible/roles/k3s_deployment/tasks/letsencrypt.yaml)
- hardening [README.md](ansible/roles/k3s_install/README.md)
- Monitoring
- Switch olaf-radicke.de
- new with https://enbizcard.vishnuraghav.com/ 