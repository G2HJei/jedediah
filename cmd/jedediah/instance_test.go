package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/G2HJei/job-search-tracker/internal/web"
)

func TestBrowserAddr(t *testing.T) {
	for in, want := range map[string]string{
		"127.0.0.1:8765": "127.0.0.1:8765",
		":8765":          "127.0.0.1:8765",
		"0.0.0.0:8765":   "127.0.0.1:8765",
		"[::]:8765":      "127.0.0.1:8765",
		"[::1]:8765":     "[::1]:8765",
		"localhost:8765": "localhost:8765",
	} {
		if got := browserAddr(in); got != want {
			t.Errorf("browserAddr(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestRunningAt(t *testing.T) {
	jedediah := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(web.VersionHeader, "test")
	}))
	defer jedediah.Close()
	other := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer other.Close()
	gone := httptest.NewServer(http.NotFoundHandler())
	gone.Close()

	if !runningAt(jedediah.URL, 0) {
		t.Error("Jedediah not detected")
	}
	if runningAt(other.URL, time.Second) {
		t.Error("another program's server taken for Jedediah")
	}
	if runningAt(gone.URL, 0) {
		t.Error("closed port taken for Jedediah")
	}
}
