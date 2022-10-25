package log

import "testing"

func TestInfo(t *testing.T) {
	Info("hello info")

	Infof("hello %s info", "log")
}

func TestError(t *testing.T) {
	Error("hello error")
}
