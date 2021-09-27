# aws-tools

###
Usage
```
A wrapper around some aws-cli commands I use to list the logs from 
AWS RDS Aurora - AWS RDS  PostgreSQL.


Create a go dir in your home directory then use go get to install the aws_tools package.

mkdir -p  go/{src,bin,pkg}

go get -v -u github.com/ChaosHour/aws_tools

````


```
./aws-tools -h
Usage of /tmp/go-build784635128/b001/exe/main:
  -d string
    	Specify your db-instance-identifier to view it's logs.
  -l string
    	Log type to use Example: slow, error. (default "error")
  -r string
    	Select a aws region to view your clusters. (default "us-west-1")
````
