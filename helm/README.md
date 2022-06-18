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
```

Indexing:

```
helm repo index --url https://olafradicke.github.io/own_dog_food/helm/charts/ ./helm/charts/
```

Use chart repo:

```
helm repo add own_dog_food https://olafradicke.github.io/own_dog_food/helm/charts/
helm search repo prometheus_rules
```