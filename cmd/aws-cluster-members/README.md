# aws-cluster-members

```
Create a go dir in your home directory then use go get to install.
package.

mkdir -p  go/{src,bin,pkg}

go get -v -u github.com/ChaosHour/aws-tools/cmd/aws-cluster-members


[klarsen@xxxxx ~]$ ls -1 go/bin
aws-cluster-members

```

### Usage
```
aws-cluster-members
Please provide a db-cluster-identifier


aws-cluster-members -h
Usage of aws-cluster-members:
  -d string
    	db-cluster-identifier
  -r string
    	Select a AWS Region to view you Clusters (default "us-west-1")


aws-cluster-members -d xxxx-us-stg
Cluster:	xxxx-us-stg
Instance:	xxxx-us-stg-2	False
Instance:	xxxx-us-stg-1	True
Instance:	xxxx-us-stg-0	False

```
True means the instance is read-write