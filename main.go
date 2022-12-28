package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Person struct {
	Id string
	Name string
}


func handler (writer http.ResponseWriter, request *http.Request) {
	person := Person{Id:"1",Name:"Foo"};
	parsedTemplate, _ := template.ParseFiles("templates/first-template.html");
	err := parsedTemplate.Execute(writer,person);

	if err != nil {
		log.Printf("Error", err);
	}
}

func main() {

	//Serving Static files
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/",fileServer))

	//Templating handler
	http.HandleFunc("/", handler); 

	//Server Lister
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT,nil);
	if err != nil {
		log.Fatal("Error while starting up server :", err);
	}
}