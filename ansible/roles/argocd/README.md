argocd
======

The role installed an ArgoCD.

Login
-----

User: admin
Get password enter:

```bash
sudo kubectl get secret \
  argocd-initial-admin-secret \
  -n argocd \
  -o jsonpath="{.data.password}" \
  | base64 -d; echo
```

