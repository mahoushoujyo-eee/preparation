package main

import (
	"errors"
	"flag"
	"mycli/tool"
	"strings"
	_ "mycli/init"
)

func main() {
	args := flag.Args()
	if len(args) == 0 {
		panic(errors.New("no args"))
	}

	if strings.Compare(args[0], "cat") == 0 {
		err := tool.Cat(args)
		if err != nil {
			panic(err)
		}
	} else if strings.Compare(args[0], "grep") == 0{
		err := tool.Grep(args)
		if err != nil {
			panic(err)
		}
	}
}
