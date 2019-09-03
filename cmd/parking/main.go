package main

import (
	"fmt"
	"parking"

	"github.com/alexflint/go-arg"
)

type arguements struct {
	Filename string `arg:"positional" help:"Souce file for commands"`
}

func main() {
	fmt.Println("running tool")
	var args arguements
	arg.MustParse(&args)
	if args.Filename != "" {
		parking.ExecuteFile(args.Filename)
	} else {
		parking.InteractiveSession()
	}
}
