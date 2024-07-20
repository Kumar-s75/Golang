// package main

// import "fmt"

//	func main() {
//		fmt.Println("Hello from NITJ")
//	}
package variables

import "fmt"

// import "fmt"
 
const LoginToken string="ghabbrkh";//public

// jwtToken:=3000000 not allowed find what is allowed yourself

func main(){
	// fmt.Println("Variables");
	var username string= "hitesh";
	fmt.Println(username);
	fmt.Printf("variable is of type:%T \n",username);

	var isLoggedIn bool= false;
	fmt.Println(isLoggedIn);
	fmt.Printf("variable is of type:%T \n",isLoggedIn);

	var smallVal uint8= 255;
	fmt.Println(smallVal);
	fmt.Printf("variable is of type:%T \n",smallVal);

	var smallFloat float32= 255.45678;
	fmt.Println(smallFloat);
	fmt.Printf("variable is of type:%T \n",smallFloat);

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable);
	fmt.Printf("Variable is of type:%T \n",anotherVariable);
	//implicit type
	var website="learncodeonline.in";
	fmt.Println(website);
	//no var style
	numberofUser:= 300000;
	fmt.Println(numberofUser);

	
	fmt.Println(LoginToken);
	fmt.Printf("Variable is of type:%T \n",LoginToken);

}