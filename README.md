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

### Routing

```bash
sudo kubectl get IngressRoute -A
```

```bash
sudo kubectl logs -f traefik-75b67cbc98-bvzzx -n kube-system
```

### ArgoCD

To find out password enter:

```
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

Connect without ingress enter:

```
sudo kubectl port-forward svc/argocd-server -n argocd 8080:443
```

```
 ssh -L localhost:8080:localhost:8080 oradicke@92.205.105.117
```

Known issue
-----------

### ArgoCD => ERR_TOO_MANY_REDIRECTS

The install yaml need a change:

```
*** 10235,10244 ****
--- 10235,10245 ----
                topologyKey: kubernetes.io/hostname
              weight: 5
        containers:
        - command:
          - argocd-server
+         - --insecure
          env:
          - name: ARGOCD_SERVER_INSECURE
            valueFrom:
              configMapKeyRef:
                key: server.insecure
```