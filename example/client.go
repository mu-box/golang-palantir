package main

import "github.com/nanobox-tools/palantir"
import "time"

func main() {
	log := palantir.PalantirClient{
		Host: "localhost",
		Port: "1234",
		Level: palantir.TRACE,
	}
	log.Fatal("fatal")
	log.Info("info")
	log.Error("error")
	log.Trace("trace")
	log.Warn("warn")
	log.Debug("debug")
	time.Sleep(10 * time.Second)
}