package main

import(
	"fmt"
)


func authRequest() bool {

		var receivedKey string;
		var authorized bool;

		fmt.Print("Enter the password: ")
		fmt.Print("\033[8m") // Hide input
		fmt.Scan(&receivedKey)
		fmt.Print("\033[28m") // Show input
		// fmt.Print(receivedKey)

	
	var hash string = "912ec803b2ce49e4a541068d495ab570"; 

	if (createHash(receivedKey) == hash){
		fmt.Print("Authorized")
		authorized = true;
	}else{
		fmt.Print("Not authorized")
		authorized = false;
	}
	return authorized;
}

