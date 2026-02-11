package main

import (
	"fmt"
)

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func (u *User) printUserInfo(name string) {
	u.name = name
	fmt.Println(u.name, u.age, u.sex, u.height, u.weight)
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    age,
		weight: weight,
		height: height,
	}
}

type DumpDatabase struct {
	m map[string]string
}

func NewDumpDatabase() *DumpDatabase {
	return &DumpDatabase{
		m: make(map[string]string, 0),
	}
}

func main() {
	db := DumpDatabase{
		m: make(map[string]string),
	}

	db.m["login"] = "root"
	db.m["pass"] = "1234"

	fmt.Println(db.m)

	user1 := NewUser("Vasya", "Male", 23, 75, 185)
	user2 := User{"Petya", 25, "Male", 85, 195}

	user1.printUserInfo("Kostya")
	user2.printUserInfo("Sergey")

	fmt.Println(user1.name)
	fmt.Println(user2.name)
	// fmt.Println(user1.name, user1.sex)

	// fmt.Printf("%+v\n", user1)
	// fmt.Printf("%+v\n", user2)
}
