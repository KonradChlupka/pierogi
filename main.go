package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/KonradChlupka/pierogi/repl"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Failed to get the current user:", err)
		os.Exit(1)
	}
	fmt.Printf("Hello %s! This is the Pierogi programming language!\n", currentUser.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
