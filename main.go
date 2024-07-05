package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/zopitty/interpreter/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("hello %s! this is the monkey language\n", user.Username)
    fmt.Printf("feel free to type some bananas\n")
    repl.Start(os.Stdin, os.Stdout)
}
