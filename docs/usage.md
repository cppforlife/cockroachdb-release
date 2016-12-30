# Usage

Asssuming you've deployed [manifests/example.yml](../manifests/example.yml)

## Create test database

```
$ bosh -d cockroachdb ssh cockroachdb_root/0
$ sudo su
# /var/vcap/jobs/cockroachdb_user/bin/cli sql

root@:26257> CREATE DATABASE bank;
CREATE DATABASE

root@:26257> SET DATABASE = bank;
SET

root@:26257> CREATE TABLE accounts (id INT PRIMARY KEY, balance DECIMAL);
CREATE TABLE

root@26257> INSERT INTO accounts VALUES (1234, 10000.50);
INSERT 1

root@26257> SELECT * FROM accounts;
```

Useful links:

- [SQL feature support](https://www.cockroachlabs.com/docs/sql-feature-support.html)

## Import example data

```
$ bosh -d cockroachdb ssh cockroachdb_root/0
$ sudo su
# source /var/vcap/jobs/cockroachdb_user/bin/env
# cockroach gen example-data | cockroach sql
```

## Scale up

Observe cluster scaling up to 5 nodes:

```
$ bosh -d cockroachdb deploy manifests/example.yml -o manifests/scale.yml -v num=5 --vars-store ./creds.yml
```
