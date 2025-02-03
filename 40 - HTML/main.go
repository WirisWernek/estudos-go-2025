package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", nil)
	})

	http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{
			Nome:  "Wiris",
			Email: "wiriswernek@gmail.com",
		}
		templates.ExecuteTemplate(w, "usuario.html", u)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
