package main

import (
	"flag"
	"fmt"
	"goas/oas"
)

func main() {

	eGen := flag.String("gen", "-", "Generate entity files in this directory")
	dir := flag.String("dir", "-", "Directory where files are located.")
	out := flag.String("out", "", "Directory in which to write the generated file.")
	name := flag.String("name", "openapi", "result filename")
	extenion := flag.String("ext", "yaml", "extension result file")

	flag.Parse()

	if *eGen != "-" {
		oas.GenerateEntityFiles(*eGen)
	} else if *dir != "-" {
		oas.Convert(*dir, *out, *name, *extenion)
	} else {
		fmt.Println("expected params")
	}

}
