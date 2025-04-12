Package main

func main(){
   fmt.Println("Welcome to array in golang");

   var fruitList []string

   fruitList[0]="Apple"
   fruitList[1]="Tomato"
   fruitList[3]="Peach"


   fmt.Println("Fruit list is : ",fruitList);
   fmt.Println("Fruit list is : ",len(fruitList));
   
   var vegList=[5]string{"Potato","beans","mushroom"};
fmt.Println("Vegy list is: ",len(vegList) );

}