package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sivchari/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Hello, %s! This is the Monkey programming language!", user.Username))
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
