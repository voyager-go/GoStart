package main

import (
	"go-start/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
