UPTIME KUMA
===========


SETUP
-----

Set password without ingress route:

```bash
kubectl port-forward service/uptime-kuma -n pulumi-apps 8080:80
```