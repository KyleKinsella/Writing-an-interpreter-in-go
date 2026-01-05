package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl" // this is for my lexer and evaluator
	//"monkey/rppl" // this is for my parser
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	
	//repl.Start(os.Stdin, os.Stdout)
	//rppl.Start(os.Stdin, os.Stdout)
	repl.StartEvaluator(os.Stdin, os.Stdout)
}
