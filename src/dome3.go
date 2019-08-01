package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

//ss := "12345"
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	//flag.StringVar(&name, "name", "everybody", "this is a name")

	// flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	// flag.CommandLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	// name = flag.String("name", "test", "this is a name")

	name = *cmdLine.String("name", "test", "this is a name")
	cmdLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "test")
		cmdLine.PrintDefaults()
	}
}

func main() {
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	cmdLine.Parse(os.Args[1:])
	//flag.Parse()

	fmt.Println("Hello, " + name)
	//	fmt.Println(ss)

}
