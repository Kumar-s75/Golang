Package main

import "fmt"

type course struct{
  Name string `json:"coursename" `
   Price int
   Platform string `json:"website" `
   Password string `json:"-" `
   Tags    []string   `json:"tags,omitempty"`
   //omitempty -if the value is null just dont den /throw that field in json.
}

func main(){
	fmt.Println("Welcome to JSON video");
	EncodeJson();
	DecodeJson();
}

func EncodeJson(){
	lcoCourses :=[]courses{
		{"ReactJs Bootcamp", 299,"LearnCodeOnline.in","abc123",[]string{"web-dev","js"} },
		{"MERN Bootcamp", 199,"LearnCodeOnline.in","bcd123",[]string{"fullstack","js"} },
		{"Angular Bootcamp", 399,"LearnCodeOnline.in","hit123",nil },
	}

	//package this dat as Json

	// finalJson,err := json.Marshal(lcoCourses)
	finalJson,err := json.Marshal(lcoCourses,"lco","\t");
	if err !=nil{
		panic(err);
	}
	fmt.Printf("%s\n",finalJson);
}

func DecodeJson(){
	 jsonDataFromWeb:=[]byte(`
	 {

"coursename":"ReactJs Bootcamp",
"Price":299,
"website":"LearnCodeOnline.in",
"tags":[
"web-dev",
"js"
],


	 },
	 
	 `);

	 var lcoCourse course

	 checkValid:=json.valid(jsonDataFromWeb);

	 if checkValid{
		fmt.Println("JSON was valid");
		json.Unmarshal(jsonDataFromWeb,&lcoCourse);
		fmt.Printf("%#v\n",lcoCourse);
	 } else{
		fmt.Println("JSON was not valid");
	 }

	 ///some cases where you just want to add data to key value pair

	 var myOnlineData map[string] interface{}
	 json.Unmarshal(jsonDataFromWeb,&myOnlineData);
	 fmt.Printf("%#v\n",myOnlineData);

	 for k,v:= range myonlineData{
		fmt.Printf("Key is %v and value is %v and Type is:%T\n ",k,v,)
	 }
}