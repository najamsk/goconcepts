package main

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	tg := os.Getenv("tg")
	if tg == "" {
		t.Error("failed getting tg from env")
	}
	t.Logf("tg:%v \n", tg)

}
