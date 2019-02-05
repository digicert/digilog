DIGILOG
========

Small wrapper around the built in Go "fmt" class to add logging levels.


## Import

### go.mod
```
require github.com/digicert/digilog v1.0
```

### go src
import "github.com/digicert/digilog"



## Usage

In your source code, you can use any of the following log levels: `DEBUG`, `INFO`, `WARN`, `ERROR`, `CRITICAL`.

Setting the environment variable `LOG_LEVEL` will determine what level is actually logged.

In the code, log levels can be used thusly:

`digilog.Warn("event_id=%s msg='%s'", "some_warning_condition", "Got a warning.")`