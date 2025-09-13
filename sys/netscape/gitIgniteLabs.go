package netscape

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// GitIgniteLabs drives the git.ignitelabs.net service, which acts as a "vanity URL" for Go imports.
//
// The system is very simple.  All access to the URI https://git.ignitelabs.net/module is redirected towards
// https://github.com/ignite-laboratories/module.  To support vanity imports, you must ALSO print out an HTML page
// when Go requests https://git.ignitelabs.net/module?go-get=1
//
//	<html>
//	  <head>
//	    <meta name="go-import" content="[importPath] git [remote]">
//	    <meta name="go-source" content="[importPath] [remote] [remote]/tree/HEAD{/dir} [remote]/blob/HEAD{/dir}/{file}#L{line}">
//	  </head>
//	  <body>OK</body>
//	</html>
//
// That's really it!  No fany libraries are needed, just a simple HTTP handler =)
//
// NOTE: The address 'git.ignitelabs.net' is implied through the request and not present in the actual code.
var GitIgniteLabs _gitIgniteLabs

type _gitIgniteLabs byte

// Navigate executes a web server that listens on port 8080 (unless otherwise specified).
func (_gitIgniteLabs) Navigate(port ...uint) {
	p := "8080"
	if len(port) > 0 {
		p = strconv.Itoa(int(port[0]))
	}

	metaTmpl := template.Must(template.New("meta").Parse(`<!doctype html>
<html><head>
<meta name="go-import" content="{{.ImportPath}} git {{.Remote}}">
<meta name="go-source" content="{{.ImportPath}} {{.Remote}} {{.Remote}}/tree/HEAD{/dir} {{.Remote}}/blob/HEAD{/dir}/{file}#L{line}">
</head><body>OK</body></html>`))

	type metaData struct {
		ImportPath string // e.g. git.ignitelabs.net/enigmaneering/enigma0
		Remote     string // e.g. https://github.com/ignite-laboratories/enigmaneering/enigma0
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		repo := "https://github.com/ignite-laboratories" + r.URL.Path

		// Go toolchain probe: serve meta tags (no redirect).
		if r.URL.Query().Get("go-get") == "1" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if err := metaTmpl.Execute(w, metaData{
				ImportPath: r.Host + r.URL.Path,
				Remote:     repo,
			}); err != nil {
				http.Error(w, "template error", http.StatusInternalServerError)
			}
			return
		}

		// Browser: redirect to GitHub. Use tree/HEAD for directories; it’s fine for blobs too.
		http.Redirect(w, r, repo, http.StatusFound)
	})

	addr := ":" + p
	log.Printf("'git.ignitelabs.net' listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
