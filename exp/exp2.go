//1-1-3 flag exp2
package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Golang", "help")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ContinueOnError)
		phpCmd.StringVar(&name, "n", "php", "help")
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)

}
