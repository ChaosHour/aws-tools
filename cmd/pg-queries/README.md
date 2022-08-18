# pg-queries

## Install
```
go get -v -u github.com/ChaosHour/aws-tools/cmd/pg-queries

For newer versions of go use:
go install github.com/ChaosHour/aws-tools/cmd/pg-queries@latest

```

### Note:  It helps if you have a .pgpass file in your home directory, but it is not required.

## USAGE
```
[klarsen@xxxxx ~]$ pg-queries 
Usage of pg-activity:
    -h string
        Specify your AWS EndPoint.
    -u string
        Specify your AWS User.

```

```SQL
pg-queries -u postgres  -h 10.8.0.11
Password for user postgres: 
Command Successfully Executed
-[ RECORD 1 ]-----------------------------------------------------------------------------------------------------
?column?    | 
query_start | 
usename     | 
state       | 
pid         | 47
query       | 
-[ RECORD 2 ]-----------------------------------------------------------------------------------------------------
?column?    | 
query_start | 
usename     | postgres
state       | 
pid         | 49
query       | 
-[ RECORD 3 ]-----------------------------------------------------------------------------------------------------
?column?    | 
query_start | 
usename     | 
state       | 
pid         | 45
query       | 
-[ RECORD 4 ]-----------------------------------------------------------------------------------------------------
?column?    | 
query_start | 
usename     | 
state       | 
pid         | 44
query       | 
-[ RECORD 5 ]-----------------------------------------------------------------------------------------------------
?column?    | 
query_start | 
usename     | 
state       | 
pid         | 46
query       | 
-[ RECORD 6 ]-----------------------------------------------------------------------------------------------------
?column?    | 00:00:00
query_start | 2022-08-16 05:06:19.443537+00
usename     | postgres
state       | active
pid         | 1520
query       | select now()-query_start,query_start,usename,state, pid, query from pg_stat_activity order by 1 desc


```
