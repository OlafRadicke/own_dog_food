#!/bin/bash
helm lint ./helm/chart_src/*/
# rm ./helm/charts/*.tgz
helm package ./helm/chart_src/*/
mv ./*.tgz ./helm/charts/
helm repo index --url https://olafradicke.github.io/own_dog_food/helm/charts/ ./helm/charts/