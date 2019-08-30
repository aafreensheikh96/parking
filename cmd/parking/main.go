package main

import (
	"parking"

	"github.com/alexflint/go-arg"
)

type arguements struct {
	Filename string `help:"Souce file for commands"`
}

func main() {
	var args arguements
	arg.MustParse(&args)
	if args.Filename != "" {
		parking.ExecuteFile(args.Filename)
	} else {
		parking.ExecuteCLI()
	}
}
