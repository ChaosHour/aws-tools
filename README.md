# aws-tools

## Usage

```GO
A wrapper around some aws-cli commands I use to list, view and gather information 
about AWS resources. AWS RDS Aurora - AWS RDS  PostgreSQL.

```

###
What is needed to run this?
awscli, jq, psql, and MySQL client.
      

```BASH
On MacOS:
brew install jq
brew install awscli
brew install libpq
brew install mysql-client
```

### You will need the mysql-client and the pgsql-client installed to use these tools.




### What is in the cmd directory?

```GO
aws-cluster-members
Lists the members of an Aurora Cluster.

go install github.com/ChaosHour/aws-tools/cmd/aws-cluster-members@latest

aws-dns-query
Queries the AWS Route53 DNS service.

go install github.com/ChaosHour/aws-tools/cmd/aws-dns-query@latest

aws-read-write
Lists the read-write instances in an Aurora Cluster.

go install github.com/ChaosHour/aws-tools/cmd/aws-read-write@latest

my-repl-stat
Lists the replication status of an Aurora Cluster.

go install github.com/ChaosHour/aws-tools/cmd/my-repl-stat@latest

mysqlcon-count
Lists the number of MySQL connections.

go install github.com/ChaosHour/aws-tools/cmd/mysqlcon-count@latest

pg-activity
Lists the activity of PostgreSQL connections.

go install github.com/ChaosHour/aws-tools/cmd/pg-activity@latest

pg-queries
Lists the queries of PostgreSQL connections.

go install github.com/ChaosHour/aws-tools/cmd/pg-queries@latest
``` 