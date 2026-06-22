package main

import (
	"flag"
	"log"
	_ "mycli/init"
	"mycli/tool"
)

func main() {
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalln("no args")
	}

	cmd := args[0]

	switch cmd {
	case "cat":
		err := tool.Cat(args)
		if err != nil {
			log.Fatalln(err)
		}
	case "grep":
		err := tool.Grep(args)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalf("unknown command: %s\n", cmd)
	}
}
