package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("Welcome to files in golang");
	content :="This needs to go in a file-go";
   file, err:=os.Create("./mylcogofile.txt");

   if err !=nil{
     panic(err);
   }

   length,err:=io.WriteString(file,content);
   if err !=nil{
	panic(err);
  }
   fmt.Println("length is: ",length);
  defer  file.Close();
  readFile("./mylocofile.txt");
    //  length,err := io.WriteString(file,content);
}
//after you run command go main.go it a new .txt file is created and it goes around with the same name
 func readFile(filname string){
	databyte,err:=ioutil.ReadFile(filname);
	if err !=nil{
		panic(err);
	}
	fmt.Println("Text data inside the file is \n",databyte);
 }