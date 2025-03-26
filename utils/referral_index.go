package utils

import (
	"surfe_test/models"
)

const REFER_USER_ACTION = "REFER_USER"

func getReferralGraph() map[int][]int {
	graph := make(map[int][]int)

	for _, action := range models.Actions {
		if action.Type == REFER_USER_ACTION {
			graph[action.UserID] = append(graph[action.UserID], action.TargetUser)
		}
	}

	return graph
}

func depthSearch(graph map[int][]int, start int, visited map[int]bool) {
	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			visited[neighbor] = true
			depthSearch(graph, start, visited)
		}
	}
}

func GetReferralIndex() map[int]int {
	graph := getReferralGraph()
	result := make(map[int]int)

	for user := range graph {
		visited := make(map[int]bool)
		depthSearch(graph, user, visited)
		result[user] = len(visited)
	}

	return result
}
