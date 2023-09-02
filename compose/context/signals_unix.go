//go:build !windows
// +build !windows

package context

import (
	"os"

	"golang.org/x/sys/unix"
)

var terminationSignals = []os.Signal{unix.SIGTERM, unix.SIGINT}
