// render_template.go
package main

import (
	"os"
	"text/template"
)

func main() {
	// Load deployment.yaml
	templateFile := os.Args[1]
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	// Load values.yaml
	valuesFile := os.Args[2]
	values, err := os.ReadFile(valuesFile)
	if err != nil {
		panic(err)
	}

	// Parse values.yaml
	var data map[string]interface{}
	if err := yaml.Unmarshal(values, &data); err != nil {
		panic(err)
	}

	// Execute template with values
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}
