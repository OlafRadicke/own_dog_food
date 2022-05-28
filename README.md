Playbook for hosteurope server setup
====================================

Mission
-------

*"Eating your own dog food"* :sweat_smile:

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
ansible-playbook -i ./ansible/hosts.yaml ./ansible/install_and_update.yaml
```

K3s / podman topics
-------------------


Generate config file (for generating image pull secret):

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

- Auto build and push of images
- [Auto update (OS)](https://linoxide.com/enable-automatic-updates-on-ubuntu-20-04/)
- Remove ssh password login
- hardening [README.md](ansible/roles/k3s_install/README.md)
- mTLS https://doc.traefik.io/traefik/https/tls/#client-authentication-mtls and https://smallstep.com/hello-mtls/doc/server/traefik
- Extended monitoring
- ArgoCD
  