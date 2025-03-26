package models

import (
	"math"
	"sort"
	"time"
)

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

func CountActionsByUserID(id int) int {
	count := 0

	for _, action := range Actions {
		if action.UserID == id {
			count++
		}
	}

	return count
}

func GetNextActionProbabilities(actionType string) map[string]float64 {
	userActions := make(map[int][]Action)
	for _, a := range Actions {
		userActions[a.UserID] = append(userActions[a.UserID], a)
	}

	counts := make(map[string]int)
	total := 0

	for _, actions := range userActions {

		sort.Slice(actions, func(i, j int) bool {
			return actions[i].CreatedAt.Before(actions[j].CreatedAt)
		})

		for i := 0; i < len(actions)-1; i++ {
			if actions[i].Type == actionType {
				next := actions[i+1].Type
				counts[next]++
				total++
			}
		}
	}

	nextActionProbabilities := make(map[string]float64)

	for actionType, count := range counts {
		probability := float64(count) / float64(total)
		nextActionProbabilities[actionType] = math.Round(probability*100)/100
	}

	return nextActionProbabilities
}
