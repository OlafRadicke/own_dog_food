Todos
=====


- Prometeus:
  - Add node-exporter
- Extended monitoring:
  - [netways prometheus-operator](https://nws.netways.de/de/tutorials/monitoring-kubernetes-mit-prometheus/)
  - [prometheus-operator](https://sysdig.com/blog/kubernetes-monitoring-prometheus-operator-part3/)
- [argocd with app-of-apps pattern](https://medium.com/dzerolabs/turbocharge-argocd-with-app-of-apps-pattern-and-kustomized-helm-ea4993190e7c)
- CI/CD
  - Review ArgoCD WorkFlow
- hardening:
  - [hardening_guide](https://rancher.com/docs/k3s/latest/en/security/hardening_guide/)
  - [kubescape](https://github.com/armosec/kubescape)
  - [kubescape docu](https://hub.armo.cloud/docs)
  - [YouTube](https://www.youtube.com/watch?v=ZATGiDIDBQk)


Won't do
--------

- CI/CD
  - Tekton
    - [Tekton Pipelines](https://tekton.dev/docs/getting-started/tasks/)
    - [Tekton Dashboard](https://tekton.dev/docs/dashboard/install/)
    - [Trigger](https://www.arthurkoziel.com/tutorial-tekton-triggers-with-github-integration/)
  - [argo-workflows](https://argoproj.github.io/argo-workflows/quick-start/)
  - [concourse](https://concourse-ci.org/)
    - Build and push containerimages
    - Build and push helm charts
  - [sealed-secrets](https://github.com/bitnami-labs/sealed-secrets)
Promethes
---------

Example:

```
# adds additional scrape configs to prometheus.yml
# must be a string so you have to add a | after extraScrapeConfigs:
# example adds prometheus-blackbox-exporter scrape config
extraScrapeConfigs:
  # - job_name: 'prometheus-blackbox-exporter'
  #   metrics_path: /probe
  #   params:
  #     module: [http_2xx]
  #   static_configs:
  #     - targets:
  #       - https://example.com
  #   relabel_configs:
  #     - source_labels: [__address__]
  #       target_label: __param_target
  #     - source_labels: [__param_target]
  #       target_label: instance
  #     - target_label: __address__
  #       replacement: prometheus-blackbox-exporter:9115
  ```