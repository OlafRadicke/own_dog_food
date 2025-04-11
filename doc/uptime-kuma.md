UPTIME KUMA
===========


SETUP
-----

Set password without ingress route:

```bash
kubectl port-forward service/uptime-kuma -n pulumi-apps 8080:80
```

or with script:

```bash
$ scrips/port-forward/uptime-kuma.sh
```

CONFIGURATION
-------------

`Uptime_Kuma_Backup_*`