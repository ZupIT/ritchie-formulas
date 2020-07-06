## Manager bucket aws

### commands
List buckets
```bash
$ rit aws list bucket
```
Create buckets
```bash
$ rit aws create bucket
```
Delete buckets
```bash
$ rit aws delete bucket
```

### local test
```bash
$ make test-local form=AWS_BUCKET
```

### description
For use this formula run before:
```bash
rit set credential 
```
for add AWS credential.

For all formulas receive field in env REGION and has received (CREDENTIAL_AWS_ACCESSKEYID, CREDENTIAL_AWS_SECRETACCESSKEY) in envs (ACCESS_KEY, SECRET_ACCESS_KEY).

For "create" command receive bucket field in env BUCKET.