digilog
========

Small wrapper around the built in Go "fmt" class to add logging levels.

## Import

### go.mod
```
go get -u github.com/digicert/digilog
```

### go src
import "github.com/digicert/digilog"


## Usage

Setting the environment variable `LOG_LEVEL` will determine what level is written to the log. In your source code, you can use any of the following log levels: `DEBUG`, `INFO`, `WARN`, `ERROR`, `CRITICAL`. The default logging level is `INFO`.

By default, digilog writes logs to both `stdout` and `stderr`. This can be overridden by creating a BuffOut struct where Out and Err are set to `io.Writer` instances where the logs should be written.

In your own application, logs can be written as such:

```
log := digilog.New()
log.Info("my_event_name", "othervar=123", "yetothervar=234")
log.Infof("event_id=my_event_name %s %s", "othervar=123", "yetothervar=234")
```

To override the writer:

```
log := digilog.New()
log.SetOutput(&digilog.BuffOut{Out: fileWriter, Err: errFileWriter})
log.Info("my_event_name", "othervar=123", "yetothervar=234")
```

## Version History

2.0.0: Mimic golang standard logger package using a Log struct and an output buffer. Is *not* a drop in replacement for Go Logger.

1.2.0: Standardized funcs to mimic go's standard. If using Sprintf formatting, use the `f` suffixed method (Debugf, Infof, etc). Otherwise use the old methods

1.1.0: Updated `Critical` to behave like `log.Panic` accepting an err object instead of a string.

1.0.0: Initial release.
