package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main(){
	fmt.Println("Welcome to web verb video-LCO");
	PerformGetRequest();
	PerformPostJsonRequest();
	PerformPostFormRequest();
    PerformPostFormRequest();
}

func PerformGetRequest(){
	const myurl="http://localhost:8000/get";

	response,err:=http.Get(myurl);
	if err !=nil{
		panic(err);
	}
	defer response.Body.Close();

	fmt.Println("Status code : ",response.StatusCode);
	fmt.Println("Content Length is: ",response.ContentLength);


	var responseString strings.Builder;
	content,_:=ioutil.ReadAll(response.Body);
     byteCount, _ := responseString.Write(content);
	 fmt.Println("Byte count is: ",byteCount);
	 fmt.Println(responseString.String());
	// fmt.Println(string(content));
	//content is in the byte format
	//bytecount data is stored in responsestring

}

func PerformPostJsonRequest(){
	const myurl="http://localhost:8000/post";

	//fake json payload

	requestBody := strings.NewReader(`
	{
	  "coursename":"	Let's go with golang",
	  "price":0,
	  "platform:":"learnCodeOnline.in"
	  	}
	`)

	response,_ :=http.Post(myurl,"application/json",requestBody);

	if err!=nil{
		panic(err)
	}
	defer response.Body.close();
	  content,_:= ioutil.ReadAll(response.Body);
	  fmt.Println(string())
}

func PerformPostFormRequest(){
	const myurl="http://localhost:8000/postform";

	///formdata
	data := url.Values{}
	data.Add("firstname","hitesh");
	data.Add("lastname","choudhary");
	data.Add("email","hitesh@go.dev");

	response,err:=http.PostForm(myurl,data);

   if err != nil{
	   panic(err);
   }

   content, _ :=ioutil.ReadAll(response.body);
  fmt.Println(string(content));


}