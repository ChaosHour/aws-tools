# aws-rds-describe

## Usage

```GO
Create a go dir in your home directory then use go get to install the aws_tools package.

mkdir -p  go/{src,bin,pkg}

go get -v -u github.com/ChaosHour/aws-tools/cmd/aws-rds-describe

```

```GO
aws-rds-describe -h
Usage
  -d string
    Specify your db-instance-identifier to view it's logs.
  -l string
    Log type to use Example: slow, error. (default "error")
  -r string
    Select a aws region to view your clusters. (default "us-west-1")

```
