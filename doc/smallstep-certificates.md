smallstep certificates
======================


LINKS
-----

- [Project site](https://github.com/smallstep/certificates)
- [Limitations](https://smallstep.com/docs/step-ca/index.html#limitations)
- [Install step-ca](https://smallstep.com/docs/step-ca/installation/index.html)



INSTALL
-------

```bash
helm repo add smallstep https://smallstep.github.io/helm-charts/
helm repo update
helm install step-certificates smallstep/step-certificates
```

parameters of the Step certificates chart:
https://artifacthub.io/packages/helm/smallstep/step-certificates