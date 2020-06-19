package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	InitLogger()
	Debug("Useful debugging information.")
	Info("Something noteworthy happened!")
	Warn("You should probably take a look at this.")
	Error("Something failed but I'm not quitting.")
	Fatal("Bye.")   //log之后会调用os.Exit(1)
	Panic("I'm bailing.")   //log之后会panic()
}

