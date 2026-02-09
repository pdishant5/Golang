package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	u1 := User{ID: 1, Name: "Dishant", Surname: "Patel"}
	jsonData, _ := json.MarshalIndent(u1, "", "    ")

	fmt.Println(jsonData)
	os.WriteFile("user.json", jsonData, 0644)

	var u2 User
	_ = json.Unmarshal(jsonData, &u2)
	fmt.Println(u2)
}
