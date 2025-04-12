package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin);
	reader.ReadString('\n');
	input,  _ :=reader.ReadString('\n');

	fmt.Println("Thanks for rating",input);
	// numRating=strconv.ParseFloat(input,64);
	numRating=strconv.ParseFloat(strings.TrimSpace(input),64);
	if err !=nil{
		fmt.Println(err);
        panic(err);
	
	} else{
       fmt.Println("Added 1 to your rating: ",numRating + 1 );
	   
	}
}