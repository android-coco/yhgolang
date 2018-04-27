// Example code for Chapter 4.2 from "Build Web Application with Golang"
// Purpose: Shows how to perform server-side validation of user input from a form.
// Also shows to use multiple template files with predefined template names.
// Run `go run main.go` and then access http://localhost:9090
package main

import (
	"html/template"
	"log"
	"net/http"
	"awesomeProject/study/gowebbookcode/apps/ch.4.2/validator"
)

const (
	PORT     = "10000"
	HOSTURL = "http://localhost:" + PORT
)

var t *template.Template

type Links struct {
	BadLinks [][2]string
}

// invalid links to display for testing.
var links Links

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, HOSTURL+"/profile", http.StatusTemporaryRedirect)
}
func profileHandler(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "profile", links)
}
func checkProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := validator.ProfilePage{&r.Form}
	t.ExecuteTemplate(w, "submission", p.GetErrors())
}

// This function is called before main()
func init() {
	// Note: we can reference the loaded templates by their defined name inside the template files.
	t = template.Must(template.ParseFiles("C:/workHome/gowork/src/awesomeProject/study/gowebbookcode/apps/ch.4.2/profile.gtpl", "C:/workHome/gowork/src/awesomeProject/study/gowebbookcode/apps/ch.4.2/submission.gtpl"))

	list := make([][2]string, 2)
	list[0] = [2]string{HOSTURL + "/checkprofile", "No data"}
	list[1] = [2]string{HOSTURL + "/checkprofile?age=1&gender=guy&shirtsize=big", "Invalid options"}
	links = Links{list}
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/checkprofile", checkProfile)

	err := http.ListenAndServe(":"+PORT, nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
