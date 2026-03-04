package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type LogItem struct {
	action   string
	timestep time.Time
}

type User struct {
	id    int
	email string
	logs  []LogItem
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID%d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i, item.action, item.timestep)
	}

	return out
}

func main() {
	rand.Seed(time.Now().Unix())

	users := generateUsers(1000)

	for _, user := range users {
		saveUserInfo(user)
	}
	// u := User{
	// 	id: 1,
	// 	email: "olof@gmail.com",
	// 	logs: []LogItem {
	// 		{actions[0], time.Now()},
	// 		{actions[3], time.Now()},
	// 		{actions[2], time.Now()},
	// 		{actions[1], time.Now()},
	// 		{actions[0], time.Now()},
	// 		{actions[3], time.Now()},
	// 	},
	// }
	// go fmt.Print("concurent hello world")
	// go fmt.Print("concurent hello world")
	// go fmt.Print("concurent hello world")

	// time.Sleep(time.Second)

	// fmt.Println("Unconcurent Hello world")
}

func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@gmail.com", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}
	return users
}

func saveUserInfo(user User) error {
	fmt.Printf("WRITING FILE FOR USER ID: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	return err

}

func generateLogs(count int) []LogItem {
	logs := make([]LogItem, count)

	for i := 0; i < count; i++ {
		logs[i] = LogItem{
			timestep: time.Now(),
			action:   actions[rand.Intn(len(actions)-1)],
		}
	}
	return logs
}
