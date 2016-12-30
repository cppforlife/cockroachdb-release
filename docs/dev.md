## Dev

Uses `-o manifests/dev.yml` to automatically rebuild release:

```
$ bosh -e vbox -d cockroachdb deploy manifests/example.yml -o manifests/dev.yml -n --vars-store ./creds.yml
```
