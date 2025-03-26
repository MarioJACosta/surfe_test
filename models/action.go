package models

import "time"

type Action struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	UserID     int       `json:"userId"`
	TargetUser int       `json:"targetUser"`
	CreatedAt  time.Time `json:"createdAt"`
}

var Actions []Action
var ActionsByID map[int]Action

func InitActions(actionData []Action) {
	Actions = actionData
	ActionsByID = make(map[int]Action)

	for _, action := range Actions {
		ActionsByID[action.ID] = action
	}
}
