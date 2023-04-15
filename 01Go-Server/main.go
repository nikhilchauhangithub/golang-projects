package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request)  {
	err :=r.ParseForm()
	if err!=nil {
		fmt.Fprintf(w,"Parseform err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name:= r.FormValue("name")
	email:= r.FormValue("email")
	fmt.Fprintf(w, "Name= %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", email)

}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method !="GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}




func main() {
   var fileServer = http.FileServer(http.Dir("./static"))
   http.Handle("/", fileServer)
   http.HandleFunc("/form", formHandler)
   http.HandleFunc("/hello",helloHandler)

   fmt.Printf("starting server at port 8080\n")

   err:=http.ListenAndServe(":8080",nil);
   if err !=nil{
	log.Fatal(err)
   }


}