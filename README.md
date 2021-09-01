# aws_tools

###
Usage
```
Create a go dir in your home directory then use go get to install the gds_logs package.

mkdir -p  go/{src,bin,pkg}

go get -v -u github.com/chaoshour/aws_tools

````


```
./aws_tools -h
Usage of /tmp/go-build784635128/b001/exe/main:
  -dbi string
    	Specify your db-instance-identifier to view it's logs.
  -l string
    	Log type to use Example: slow, error. (default "error")
  -r string
    	Select a aws region to view your clusters. (default "us-west-1")
````
