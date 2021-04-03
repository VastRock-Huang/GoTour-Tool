// 1-4-2 tmplate exp
package main

import (
	"os"
	"strings"
	"text/template"
)

const templateText = `
Output 1: {{title .Name1}}
Output 2: {{title .Name2}}
Output 3: {{.Name3 | title}}
`

func main() {
	funcMap := template.FuncMap{"title": strings.Title}
	tpl := template.New("go tour")
	tpl, _ = tpl.Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "cpp",
		"Name3": "java",
	}
	_ = tpl.Execute(os.Stdout, data)
}
