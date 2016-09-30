// rudes.me is my personal website
package main

import (
	"html/template"
	"io"
	"net/http"
	"time"
)

const (
	_staticURL   = "/static/"
	_staticRoot  = "/go/src/github.com/rudes/rudes.me/static/"
	_templateDir = "/go/src/github.com/rudes/rudes.me/static/templates/"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc(_staticURL, staticHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		render(w, "index.tmpl")
		return
	}
	http.NotFound(w, r)

}

func render(w http.ResponseWriter, content string) {
	var ctx string
	t, err := template.ParseFiles(_templateDir+"base.tmpl",
		_templateDir+"header.tmpl", _templateDir+content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, ctx)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	sf := r.URL.Path[len(_staticURL):]
	if len(sf) != 0 {
		f, err := http.Dir(_staticRoot).Open(sf)
		if err != nil {
			f.Close()
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		content := io.ReadSeeker(f)
		http.ServeContent(w, r, sf, time.Now(), content)
	}
}
