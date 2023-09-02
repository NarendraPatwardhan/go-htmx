//go:build windows
// +build windows

package context

import (
	"os"
)

var terminationSignals = []os.Signal{os.Interrupt}
