package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Neel-shetty/monkey/repl"
	"github.com/Neel-shetty/monkey/runner"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		err := runner.RunFile(os.Args[1], os.Stdout)
		if err != nil {
			fmt.Printf("Error executing file: %s\n", err)
			os.Exit(1)
		}
		return
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
