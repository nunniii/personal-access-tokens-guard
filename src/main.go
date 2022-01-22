package main

import(
	"fmt"
	"os"
)

func main(){

	arg := os.Args[1]
	authRequest();


	
	fmt.Print("\nArg = ", arg, "\n")
}

