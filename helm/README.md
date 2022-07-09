About this helmchart repository
===============================

Check:

```
helm lint ./helm/charts/*/*/
```

Packaging:

```
helm lint ./helm/charts/*/
# rm ./helm/charts/*.tgz
helm package ./helm/charts/*/
mv ./*.tgz ./helm/charts/
helm repo index --url https://olafradicke.github.io/own_dog_food/helm/charts/ ./helm/charts/
```

Use chart repo:

```
helm repo add own_dog_food https://olafradicke.github.io/own_dog_food/helm/charts/
helm repo update
helm search repo mutual_tls
helm install \
  -f ./values.yaml \
  -n monitoring  \
  --dry-run \
  --debug  \
  --version 0.1.2  \
  myroutes \
  own_dog_food/mutual_tls

helm install -f ./values.yaml -n monitoring mutual_tls own_dog_food/mutual_tls
helm get all mutual_tls -n monitoring
```
