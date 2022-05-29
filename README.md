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
sudo kubectl get IngressRoute -A
```

```bash
sudo kubectl logs -f traefik-75b67cbc98-bvzzx -n kube-system
```


Todos
-----

- [Auto update (OS)](https://linoxide.com/enable-automatic-updates-on-ubuntu-20-04/)
- Remove ssh password login
- hardening: 
  - [hardening_guide](https://rancher.com/docs/k3s/latest/en/security/hardening_guide/)
  - [kubescape](https://github.com/armosec/kubescape)
  - [kubescape docu](https://hub.armo.cloud/docs)
  - [YouTube](https://www.youtube.com/watch?v=ZATGiDIDBQk)
- Extended monitoring:
  - [netways prometheus-operator](https://nws.netways.de/de/tutorials/monitoring-kubernetes-mit-prometheus/)
  - [prometheus-operator](https://sysdig.com/blog/kubernetes-monitoring-prometheus-operator-part3/)
- ArgoCD
  - Auto build and push of images
  