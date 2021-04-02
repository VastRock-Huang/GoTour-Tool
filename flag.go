//1-1-1 flag exp1
//package main
//
//import (
//	"flag"
//	"log"
//)
//
//func main() {
//	var name string
//	flag.StringVar(&name, "name", "GoTour", "helpInfo : this is name")
//	flag.StringVar(&name, "n", "GoTour", "helpInfo: this is n")
//	flag.Parse()
//	log.Printf("name: %s", name)
//}

////1-1-2 flag exp2
//package main
//
//import (
//	"flag"
//	"log"
//)
//
//func main() {
//	var name string
//	flag.Parse()
//	args := flag.Args()
//	if len(args) <= 0 {
//		return
//	}
//	switch args[0] {
//	case "go":
//		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
//		goCmd.StringVar(&name, "name", "Golang", "help")
//		_ = goCmd.Parse(args[1:])
//	case "php":
//		phpCmd := flag.NewFlagSet("php", flag.ContinueOnError)
//		phpCmd.StringVar(&name, "n", "php", "help")
//		_ = phpCmd.Parse(args[1:])
//	}
//	log.Printf("name: %s", name)
//
//}

//1-1-3 flag exp3
//package main
//
//import (
//	"errors"
//	"flag"
//	"fmt"
//	"log"
//)
//
//type Name string
//
//func (n *Name) String() string {
//	return fmt.Sprint(*n)
//}
//
//func (n *Name) Set(value string) error {
//	// 若n指向的字符串已被解析得到了参数, 又得到参数报错
//	// 如使用: -name a -n b  b就会使得报错
//	if len(*n) > 0 {
//		return errors.New("name flag already set")
//	}
//	*n = Name("NameArg:" + value)
//	return nil
//}
//
//func main() {
//	var name Name
//	flag.Var(&name, "name", "This is name")
//	flag.Var(&name, "n", "this is n")
//	flag.Parse()
//	log.Printf("name: %s", name)
//}

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
