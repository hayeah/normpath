package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func help() {
	fmt.Println(
		`Usage: normpath <path>

Expand relative path to absolute path. Resolve symbolic links.
`)

	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid usage")
		help()
	}

	pathname := os.Args[1]

	abspath, err := filepath.Abs(pathname)
	if err != nil {
		log.Fatalln(err)
	}

	resolvedPath, err := filepath.EvalSymlinks(abspath)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resolvedPath)
}
