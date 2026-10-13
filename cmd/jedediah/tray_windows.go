//go:build windows

package main

import (
	"context"
	_ "embed"

	"fyne.io/systray"
	"golang.org/x/sys/windows"
)

//go:embed winres/tray.ico
var trayIcon []byte

const hasTray = true

var procPostThreadMessage = windows.NewLazySystemDLL("user32.dll").NewProc("PostThreadMessageW")

// runTray shows the notification-area icon and blocks until ctx is done.
// Left-click opens the browser; the menu has Open and Quit. stop must cancel
// ctx and return only once Jedediah has shut down: systray also calls it when
// Windows ends the session, and the process may be killed right after.
//
// It must run on the main goroutine, which systray locks to the main thread.
func runTray(ctx context.Context, tooltip string, open, stop func()) {
	onReady := func() {
		systray.SetIcon(trayIcon)
		systray.SetTooltip(tooltip)
		systray.SetOnTapped(open)
		openItem := systray.AddMenuItem("Open Jedediah", "")
		systray.AddSeparator()
		quitItem := systray.AddMenuItem("Quit", "Stop Jedediah")
		go func() {
			for {
				select {
				case <-openItem.ClickedCh:
					open()
				case <-quitItem.ClickedCh:
					stop()
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	thread := windows.GetCurrentThreadId()
	go func() {
		<-ctx.Done()
		systray.Quit()
		// Ends the message loop even if systray could not create its window.
		const wmQuit = 0x0012
		procPostThreadMessage.Call(uintptr(thread), wmQuit, 0, 0)
	}()
	systray.Run(onReady, stop)
}
