Aout this helmchart repository
===============================

Check:

```
helm lint ./helm/charts/prometheus_rules/*/
```

Packaging:

```
rm ./helm/charts/*.tgz
helm package ./helm/charts/*/
mv ./*.tgz ./helm/charts/
helm repo index --url https://olafradicke.github.io/own_dog_food/helm/charts/ ./helm/charts/
```

Use chart repo:

```
helm repo add own_dog_food https://olafradicke.github.io/own_dog_food/helm/charts/
helm repo update
helm search repo mutual_tls
helm install -f ./values.yaml -n monitoring myroute own_dog_food/mutual_tls
helm get all myroutes -n monitoring
```