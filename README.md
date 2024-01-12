Playbook for hosteurope server setup
====================================

Mission
-------

*"Eating your (my) own dog food"* :sweat_smile:



Change workflow
---------------

- Change helm chart (if necessary )
- Rebuild chart and reposytory (see [helm/README.md](helm/README.md))
- Commit changes (include add new helm tgz file)
- Push changes
- Change playbook
  - Change helm chart version (if necessary )
  - Change helm values (if necessary )
- Run playbook

***Note:*** the helm chart repository is maped onliy on the develop branch!

EXTENDED DOKUMENTATION
----------------------

See [doc/_index.md](doc/_index.md)


Links
-----

- [bitnami/argo-workflows](https://github.com/bitnami/charts/tree/master/bitnami/argo-workflows)
- [helm repos on github](https://medium.com/@mattiaperi/create-a-public-helm-chart-repository-with-github-pages-49b180dbb417)
- [concourse chart](https://github.com/concourse/concourse-chart)
- [tekton-pipeline](https://github.com/cdfoundation/tekton-helm-chart