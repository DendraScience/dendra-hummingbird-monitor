To set partition expiration, run the following sql statement:
```
ALTER SCHEMA `host_monitoring.host_metrics`
 SET OPTIONS(
     default_partition_expiration_days=31
 )
```
