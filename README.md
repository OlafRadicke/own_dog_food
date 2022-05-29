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
sudo kubectl get IngressRoute -A
```

```bash
sudo kubectl logs -f traefik-75b67cbc98-bvzzx -n kube-system
```
  