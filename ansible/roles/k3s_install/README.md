k3s_install
===========

This role installed a k3s.

Requirements
------------


Role Variables
--------------



Dependencies
------------

```bash
oradicke@ip-92-205-105-117:~$ sudo cat /var/lib/rancher/k3s/server/manifests/traefik.yaml
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: traefik-crd
  namespace: kube-system
spec:
  chart: https://%{KUBERNETES_API}%/static/charts/traefik-crd-10.14.100.tgz
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: traefik
  namespace: kube-system
spec:
  chart: https://%{KUBERNETES_API}%/static/charts/traefik-10.14.100.tgz
  set:
    global.systemDefaultRegistry: ""
  valuesContent: |-
    rbac:
      enabled: true
    ports:
      websecure:
        tls:
          enabled: true
    podAnnotations:
      prometheus.io/port: "8082"
      prometheus.io/scrape: "true"
    providers:
      kubernetesIngress:
        publishedService:
          enabled: true
    priorityClassName: "system-cluster-critical"
    image:
      name: "rancher/mirrored-library-traefik"
      tag: "2.6.1"
    tolerations:
    - key: "CriticalAddonsOnly"
      operator: "Exists"
    - key: "node-role.kubernetes.io/control-plane"
      operator: "Exists"
      effect: "NoSchedule"
    - key: "node-role.kubernetes.io/master"
      operator: "Exists"
      effect: "NoSchedule"
```


Example Playbook
----------------


Background information
----------------------

See:

- [install-options](https://rancher.com/docs/k3s/latest/en/installation/install-options/)
- [hardening_guide](https://rancher.com/docs/k3s/latest/en/security/hardening_guide/)
- [k3s/networking](https://rancher.com/docs/k3s/latest/en/networking/)
- [traefik/providers](https://doc.traefik.io/traefik/providers/kubernetes-ingress/)
- [Traefik & CRD & Let's Encrypt](https://doc.traefik.io/traefik/v2.0/user-guides/crd-acme/)
- [traefik/lets-encrypt](https://doc.traefik.io/traefik/https/acme/#lets-encrypt)
- [-> https-on-kubernetes-using-traefik-proxy](https://traefik.io/blog/https-on-kubernetes-using-traefik-proxy/)

License
-------

BSD

Author Information
------------------

- Olaf Radicke <briefkasten@olaf-radicke.de>
