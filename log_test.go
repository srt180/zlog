package zlog

import (
	"flag"
	"testing"
)

func TestLog(t *testing.T) {
	flag.Parse()
	InitLogger()

	Info("info")
	WithField("with", "field").Debug("withfield debug")
	With("just", "with").Debugf("just with debugf:%d", 2020)
}
