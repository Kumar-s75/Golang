package main

import "fmt"

func main() {
	fmt.Println("Switch and case in golang");

	rand.seed(time.Now().UnixNano())
	diceNumber:=rand.Intn(6) + 1
	fmt.Println("Value of dice is ",diceNumber);

	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 and you can open");
	case 2:
		fmt.Println("You can move 2 spot");	
    case 3:
		fmt.Println("You can move 3 spot");	
		fallthrough;
	case 4:
		fmt.Println("You can move 4 spot");
		fallthrough;		
	case 5:
		fmt.Println("You can move 5 spot");	
	case 6:
		fmt.Println("You can move 6 spot and roll dice again");	
		default:
			fmt.Println("What was That!")			
	}

}