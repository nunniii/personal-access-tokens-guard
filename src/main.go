package main

import(
	"fmt"
	"os"
)


func enterPassword(receivedKey string, hash string) bool {

	var authorized bool = false;


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

	return authorized;
}


func getToken(receivedKey string){
	fmt.Print("getToken")
}

func createToken(){
	fmt.Print("createToken")
}

func changePassword(){
	fmt.Print("changePassword")
}


func main(){


	arg := os.Args[1]

	var receivedKey string;
	var hash string = "e10adc3949ba59abbe56e057f20f883e"; // hash de 123456




	// Start
	if(arg == "--new-key"){
		createToken()

	}else if(arg == "--change-password"){
		if(enterPassword(receivedKey, hash)){
			changePassword()
		}else{
			fmt.Fprintf(os.Stderr, "Not authorized")
			os.Exit(1)
		}

	}else{
		if(enterPassword(receivedKey, hash)){
			getToken(arg)
		}else{
			fmt.Fprintf(os.Stderr, "Not authorized")
			os.Exit(1)
		}

	}

	

	
	fmt.Print("\nArg = ", arg, "\n")
}

