//1-1-5 flag exp3
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

func (n *Name) String() string {
	return fmt.Sprint(*n)
}

func (n *Name) Set(value string) error {
	// 若n指向的字符串已被解析得到了参数, 又得到参数报错
	// 如使用: -name a -n b  b就会使得报错
	if len(*n) > 0 {
		return errors.New("name flag already set")
	}
	*n = Name("NameArg:" + value)
	return nil
}

func main() {
	var name Name
	flag.Var(&name, "name", "This is name")
	flag.Var(&name, "n", "this is n")
	flag.Parse()
	log.Printf("name: %s", name)
}
