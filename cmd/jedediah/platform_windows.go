//go:build windows

package main

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Not in golang.org/x/sys/windows.
var (
	kernel32                  = windows.NewLazySystemDLL("kernel32.dll")
	procGetConsoleProcessList = kernel32.NewProc("GetConsoleProcessList")
	procFreeConsole           = kernel32.NewProc("FreeConsole")
)

// background is true when nobody can see stderr (started from Explorer, a
// shortcut or at sign-in), so fatal errors go to a dialog instead.
var background bool

// detachConsole makes a double-clicked Jedediah run without a console window.
// On Windows 11 24H2 and later the manifest (winres/jedediah.manifest) already
// does this, and a terminal launch keeps its console. Older Windows still opens
// a console, so drop it when no other process shares it: then Explorer started
// us, not a terminal.
func detachConsole() {
	if windows.RtlGetVersion().BuildNumber < 26100 {
		var pids [2]uint32
		n, _, _ := procGetConsoleProcessList.Call(uintptr(unsafe.Pointer(&pids[0])), uintptr(len(pids)))
		if n == 1 {
			procFreeConsole.Call()
			background = true
			return
		}
	}
	background = !stderrUsable()
}

func stderrUsable() bool {
	h, err := windows.GetStdHandle(windows.STD_ERROR_HANDLE)
	if err != nil || h == 0 || h == windows.InvalidHandle {
		return false
	}
	t, err := windows.GetFileType(h)
	return err == nil && t != windows.FILE_TYPE_UNKNOWN
}

// reportFatal prints err, and also shows it in a dialog when running in the background.
func reportFatal(err error) {
	fmt.Fprintln(os.Stderr, "jedediah:", err)
	if !background {
		return
	}
	text, _ := windows.UTF16PtrFromString(err.Error())
	title, _ := windows.UTF16PtrFromString("Jedediah")
	windows.MessageBox(0, text, title, windows.MB_OK|windows.MB_ICONERROR|windows.MB_SETFOREGROUND)
}
