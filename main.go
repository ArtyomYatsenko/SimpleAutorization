package main

import (
	"html/template"
	"log"
	"net/http"
)

type DataAutorisation struct {
	Login        string
	Password     string
	Succes       bool
	StorageAcces string
}

var tmpl = template.Must(template.ParseFiles("autorisation.html"))

func handler(res http.ResponseWriter, rec *http.Request) {
	data := DataAutorisation{
		Login:        rec.FormValue("login"),
		Password:     rec.FormValue("password"),
		StorageAcces: "Вы авторизованы.",
	}
	if data.Password == "123456" && data.Login == "test" {
		data.Succes = true
	}
	err := tmpl.Execute(res, data)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
