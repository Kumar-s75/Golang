package main

import (
	"bufio"
	"fmt"
)

func main(){
	welcome := "Welcome to user input"
	fmt.Println(welcome);

	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:");

	////comma ok syntax///err ok syntax

	input, err:=reader.ReadString('\n')/// this can be trated as try and catch , input
	//is try here and err is catch , if any error comes in this is the error catch up here!
	fmt.Println("Thanks for Rating,",input );
	fmt.Printf("Type of this rating is %T,",input);
}