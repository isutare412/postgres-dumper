- [Postgres Dumper](#postgres-dumper)
- [Usage](#usage)
- [Deployment (CronJob)](#deployment-cronjob)

# Postgres Dumper

Backup your PostgreSQL to AWS S3 periodically. You can use linux crontab or k8s
[cronjob](#deployment-cronjob).

This repository was copied from a backup module of [Web
Memo](https://github.com/isutare412/web-memo) due to its feature
independentness.

# Usage

```bash
# Prepare your own config
cat <<EOF >> config.local.yaml
process-timeout: 10m
log:
  format: <format> # text / json
  level: <level> # debug / info / warn / error
aws:
  access-key: <access_key>
  secret: <secret_access_key>
  region: <aws_region>
  s3:
    backup-bucket: <bucket>
backup:
  prefix: <object_prefix>
  retention: 720h # 30 days
postgres:
  host: localhost
  port: 5432
  user: <user>
  password: <password>
  database: <database_name>
EOF

make run
```

# Deployment (CronJob)

```bash
# Prepare your own values
cat <<EOF >> my_values.yaml
config:
  aws:
    access-key: <access_key>
    secret: <secret_access_key>
    region: <region>
# ...
EOF

helm upgrade --install postgres-dumper ./helm/postgres-dumper -f my_values.yaml
```
