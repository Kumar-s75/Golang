package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Hello mod in golang");
	greeter();
	r:=mux.NewRouter();
	r.HandleFunc("/",serveHome).Methods("GET");

	// http.ListenAndServe(":4000",r);
	log.Fatal(http.ListenAndServe(":4000",r));
}

func greeter(){
	fmt.Println("Hey There mod users");
}

func serveHome(w http.ResponseWriter,r *http.Request){
    w.Write([]byte("<h1> Welcome to golang series on Youtube</h1>"));

}