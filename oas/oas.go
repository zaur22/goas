package oas

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func Convert(from string, to string, fileName string, extension string) {
	fileName += "." + extension
	OAS := readAllFrom(from)
	writeAllTo(to, fileName, OAS)
}

type oas struct {
	info            string
	paths           []string
	schemas         []string
	parameters      []string
	securitySchemes []string
	requestBodies   []string
	responses       []string
	headers         []string
	examples        []string
	links           []string
	callbacks       []string
}

const (
	INFO             = "info"
	PATHS            = "paths"
	SCHEMAS          = "schemas"
	PARAMETERS       = "parameters"
	SECURITY_SCHEMES = "securitySchemas"
	REQUEST_BODIES   = "requestBodies"
	RESPONSES        = "responses"
	HEADERS          = "headers"
	EXAMPLES         = "examples"
	LINKS            = "links"
	CALLBACKS        = "callbacks"

	COMPONENTS = "components"
	END        = ":\n"
)

var extensions = []string{".yaml", ".yml"}

func readAllFrom(from string) oas {
	var OAS = oas{}
	files, err := ioutil.ReadDir(from)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			var dir = path.Join(from, f.Name())
			readFromEntityPath(&OAS, dir)
		} else {
			var filename = f.Name()
			var extension = filepath.Ext(filename)
			var name = filename[0 : len(filename)-len(extension)]
			if name == INFO && stringInSlice(extension, extensions) {
				info, err := ioutil.ReadFile(path.Join(from, filename))
				if err != nil {
					log.Fatal(err)
				}
				OAS.info = string(info)
			}
		}
	}
	return OAS
}

func readFromEntityPath(OAS *oas, dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileValByte, err := ioutil.ReadFile(path.Join(dir, f.Name()))
		if err != nil {
			log.Fatal(err)
		}

		if len(fileValByte) == 0 {
			continue
		}

		var fileVal = string(fileValByte)
		var filename = f.Name()
		var extension = filepath.Ext(filename)
		var name = filename[0 : len(filename)-len(extension)]
		if !stringInSlice(extension, extensions) {
			log.Fatalf("Bad extension %v, expected %v, got %v", filename, extensions, extension)
		}

		var fileValRune = []rune(fileVal)
		if len(fileValRune) != 0 {
			var lastSymb = string(
				fileValRune[len(fileValRune)-1:],
			)

			if lastSymb != "\n" {
				fileVal += "\n"
			}
		}

		switch name {
		case PATHS:
			OAS.paths = append(OAS.paths, fileVal)
		case SCHEMAS:
			OAS.schemas = append(OAS.schemas, fileVal)
		case PARAMETERS:
			OAS.parameters = append(OAS.parameters, fileVal)
		case SECURITY_SCHEMES:
			OAS.securitySchemes = append(OAS.securitySchemes, fileVal)
		case REQUEST_BODIES:
			OAS.requestBodies = append(OAS.requestBodies, fileVal)
		case RESPONSES:
			OAS.responses = append(OAS.responses, fileVal)
		case HEADERS:
			OAS.headers = append(OAS.headers, fileVal)
		case EXAMPLES:
			OAS.examples = append(OAS.examples, fileVal)
		case LINKS:
			OAS.links = append(OAS.links, fileVal)
		case CALLBACKS:
			OAS.callbacks = append(OAS.callbacks, fileVal)
		}
	}
}

func writeAllTo(to string, filename string, OAS oas) {

	file, err := os.OpenFile(
		path.Join(to, filename),
		os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var res = formatOAS(OAS)
	file.WriteString(res)
}

func formatOAS(OAS oas) string {
	var res = ""

	res += OAS.info
	res += PATHS + END
	for _, str := range OAS.paths {
		res += tab(str, 1)
	}

	res += COMPONENTS + END

	if len(OAS.schemas) != 0 {
		res += tab(SCHEMAS+END, 1)
		for _, str := range OAS.schemas {
			res += tab(str, 2)
		}
	}

	if len(OAS.parameters) != 0 {
		res += tab(PARAMETERS+END, 1)
		for _, str := range OAS.parameters {
			res += tab(str, 2)
		}
	}

	if len(OAS.securitySchemes) != 0 {
		res += tab(SECURITY_SCHEMES+END, 1)
		for _, str := range OAS.securitySchemes {
			res += tab(str, 2)
		}
	}

	if len(OAS.requestBodies) != 0 {
		res += tab(REQUEST_BODIES+END, 1)
		for _, str := range OAS.requestBodies {
			res += tab(str, 2)
		}
	}

	if len(OAS.responses) != 0 {
		res += tab(RESPONSES+END, 1)
		for _, str := range OAS.responses {
			res += tab(str, 2)
		}
	}

	if len(OAS.headers) != 0 {
		res += tab(HEADERS+END, 1)
		for _, str := range OAS.headers {
			res += tab(str, 2)
		}
	}

	if len(OAS.examples) != 0 {
		res += tab(EXAMPLES+END, 1)
		for _, str := range OAS.examples {
			res += tab(str, 2)
		}
	}

	if len(OAS.links) != 0 {
		res += tab(LINKS+END, 1)
		for _, str := range OAS.links {
			res += tab(str, 2)
		}
	}

	if len(OAS.callbacks) != 0 {
		res += tab(CALLBACKS+END, 1)
		for _, str := range OAS.callbacks {
			res += tab(str, 2)
		}
	}

	return res
}

func tab(str string, tabCount int) string {
	var tabDubl = ""
	for i := 0; i < tabCount; i++ {
		tabDubl += "  "
	}
	var re = regexp.MustCompile(`\r?\n`)
	str = re.ReplaceAllString(str, "\n"+tabDubl)
	var reLast = regexp.MustCompile(`\r?\n\s*$`)
	str = reLast.ReplaceAllString(str, "\n")
	str = tabDubl + str
	return str
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
