package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	// Args hold command line arguments, starting with program name
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

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

	newfile, error := os.Create("index.html")
	if error != nil {
		log.Fatal("error creating file", error)
	}

	defer newfile.Close()
	io.Copy(newfile, strings.NewReader(str))
}
