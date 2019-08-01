package main

import (
	"ex03/q4/lib"
	myflag "ex03/q4/lib/flag"
	il "ex03/q4/lib/internal"
	"flag"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	myflag.Test()
	il.Hello(os.Stdout, "fffff")
}
