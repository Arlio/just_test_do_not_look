package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Fname     string
	Lname     string
	Age       int8
	Happyness float32
	Langs     []string
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf(
		"name: %v\nlastname: %v\nage: %v\nhap: %v\n",
		u.Fname, u.Lname, u.Age, u.Happyness,
	)
}

func (u *User) setHappyness(newHappyness float32) {
	u.Happyness = newHappyness
}

func homePage(w http.ResponseWriter, r *http.Request) {

	user1 := User{"Dmitry", "Bakhmetyev", 26, .9, []string{
		"JS", "PHP", "Python", "GO",
	}}

	user1.setHappyness(.4)

	//fmt.Fprintf(w, user1.getAllInfo())

	tmpl, _ := template.ParseFiles("templates/home_page.html")

	tmpl.Execute(w, user1)

}

func initHttpSrv() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	initHttpSrv()
}
