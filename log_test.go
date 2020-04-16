package zlog

import (
	"flag"
	"testing"
)

func TestLog(t *testing.T) {
	flag.Parse()
	InitLogger()

	Info("info")
}
