package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the home page when the root path is requested
	http.ServeFile(w, r, "static/home.html")
}

func main() {
	// Handle root path first
	http.HandleFunc("/", rootHandler)

	// Then handle specific paths
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}