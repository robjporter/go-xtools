package main

import (
	"../xlogger"
)

// TODO:

func main() {
	//log := xlogger.NewLogger()
	//log.SetOutput(os.Stdout)
	//log.SetFormat("{TIME} <{TAG}>: {MSG}")
	//log.SetTimeFormat("Mon Jan _2 15:04:05 2006")
	xlogger.Info("INFO","number",4)
	xlogger.Debug("DEBUG")
	xlogger.Warn("WARN")
	xlogger.Error("ERROR")
	xlogger.Infof("INFO 2",xlogger.NewTag("GAG").Get(), 1)
}