// 1-1-3 flag exp1
package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "GoTour", "helpInfo : this is name")
	flag.StringVar(&name, "n", "GoTour", "helpInfo: this is n")
	flag.Parse()
	log.Printf("name: %s", name)
}
