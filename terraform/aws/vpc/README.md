## terraform aws vpc

### command
```bash
$ rit terraform aws add vpc
```

### local test
```bash
$ make test-local form=terraform/aws/vpc
```

### description
This formula receives 5 inputs (region, vpc name, vpc cidr, vpc azs, customer name) and adds vpc module files into the project.