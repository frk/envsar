package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	leftDelim  = "{{"
	rightDelim = "}}"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no template provided")
	}

	file := os.Args[1]
	name := filepath.Base(file)
	fmap := makefuncmap()
	tmpl := template.New(name).Delims(leftDelim, rightDelim).Funcs(fmap)

	var err error
	if tmpl, err = tmpl.ParseFiles(file); err != nil {
		log.Fatal(err.Error())
	}
	if err := tmpl.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err.Error())
	}
}

func makefuncmap() template.FuncMap {
	environ := os.Environ()
	funcmap := make(template.FuncMap)
	for _, kv := range environ {
		if i := strings.Index(kv, "="); i > -1 {
			k, v := kv[:i], kv[i+1:]
			funcmap[k] = envfunc(v)
		}
	}
	return funcmap
}

func envfunc(v string) (fn func() string) {
	return func() string {
		return v
	}
}