package main

import "fmt"

func main() {
	fmt.println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Javascript"

	fmt.Println("List of all languages: ",languages);
	fmt.Println("Js shorts for: ",languages["JS"]);

	delete(languages,"RB");
	fmt.Println("List of all languages:",languages);


	////loops are intersting in golang

	for _,value  := range languages{
		fmt.Printf("For key v,value is  %v\n",value);
	}
     
	for key,value  := range languages{
		fmt.Printf("For key v,value is  %v\n",key,value);
	}
}