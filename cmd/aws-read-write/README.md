# aws-read-write

```
Create a go dir in your home directory then use go get to install.
package.

mkdir -p  go/{src,bin,pkg}

go get -v -u github.com/ChaosHour/aws-tools/cmd/aws-read-write


[klarsen@xxxxx ~]$ ls -1 go/bin
aws-read-write
```


### Usage
```
[klarsen@xxxxx ~]$ aws-read-write
Executing command:
[aws rds describe-db-cluster-endpoints --db-cluster-identifier ]

{
    "DBClusterEndpoints": []
}



[klarsen@xxxxx ~]$ aws-read-write -h
Usage of aws-read-write:
  -d string
    	db-cluster-identifier Example: my-cluster-us-stg
  -r string
    	Select a aws region to view your clusters. (default "us-west-1")


[klarsen@xxxxx ~]$ aws-read-write -d my-cluster-us-stg
Executing command:
[aws rds --region us-west-1 describe-db-cluster-endpoints --db-cluster-identifier my-cluster-us-stg]

{
    "DBClusterEndpoints": [
        {
            "Status": "available",
            "Endpoint": "my-cluster-us-stg.cluster-xxxx.us-west-1.rds.amazonaws.com",
            "DBClusterIdentifier": "my-cluster-us-stg",
            "EndpointType": "WRITER"
        },
        {
            "Status": "available",
            "Endpoint": "my-cluster-us-stg.cluster-ro-xxxxx.us-west-1.rds.amazonaws.com",
            "DBClusterIdentifier": "my-cluster-us-stg",
            "EndpointType": "READER"
        }
    ]
}

```
