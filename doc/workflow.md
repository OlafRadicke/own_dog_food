WORKFLOWS
=========

Log in over token
-----------------


Get auth token of web-user

```
$ kubectl -n argo get secret web-user.service-account-token -o=jsonpath='{.data.token}' | base64 --decode
```

For use as login in web ui you have to add the **prefix "Bearer"**. For example `Bearer ZXlKaGJHY2lPaUpTVXpJMU5pSXNJbXRwWkNJNkltS...`

Details see [docu](https://argo-workflows.readthedocs.io/en/latest/access-token/#token-creation)