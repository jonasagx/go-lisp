package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jonasagx/go-lisp/lisp"
	"io/ioutil"
	"os"
)

func main() {
	var file = flag.String("file", "", "File to interpret")
	var repl = flag.Bool("i", false, "Interactive mode")
	flag.Parse()

	if *repl {
		lisp.Repl()
	} else if *file != "" {
		if output, err := ioutil.ReadFile(*file); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			if _, err := lisp.EvalString(string(output)); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	} else {
		if output, err := ioutil.ReadAll(bufio.NewReader(os.Stdin)); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			lisp.EvalString(string(output))
		}
	}
}
