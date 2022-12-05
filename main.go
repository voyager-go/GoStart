package main

import (
	"go-start/cmd"
	"os"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		panic(err)
	}
}
