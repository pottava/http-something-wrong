package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pottava/http-something-wrong/lib"
)

var (
	version string
	date    string
)

func main() {
	http.Handle("/", lib.Wrap(index))
	http.Handle("/sleep/", lib.Wrap(sleep))
	http.Handle("/errors/", lib.Wrap(errors))

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		if len(version) > 0 && len(date) > 0 {
			fmt.Fprintf(w, "version: %s (built at %s)\n", version, date)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	log.Printf("[service] listening on port %s", lib.Config.Port)
	log.Fatal(http.ListenAndServe(":"+lib.Config.Port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html><html lang="en">
<head><meta charset="utf-8"><title>something wrong</title></head><body>
<h3>Sleep</h3>
<ul>
  <li><a href="/sleep/50">50ms</a></li>
  <li><a href="/sleep/100">100ms</a></li>
  <li><a href="/sleep/1000">1s</a></li>
  <li><a href="/sleep/5000">5s</a></li>
  <li>..</li>
</ul>
<h3>HTTP Errors</h3>
<ul>
  <li><a href="/errors/400">Bad Request</a></li>
  <li><a href="/errors/401">Unauthorized</a></li>
  <li><a href="/errors/403">Forbidden</a></li>
  <li><a href="/errors/404">Not Found</a></li>
  <li><a href="/errors/429">Too Many Requests</a></li>
  <li><a href="/errors/500">Internal Server Error</a></li>
  <li><a href="/errors/501">Not Implemented</a></li>
  <li><a href="/errors/502">Bad Gateway</a></li>
  <li><a href="/errors/503">Service Unavailable</a></li>
  <li>..</li>
</ul>
</body></html>`)
}

func sleep(w http.ResponseWriter, r *http.Request) {
	var duration int64
	if candidate, err := strconv.ParseInt(r.URL.Path[len("/sleep/"):], 10, 64); err == nil {
		duration = candidate
	}
	time.Sleep(time.Duration(duration) * time.Millisecond)
	fmt.Fprintf(w, "%d\n", duration)
}

func errors(w http.ResponseWriter, r *http.Request) {
	code := http.StatusBadRequest
	if candidate, err := strconv.ParseInt(r.URL.Path[len("/errors/"):], 10, 32); err == nil {
		code = int(candidate)
	}
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%d %s\n", code, http.StatusText(code))
}
