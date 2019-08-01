package lib

import (
	in "ex03/q4/lib/internal"
	"os"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
