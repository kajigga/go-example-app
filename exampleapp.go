package main

import (
	"flag"
	"fmt"
)

const VERSION = "0.4"

type Args struct {
	GetVersion bool
	Message    string
}

func main() {

	a := Args{}
	flag.BoolVar(&a.GetVersion, "version", false, "get the version")
	flag.StringVar(&a.Message, "message", "", "message to share")

	flag.Parse()

	if a.GetVersion {
		version()
	} else {
		fmt.Println(a.Message)
	}

}

func version() {
	fmt.Println(VERSION)
}
