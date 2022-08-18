# my-repl-stat

## Install
```
go get -v -u github.com/ChaosHour/aws-tools/cmd/my-repl-stat

```

### Note:  You must have a .my.cnf file with correct GDS user and pass to use:

## USAGE
```
[klarsen@xxxxx ~]$ my-repl-stat
  -h string
    	Specify your AWS EndPoint.


[klarsen@xxxxx ~]$ my-repl-stat -h
flag needs an argument: -h
Usage of my-repl-stat:
  -h string
    	Specify your AWS EndPoint.


```

```SQL
[klarsen@xxxxx ~]$ my-repl-stat -h app-prod-1.xxxxx.eu-west-1.rds.amazonaws.com
Command Successfully Executed
                  Master_Host: app-prod-2.xxxxx.eu-west-1.rds.amazonaws.com
             Slave_IO_Running: Yes
            Slave_SQL_Running: Yes
        Seconds_Behind_Master: 0

```
