package main


import "fmt"

func main(){
    fmt.Println("Welcome to a class on pointers");

	// var one int =2;
     var ptr *int 
	  fmt.Println("Value of pointer is ",ptr);

	  myNumber:=23;
	   var ptr= &myNumber;

	   fmt.Println("value of actual pointer is ",ptr);
	   fmt.Println("value of actual pointer is ",*ptr);

	   *ptr=*ptr*2;
	   fmt.Println("New Value is: ",myNumber);
}