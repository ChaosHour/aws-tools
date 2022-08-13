# aws-tools

## Usage

```GO
A wrapper around some aws-cli commands I use to list, view and gather information 
about AWS resources. AWS RDS Aurora - AWS RDS  PostgreSQL.

```

###
What is needed to run this?
awscli and jq

```BASH
On MacOS:
brew install jq
brew install awscli
```


### What is in the cmd directory?

```GO
aws-cluster-members
Lists the members of an Aurora Cluster.

aws-dns-query
Queries the AWS Route53 DNS service.


aws-rds-describe
Describes an AWS RDS instance.

aws-read-write
Lists the read-write instances in an Aurora Cluster.
``` 