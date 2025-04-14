package main

import "fmt"

func main() {
	greeter();
	fmt.Println("Welcome to functions in golang");

	fmt.Println("Result is:");
	greeterTwo();

	proRes:=proAdder(2,5,8,7);
	fmt.Println("Pro result is : ",proRes);

	proRes,myMessage:=proAdder(2,5,8,7,3);
	fmt.Println("Pro Result is: ",proRes);
	fmt.Println("Pro Message is: ",myMessage);
}
func greeterTwo(){
	fmt.Println("Another Method");
}
func greeter(){
	fmt.Println("Hello from golang");
}
func adder(valOne int,valTwo int) int {
	return valOne + valTwo; 
}

func proAdder(values ...int) int{
   total := 0

   for _,val := range  values{
	  total += val;
   }

   return total, "Hi Pro result function";
}