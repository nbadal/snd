//go:build darwin

package main

import (
	"os"
	"syscall"
)

func ignoredSignals() []os.Signal {
	return []os.Signal{syscall.SIGURG}
}
