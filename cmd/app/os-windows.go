//go:build windows

package main

import (
	"os"
)

func ignoredSignals() []os.Signal {
	return []os.Signal{} // Does not have SIGURG
}
