package main

import(
	"fmt"
	"os"
	"github.com/go-vgo/robotgo/clipboard"
)


func enterPassword(receivedKey string, hash string) bool {

	var authorized bool = false;

	fmt.Print("Enter the password: ")
	fmt.Print("\033[8m") // Hide input
	fmt.Scan(&receivedKey)
	fmt.Print("\033[28m") // Show input

	authorized = authRequest(receivedKey, hash);

	return authorized;
}


func getToken(receivedKey string){

	path := fmt.Sprintf("./data/tokens/%s", receivedKey)

	token := readFile(path);


	clipboard.WriteAll(token)
	fmt.Print("Token was copied\n")

}

func createToken(){

	var name, token string;

	fmt.Print("Enter a name for the token: ")
	fmt.Scan(&name)
	name = fmt.Sprintf("tokens/%s", name)
	fmt.Print("Enter the token: ")
	fmt.Print("\033[8m") // Hide input
	fmt.Scan(&token)
	fmt.Print("\033[28m") // Show input

	writeFile(name, token);
}

func changePassword(){

	var password string;

	fmt.Print("Enter the new password: ")
	fmt.Print("\033[8m") // Hide input
	fmt.Scan(&password)
	fmt.Print("\033[28m") // Show input
	writeFile("psswd", createHash(password));

	fmt.Print("Changed password\n")


}


func main(){


	arg := os.Args[1]

	var receivedKey string;
	var hash string = readFile("./data/psswd"); // hash de 123456




	// Start
	if(arg == "--new-key"){
		createToken()

		// a := readFile("./data/tokens");
		// fmt.Print(a)

	}else if(arg == "--change-password"){
		if(enterPassword(receivedKey, hash)){
			changePassword()
		}else{
			fmt.Fprintf(os.Stderr, "\tNot authorized\n\t")
			os.Exit(1)
		}

	}else{
		if(enterPassword(receivedKey, hash)){
			getToken(arg)
		}else{
			fmt.Fprintf(os.Stderr, "\tNot authorized\n\t")
			os.Exit(1)
		}

	}




}

