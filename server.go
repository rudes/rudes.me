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
	tFuncs := template.FuncMap{
		"printLanguages": printLangs,
	}
	t, err := template.New("base.tmpl").Funcs(tFuncs).ParseFiles(_templateDir+"base.tmpl",
		_templateDir+"header.tmpl", _templateDir+content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, ctx)
}

func printLangs() template.HTML {
	var res string

	type language struct {
		Language, Link string
	}
	langs := []language{
		{"C", "https://en.wikipedia.org/wiki/The_C_Programming_Language"},
		{"C++", "https://en.wikipedia.org/wiki/C%2B%2B"},
		{"C#", "https://en.wikipedia.org/wiki/C_Sharp_(programming_language)"},
		{"Go", "https://golang.org/"},
		{"Java", "https://en.wikipedia.org/wiki/Java_(programming_language)"},
		{"Ruby", "https://www.ruby-lang.org/en/"},
		{"Bash", "https://en.wikipedia.org/wiki/Bash_(Unix_shell)"},
		{"Powershell", "https://en.wikipedia.org/wiki/PowerShell"},
		{"Perl", "https://en.wikipedia.org/wiki/Perl"},
		{"Python", "https://en.wikipedia.org/wiki/Python_(programming_language)"},
		{"Rust", "https://www.rust-lang.org/en-US/"},
		{"Haskell", "https://www.haskell.org/"},
		{"TCL", "https://en.wikipedia.org/wiki/Tcl"},
		{"VimL", "https://en.wikipedia.org/wiki/Vim_(text_editor)#Vim_script"},
		{"JavaScript", "https://www.javascript.com/"},
		{"KSH", "https://en.wikipedia.org/wiki/Korn_shell"},
		{" ", " "},
		{"PHP", "https://secure.php.net/"},
	}
	for i := 0; i < len(langs); i++ {
		res += "<tr>"
		res += "<td><a href='" + langs[i].Link + "'>" + langs[i].Language + "</a></td>"
		i++
		res += "<td><a href='" + langs[i].Link + "'>" + langs[i].Language + "</a></td>"
		i++
		res += "<td><a href='" + langs[i].Link + "'>" + langs[i].Language + "</a></td>"
		res += "</tr>"
	}
	return template.HTML(res)

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
