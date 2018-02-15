package main

import "fmt"

type Member struct {
	username string
	email string
	age int
	enabled bool
}

func (u Member) GetUsername() string { // without pointer
	return u.username
}

func (u* Member) ChangeStatus() { // with pointer
	u.enabled = !u.enabled
}



func main() {
	var member1 Member
	member1.username = "mdiaz"
	member1.email = "mdiaz@is4tech.com"
	member1.age = 28
	member1.enabled = true

	member2 := Member{username: "clopez", email: "clopez@is4tech.com", age: 33}

	fmt.Println(member1)
	fmt.Println(member2)


	fmt.Println("\nChanging status for user:", member2.GetUsername())
	member2.ChangeStatus()
	fmt.Println(member2)

}