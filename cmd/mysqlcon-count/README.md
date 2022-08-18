# mysqlcon-count

## Install
```
go get -v -u github.com/ChaosHour/aws-tools/cmd/mysqlcon-count

For newer versions of go use:
go install github.com/ChaosHour/aws-tools/cmd/mysqlcon-count@latest


```

### Note:  You must have a .my.cnf file with correct user and pass to use:
 ~/.my.cnf

## USAGE
```
mysqlcon-count -h app-stg.cluster-xxxxx.us-west-1.rds.amazonaws.com


mysqlcon-count

  -h string
    	Specify your AWS EndPoint.


```
```SQL
mysqlcon-count -h app-stg.cluster-xxxxx.us-west-1.rds.amazonaws.com
Command Successfully Executed
+---------------+--------------------------+----------+
| host_short    | users                    | COUNT(*) |
+---------------+--------------------------+----------+
| 10.2xx.xxx.xx | xxx_xxxxx                |        1 |
|               | system user              |        1 |
| localhost     | event_scheduler,rdsadmin |        6 |
| 10.23.xxx.x   | xxxx_app                 |        6 |
| 10.21x.xxx.x  | xxxxxxx_xxxx_srv         |       10 |
| 10.21x.xxx.x  | xxxxxxx_xxxx_srv         |       10 |
| 10.2x.xxx.xx  | xxxxxxx_xxxx_srv         |       10 |
| 10.2x.xxx.x   | xxxx                     |       10 |
| 10.2x.xxx.xx  | xxxx                     |       10 |
| 10.2x.xxx.xx  | xxxx_app                 |       14 |
| 10.2x.xxx.xx  | xxxxx_xxx                |       16 |
| 10.2x.xxx.xx  | xxxxx_xxx                |       16 |
| 10.2x.xxx.xx  | xxxxxx_xxx_app           |       16 |
| 10.2x.xxx.xx  | xxxx_app                 |       20 |
| 10.2x.xxx.xx  | xxxxxxx_xxxx             |       32 |
| total         |                          |      178 |
+---------------+--------------------------+----------+
```
