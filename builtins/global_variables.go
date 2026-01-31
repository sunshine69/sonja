package builtins

import "github.com/sunshine69/gonja/v2/exec"

var GlobalVariables = exec.NewContext(map[string]interface{}{
	"gonja": map[string]interface{}{
		"version": "v0.0.0+trunk",
	},
})
