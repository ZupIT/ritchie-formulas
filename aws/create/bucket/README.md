# Create bucket aws

## command

Create bucket

```bash
rit aws create bucket
```

## description

For use this formula run before:

```bash
rit set credential
```

for add AWS credential.

Formula receive field in env REGION and has received
(CREDENTIAL_AWS_ACCESSKEYID, CREDENTIAL_AWS_SECRETACCESSKEY) in envs
(ACCESS_KEY, SECRET_ACCESS_KEY).

Formula receive bucket field in env BUCKET.
