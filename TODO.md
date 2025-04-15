Todos
=====

- Use own hugo operator for tif page.
- [Check ssl](https://www.ssllabs.com/ssltest/?ref=traefik.io)
- https://www.pulumi.com/registry/packages/kubernetes/api-docs/helm/v3/chart/
   - [HTTPS with Cert-Manager and Letsencrypt](https://k3s.rocks/https-cert-manager-letsencrypt/)
- Grafana:
  - https://github.com/grafana-operator/grafana-operator/tree/v4.8.0/documentation
- Loki https://loki-operator.dev/docs/prologue/introduction.md/
- hardening:
  - [hardening_guide](https://rancher.com/docs/k3s/latest/en/security/hardening_guide/)
  - [kubescape](https://github.com/armosec/kubescape)
  - [kubescape docu](https://hub.armo.cloud/docs)
  - [YouTube](https://www.youtube.com/watch?v=ZATGiDIDBQk)


Won't do
--------

- (Grafana) backup https://velero.io/
- CI/CD
  - Tekton
    - [Tekton Pipelines](https://tekton.dev/docs/getting-started/tasks/)
    - [Tekton Dashboard](https://tekton.dev/docs/dashboard/install/)
    - [Trigger](https://www.arthurkoziel.com/tutorial-tekton-triggers-with-github-integration/)
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