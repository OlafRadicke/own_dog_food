Troubleshooting
===============


Routing
-------

```bash
sudo kubectl get IngressRoute -A
```

```bash
sudo kubectl logs -f traefik-75b67cbc98-bvzzx -n kube-system
```

ArgoCD
------

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