package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Yuta Lin"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang = "en">
	<head>
	<meta charset = "UTF-8">
	<title>Hello Go!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	// os.Create: func Create(name string) (*File, error)
	newfile, error := os.Create("index.html")
	if error != nil {
		log.Fatal("error creating file", error)
	}

	defer newfile.Close()

	// io.Copy: func Copy(dst Writer, src Reader)(written int64, error error)
	// strings.NewReader: func NewReader(s string) *Reader
	io.Copy(newfile, strings.NewReader(str))
}
