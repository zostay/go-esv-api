package main

//go:generate go run main.go

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"
)

type APISpec struct {
	Root      string     `yaml:"root"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Name     string          `yaml:"name"`
	Path     string          `yaml:"path"`
	Required []ParameterSpec `yaml:"required"`
	Optional []ParameterSpec `yaml:"optional"`
	Result   ParameterSpec   `yaml:"result"`
}

type ParameterSpec struct {
	Name   string          `yaml:"name"`
	Type   string          `yaml:"type"`
	Struct []ParameterSpec `yaml:"struct"`
}

type StructSpec struct {
	Struct   []ParameterSpec
	Indent   string
	IndentBy int
}

func main() {
	specBytes, err := ioutil.ReadFile("esv-api.yaml")
	if err != nil {
		panic(fmt.Errorf("unable to read esv-api.yaml: %w", err))
	}

	var s APISpec
	err = yaml.Unmarshal(specBytes, &s)
	if err != nil {
		panic(fmt.Errorf("error reading API spec: %w", err))
	}

	fh, err := os.Create("../../pkg/esv/api.go")
	if err != nil {
		panic(fmt.Errorf("error writing api.go: %w", err))
	}
	defer fh.Close()

	tmplBytes, err := ioutil.ReadFile("api.go.tmpl")
	if err != nil {
		panic(fmt.Errorf("unable to read api.go.tmpl: %w", err))
	}

	tmpl := template.New("api")
	tmpl.Funcs(map[string]interface{}{
		"ToCamel": strcase.ToCamel,
		"FunctionArgs": func(ps []ParameterSpec) string {
			args := make([]string, len(ps))
			for i, p := range ps {
				args[i] = p.Name + " " + p.Type
			}
			return strings.Join(args, ", ")
		},
		"PrepareStruct": func(ps []ParameterSpec, indent int) StructSpec {
			return StructSpec{ps, strings.Repeat(" ", indent), indent}
		},
		"Add": func(a, b int) int { return a + b },
		"Set": func(wo map[string]bool, name string) string {
			wo[name] = true
			return ""
		},
	})

	_, err = tmpl.Parse(string(tmplBytes))
	if err != nil {
		panic(fmt.Errorf("error parsing api.go.tmpl: %w", err))
	}

	vars := make(map[string]interface{})
	vars["codeType"] = "generated"
	vars["codeGenerator"] = "github.com/zostay/go-esv-api-sdk/tools/gen"
	vars["codeEditable"] = "NOT"
	vars["spec"] = s
	vars["wo"] = make(map[string]bool)

	err = tmpl.Execute(fh, vars)
	if err != nil {
		panic(fmt.Errorf("error executing template: %w", err))
	}

	err = exec.Command("go", "fmt", "../../pkg/esv/api.go").Run()
	if err != nil {
		panic(fmt.Errorf("error executing go fmt: %w", err))
	}
}
