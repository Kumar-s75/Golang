package main

import (
	"fmt";
	"sort"
)


func main(){
	fmt.Println("Welcome to video on slices");
    fmt.Println("Fruit list is: ",fruitList);

    var vegList=[5]string{"potato","beans","mushroom"};
	fmt.Println("Vegy list is: ",len(vegList));

	var fruitList =[]string{"Apple","Tomato","Peach"};
	fmt.Printf("Type of fruitlist is %T\n",fruitList);

	fruitList=append(fruitList,"Mango","Banana");
	fmt.Println(fruitList); 

	fruitList=append(fruitList[1:3]);
	fmt.Println(fruitList); 
	
	fruitList=append(fruitList[:3]);
	fmt.Println(fruitList); 

	highScores:=make([]int,4);

	highScores[0]=234; 
	highScores[1]=345;
	highScores[2]=465;
	highScores[3]=867;
    // highScores[4]=777;// this can catch error.

	highScores=append(highScores,555,666,321); 
	/// in this case all values get accommodated.
	fmt.Println(highScores);

	///things in slices but not in array

	sort.Ints(highScores);
	fmt.Println(highScores);
	fmt.Println(sort.IntsAreSorted(highScores)); 

	///how to remove a value from slices based on index

	var courses=[]string{"reactjs","javascript","swift","python","ruby"};
	fmt.Println(courses);
	var index int=2;
	courses=append(courses[:index],courses[index+1:]...);
	fmt.Println(courses); 

	 
}