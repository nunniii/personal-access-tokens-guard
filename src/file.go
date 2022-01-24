package main

import(
	"os"
	"fmt"
	"io/ioutil"
)

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err);
		fmt.Fprintf(os.Stderr, "Reading error")
		os.Exit(1)
	}
	return string(data)
}

func writeFile(fileName string, content string){

	text := []byte(content)

	fileName = fmt.Sprintf("./data/%s", fileName);

    err := ioutil.WriteFile(fileName, text, 0644)
     
    if err != nil {
        fmt.Fprintf(os.Stderr, "Writing error")
		os.Exit(1)
    }
}
