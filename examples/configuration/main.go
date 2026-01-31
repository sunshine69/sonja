package main

import (
	"os"

	"github.com/sunshine69/sonja/v2"
	"github.com/sunshine69/sonja/v2/exec"
)

func main() {
	gonja.DefaultConfig.VariableStartString = "{$"
	gonja.DefaultConfig.VariableEndString = "$}"
	gonja.DefaultConfig.StrictUndefined = true

	template, err := gonja.FromString(`{$ var ~ " is fun!" | capitalize $}`)
	if err != nil {
		panic(err)
	}

	ctx := exec.NewContext(map[string]interface{}{
		"var": "gonja",
	})
	if err = template.Execute(os.Stdout, ctx); err != nil { // Prints Gonja is fun! to stdout
		panic(err)
	}
}
