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
ansible-playbook -i ./ansible/hosts.yaml ./ansible/run_all.yaml
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

- Auto build and push of images
- Auto update (OS)
- Remove ssh password login
- tlsoptions [letsencrypt](ansible/roles/k3s_deployment/tasks/letsencrypt.yaml) https://www.ssllabs.com/ssltest/analyze.html?d=the-independent-friend.de
- hardening [README.md](ansible/roles/k3s_install/README.md)
- mTLS https://doc.traefik.io/traefik/https/tls/#client-authentication-mtls and https://smallstep.com/hello-mtls/doc/server/traefik
- Monitoring
- Switch olaf-radicke.de
- new with https://enbizcard.vishnuraghav.com/ 
- ArgoCD
  