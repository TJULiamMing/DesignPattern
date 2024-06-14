package main

import (
	"designPattern/BehavioralPattern/iterator"
	"fmt"
)

func main() {
	user1 := &iterator.User{Name: "user1"}
	user2 := &iterator.User{Name: "user2"}
	user3 := &iterator.User{Name: "user3"}

	ui := &iterator.UserCollection{Users: []*iterator.User{user1, user2, user3}}

	i := ui.NewIterator()

	for i.HasNext() {
		user := i.Next()
		fmt.Printf("User is %+v\n", user)
	}

}
