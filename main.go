package main

import (
	"flag"
	"goas/oas"
)

func main() {
	dir := flag.String("dir", "", "Directory where files are located.")
	out := flag.String("out", "", "Directory in which to write the generated file.")
	name := flag.String("name", "openapi", "result filename")
	extenion := flag.String("ext", "yaml", "extension result file")
	oas.Convert(*dir, *out, *name, *extenion)
}
