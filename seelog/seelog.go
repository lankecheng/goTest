package main

import (
	//"fmt"
	"github.com/cihub/seelog"
)

func main() {
	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		panic(err.Error())
	}
	seelog.ReplaceLogger(logger)

	defer seelog.Flush()
	i := []string{"a", "b", "c"}
	seelog.Tracef("test format %v", i)
	seelog.Debug("Hello from seelog debug")
	seelog.Info("Hello from seelog info")
	seelog.Warn("Hello from seelog warn")
	seelog.Error("Hello from seelog error")
	seelog.Critical("Hello from seelog debug")
}
