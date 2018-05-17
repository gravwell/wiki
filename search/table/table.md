## Table

The table renderer is used to create tables. Building tables is done by providing arguments to the table renderer. Arguments must be enumerated values, TAG, TIMESTAMP, or SRC.

### Sample Query

```
tag=syslog grep sshd | grep "Failed password for" | regex "Failed\spassword\sfor\s(?P<user>\S+)" | count by user | table user count
```

![](table-render.png)