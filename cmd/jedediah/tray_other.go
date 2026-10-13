//go:build !windows

package main

import "context"

const hasTray = false

func runTray(ctx context.Context, tooltip string, open, stop func()) { <-ctx.Done() }
