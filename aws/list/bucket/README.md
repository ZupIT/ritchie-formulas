# List bucket aws

## command

List buckets

```bash
rit aws list bucket
```

## description

For use this formula run before:

```bash
rit set credential
```

for add AWS credential.

For all formulas receive field in env REGION and has received (CREDENTIAL_AWS_ACCESSKEYID,
CREDENTIAL_AWS_SECRETACCESSKEY) in envs (ACCESS_KEY, SECRET_ACCESS_KEY).

For "create" command receive bucket field in env BUCKET.
