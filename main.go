package main

import (
	"fmt"
	"net/http"
)

type User struct {
	fname     string
	lname     string
	age       int8
	happyness float32
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf(
		"name: %v\nlastname: %v\nage: %v\nhap: %v\n",
		u.fname, u.lname, u.age, u.happyness,
	)
}

func (u *User) setHappyness(newHappyness float32) {
	u.happyness = newHappyness
}

func homePage(w http.ResponseWriter, r *http.Request) {

	user1 := User{"Dmitry", "Bakhmetyev", 26, .9}

	user1.setHappyness(.4)

	fmt.Fprintf(w, user1.getAllInfo())

}

func initHttpSrv() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	initHttpSrv()
}
