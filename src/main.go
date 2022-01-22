package main

import(
	"fmt"
	"os"
)

func main(){


	authRequest();

	arg := os.Args[1]
	fmt.Print("\nArg = ", arg, "\n")
}

