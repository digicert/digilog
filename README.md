DIGILOG
========

Small wrapper around the built in Go "fmt" class to add logging levels.


## Import

### go.mod
```
require github.com/digicert/digilog v1.2.0 (<= whatever version here)
```

### go src
import "github.com/digicert/digilog"



## Usage

In your source code, you can use any of the following log levels: `DEBUG`, `INFO`, `WARN`, `ERROR`, `CRITICAL`.

Setting the environment variable `LOG_LEVEL` will determine what level is actually logged.

In the code, log levels can be used thusly:

`digilog.Warnf("event_id=%s msg='%s'", "some_warning_condition", "Got a warning.")`

OR:

`digilog.Critical(err)`


## Version History

1.2.0: Standardized funcs to mimic go's standard. If using Sprintf formatting, use the `f` suffixed method (Debugf, Infof, etc). Otherwise use the old methods

1.1.0: Updated `Critical` to behave like `log.Panic` accepting an err object instead of a string.

1.0.0: Initial release.
