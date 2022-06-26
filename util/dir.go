package util

import (
	"path"
	"runtime"
)

func GetCurrentAbPathByCaller(skip int) string {
	var abPath string
	_, filename, _, ok := runtime.Caller(skip)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
