package main

import (
	"github.com/vastrock-huang/gotour-tool/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Excute err: %v", err)
	}
}
