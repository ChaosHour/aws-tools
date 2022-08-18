# pg-activity

## Install
```
go get -v -u github.com/ChaosHour/aws-tools/cmd/pg-activity

For newer versions of go use:
go install github.com/ChaosHour/aws-tools/cmd/pg-activity@latest


```

### Note:  It helps if you have a .pgpass file in your home directory, but it is not required.

## USAGE
```
[klarsen@xxxxx ~]$ pg-activity 
  -h string
        Specify your AWS EndPoint.
  -u string
        Specify your AWS User.
```

```SQL
pg-activity -u postgres  -h 10.8.0.11
Password for user postgres: 
Command Successfully Executed
-[ RECORD 1 ]-----
usename | postgres
state   | 
count   | 1
-[ RECORD 2 ]-----
usename | 
state   | 
count   | 4
```
