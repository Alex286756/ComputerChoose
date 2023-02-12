package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bobc", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	templ, err := template.ParseFiles("../templates/html/home_page.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	templ.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
