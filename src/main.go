package main

import(
	"fmt"
	"os"
)

func main(){

	arg := os.Args[1]
	var receivedKey string;
	var hash string = "e10adc3949ba59abbe56e057f20f883e"; // hash de 123456
	var authorized bool;


	fmt.Print("Enter the password: ")
	fmt.Print("\033[8m") // Hide input
	fmt.Scan(&receivedKey)
	fmt.Print("\033[28m") // Show input



	 authorized = authRequest(receivedKey, hash);
	if(authorized){
		fmt.Print("Authorized")
	}else{
		fmt.Print("Not authorized")
	}
	


	
	fmt.Print("\nArg = ", arg, "\n")
}

