package main

import "fmt"

func main() {
	name := "Yuta Lin"

	tpl := `
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
	`
	fmt.Println(tpl)
}