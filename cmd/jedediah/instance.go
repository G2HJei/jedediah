package main

import (
	"net"
	"net/http"
	"time"

	"github.com/G2HJei/job-search-tracker/internal/web"
)

// browserAddr turns a listen address into one a browser can open.
func browserAddr(addr string) string {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return addr
	}
	if ip := net.ParseIP(host); host == "" || ip != nil && ip.IsUnspecified() {
		host = "127.0.0.1"
	}
	return net.JoinHostPort(host, port)
}

// runningAt reports whether Jedediah is serving url, waiting up to wait for it:
// a running instance binds its port before it loads the data folder, and only
// answers once that is done.
func runningAt(url string, wait time.Duration) bool {
	client := &http.Client{Timeout: time.Second}
	deadline := time.Now().Add(wait)
	for {
		if resp, err := client.Get(url + "/healthz"); err == nil {
			resp.Body.Close()
			return resp.StatusCode == http.StatusOK && resp.Header.Get(web.VersionHeader) != ""
		}
		if time.Now().After(deadline) {
			return false
		}
		time.Sleep(200 * time.Millisecond)
	}
}
