package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"log"
	"os"
	"text/template"

	"golang.org/x/tools/imports"
)

type Name struct {
	First string
	Last  string
}

type Person struct {
	Name Name
}

type Tmpl struct {
	Value []byte
}

func main() {
	t, err := template.New("gen.tmpl").ParseFiles("gen.tmpl")
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	var buf bytes.Buffer

	v := Person{
		Name: Name{
			First: "John",
			Last:  "Doe",
		},
	}

	vdata, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("failed to marshal: %v", err)
	}

	if err = t.Execute(&buf, &Tmpl{Value: vdata}); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}

	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("failed to format source: %v", err)
	}

	data, err = imports.Process("gen.go", data, nil)
	if err != nil {
		log.Fatalf("failed to process imports: %v", err)
	}

	err = os.WriteFile("gen.go", data, 0644)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}
