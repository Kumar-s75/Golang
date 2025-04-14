package main

import "fmt"

const url="https://lco.dev"

func main() {
	fmt.Println("Lco web requests");

	 response,err:=http.Get(url);

	 if err!=null{
        panic(err)
}

fmt.Printf("Response is of type:%T\n",response);

 defer response.Body.Close();
///callers responsibillity to close the connection.
   databytes,err:=ioutil.ReadAll(response.body);
   if(err!=nil){
	panic(err);
   }
   content:=string(databytes);
}