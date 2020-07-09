package main

import (
	"net/http"
	"html/template"
	"time"
	"fmt"
)

type Welcome struct{
	Time string
}

type About struct {

}

var welcome_template = template.Must(template.ParseFiles("templates/welcome.html"))
var about_template = template.Must(template.ParseFiles("templates/about.html"))
func main() {
	
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	http.HandleFunc("/home", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/health", health)

	http.ListenAndServe(":8000", nil)
}


func home(w http.ResponseWriter, r *http.Request){

	welcome := Welcome{time.Now().Format(time.Stamp)}

	{
		if err := welcome_template.ExecuteTemplate(w, "welcome.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func health(w http.ResponseWriter, r *http.Request){

	message := "All Good here!"

	fmt.Println(message)

}


func about(w http.ResponseWriter, r *http.Request){

	about := About{}

	if err := about_template.ExecuteTemplate(w, "about.html", about); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}