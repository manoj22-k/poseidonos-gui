package globals

import (
	"github.com/shettyh/tlock"
	"kouros/model"
	"sync"
	"time"
)

var Telemetry_config_path = "/usr/local/dagent/telemetry_config.yaml"
var TelemetryWriteLock sync.Mutex
var InitialTime = int64(0)
var InitialRes model.Response
var InitialErr error
var TimeLimit = int64(8)            //seconds
var LockTimeout = time.Duration(10) //seconds
var APILock tlock.Lock = tlock.New()