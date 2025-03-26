package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

var Users []User
var UsersByID map[int]User

func InitUsers(usersData []User) {
	Users = usersData
	UsersByID = make(map[int]User)

	for _, user := range Users {
		UsersByID[user.ID] = user
	}
}
