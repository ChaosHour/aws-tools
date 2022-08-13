# aws-cluster-members

```
Create a go dir in your home directory then use go get to install.
package.

mkdir -p  go/{src,bin,pkg}

go get -v -u github.com/ChaosHour/aws-tools/cmd/aws-cluster-members


[klarsen@xxxxx ~]$ ls -lrt go/bin
total 20820
-rwxr-xr-x 1 klarsen xxx  2491007 Sep 16 16:25 gds-logs
-rwxr-xr-x 1 klarsen xxx  2499399 Sep 16 16:25 download-logs
-rwxr-xr-x 1 klarsen xxx 11329531 Sep 21 22:34 sync-rds-logs
-rwxr-xr-x 1 klarsen xxx  2490841 Sep 22 21:19 mysqlcon-count
-rwxr-xr-x 1 klarsen xxx  2495266 Sep 23 17:50 aws-read-write
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