// Command modzip packs a directory into a Go module zip using
// golang.org/x/mod/zip — the same packer proxy.golang.org uses, so the member
// set of the produced zip matches the published one by construction.
//
// It lives under .github/ so that it is outside the `./...` package pattern
// (the go tool skips dot-directories) and outside the packed tree.
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: modzip <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
