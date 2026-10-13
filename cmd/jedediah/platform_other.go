//go:build !windows

package main

import (
	"fmt"
	"os"
)

func detachConsole() {}

func reportFatal(err error) {
	fmt.Fprintln(os.Stderr, "jedediah:", err)
}
