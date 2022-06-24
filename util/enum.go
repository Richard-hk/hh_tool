package util

import "os"

const (
	V2RAY = "v2ray"
)

var MonitorApp = []string{V2RAY}
var StopSignal = make(chan os.Signal, 1)
