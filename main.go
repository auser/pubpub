package main

import "github.com/auser/pubpub/pubpub"

var (
	// VERSION set during build
	VERSION = "0.0.1"
)

func main() {
	pubpub.Execute(VERSION)
}
