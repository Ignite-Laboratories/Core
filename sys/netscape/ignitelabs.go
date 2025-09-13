package netscape

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"git.ignitelabs.net/core/sys/netscape/igniteLabs"
)

var IgniteLabs _igniteLabs

type _igniteLabs byte

// Navigate executes a web server that listens on port 8080 (unless otherwise specified).
func (_igniteLabs) Navigate(port ...uint) {
	p := "8080"
	if len(port) > 0 {
		p = strconv.Itoa(int(port[0]))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, igniteLabs.Static)
		return
	})

	addr := ":" + p
	log.Printf("'ignitelabs.net' listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
