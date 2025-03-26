package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"surfe_test/models"
)

func main() {
	users, actions, err := loadData()
	if err != nil {
		log.Fatalf("couldn't load data: %s", err.Error())
	}

	models.InitUsers(users)
	models.InitActions(actions)
}

func loadData() ([]models.User, []models.Action, error) {
	var (
		users   []models.User
		actions []models.Action
	)

	usersData, err := os.ReadFile("data/users.json")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't load users data: %s", err.Error())
	}

	if err = json.Unmarshal(usersData, &users); err != nil {
		return nil, nil, fmt.Errorf("couldn't unmarshall users data: %s", err.Error())
	}

	actionsData, err := os.ReadFile("data/actions.json")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't load actions data: %s", err.Error())
	}

	if err = json.Unmarshal(actionsData, &actions); err != nil {
		return nil, nil, fmt.Errorf("couldn't unmarshall actions data: %s", err.Error())
	}

	return users, actions, nil

}
