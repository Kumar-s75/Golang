package main

import "fmt"

func main(){
	fmt.Println("Structs in golang");
	//no inheritance in golang ;No super or parent

	hitesh:=User{"Hitesh","hitesh@go.dev",true,16};
	fmt.Println(hitesh);
	fmt.Printf("Hitesh details are: %+v\n",hitesh);
	//+v gives you the structure and the naming convection
	fmt.Printf("Name is %v and email is %v\n.",hitesh.Name,hitesh.Email);
    hitesh.GetStatus();
	hitesh.NewMail();
	fmt.Printf("Name is %v and email is %v\n.",hitesh.Name,hitesh.Email);
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
	oneAge int
}

func(u User) GetStatus(){
	fmt.Println("Is user active :",u.Status);

}

func (u User) NewMail(){
   u.Email="test@go.dev";
   fmt.Println("Email of this user is :  ",u.Email);
}