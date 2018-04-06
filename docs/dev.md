## Dev

Uses `-o manifests/dev.yml` to automatically rebuild release:

```
$ bosh -e vbox -d cockroachdb deploy manifests/example.yml -o manifests/dev.yml -n --vars-store ./creds.yml
```

Perf issues:

- https://github.com/cockroachdb/cockroach/issues/17777
  - CockroachDB v1.0 write performance is half as fast as postgreSQL?
- https://github.com/cockroachdb/cockroach/issues/24507
