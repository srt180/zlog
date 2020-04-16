# zlog

```
flags:
  -log.age int
        MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename. (default 7)
  -log.backups int
        MaxBackups is the maximum number of old log files to retain. (default 5)
  -log.filename string
        log file name (default "default")
  -log.level string
        log levels:debug/info/warn/error/dpanic/panic/fatal (default "info")
  -log.path string
        log save path (default "/tmp")
  -log.size int
        MaxSize is the maximum size in megabytes of the log file before it gets rotated. It defaults to 10 megabytes. (default 10)
```
