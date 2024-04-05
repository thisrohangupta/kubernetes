// render_template.go
package main

import (
    "os"
    "text/template"
    "io/ioutil"
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
    values, err := ioutil.ReadFile(valuesFile)
    if err != nil {
        panic(err)
    }

    // Execute template with values
    if err := tmpl.Execute(os.Stdout, string(values)); err != nil {
        panic(err)
    }
}
