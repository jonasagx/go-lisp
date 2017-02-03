package lisp

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"syscall"
	"os/signal"
)

// Global variable
var historyHolder Timeline

func saveToHistory(historyParsed string) {
	fmt.Println("\nsaveToHistory")
	fmt.Print(historyParsed)
}

func 

func CaptureSIGTERM() {
    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    
    go func() {
        <-c
        saveToHistory(historyHolder.toString())
        os.Exit(0)
    }()
}

func Repl() {
	historyHolder = NewTimeline()
	CaptureSIGTERM()

	fmt.Printf("Welcome to the Lisp REPL\n")
	reader := bufio.NewReader(os.Stdin)
	expr := ""
	
	for {
		if expr == "" {
			fmt.Printf("\n> ")
		}
		line, _ := reader.ReadString('\n')
		historyHolder.NewLine(line)

		expr = fmt.Sprintf("%v%v", expr, line)
		openCount := strings.Count(expr, "(")
		closeCount := strings.Count(expr, ")")

		if openCount < closeCount {
			fmt.Printf("ERROR: Malformed expression: %v", line)
			expr = ""
		} else if openCount == closeCount {
			if strings.TrimSpace(expr) != "" {
				if response, err := EvalString(expr); err != nil {
					fmt.Printf("ERROR: %v\n", err)
				} else {
					if response == Nil {
						fmt.Println(";Unspecified return value")
					} else {
						fmt.Printf(";Value: %v\n", response.Inspect())
					}
				}
			}
			expr = ""
		}
	}
}
