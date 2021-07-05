package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func openBoard(boardName string) TrelloBoard {
	jsonFile, err := os.Open(boardName)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var board TrelloBoard

	json.Unmarshal(byteValue, &board)

	return board
}

// check the ID of the completed list, will return an empty string if not found
func getCompleteListID(board TrelloBoard) string {
	lists := board.Lists

	for _, list := range lists {
		if strings.Contains(strings.ToLower(list.Name), "complete") {
			return list.ID
		}
	}

	return ""
}

// check the ID of the point custom field, will return an empty string if not found
// TODO: Multiple points custom field
func getCustomFieldPointsID(board TrelloBoard) string {
	customFields := board.CustomFields

	for _, customField := range customFields {
		if strings.Contains(strings.ToLower(customField.Name), "point") {
			return customField.ID
		}
	}

	return ""
}

// compile the completed card on the list, and return an array of card name and points
func getCardNameandPoints(board TrelloBoard, completeListID string, customFieldPointsID string) []CardNamePoints {
	var results []CardNamePoints

	cards := board.Cards

	for _, card := range cards {
		var cardNamePoints CardNamePoints

		if card.IDList == completeListID {
			cardNamePoints.Name = card.Name
			for _, customField := range card.CustomFieldItems {
				if customField.IDCustomField == customFieldPointsID {
					cardNamePoints.Points = customField.Value.Number
				}
			}
			results = append(results, cardNamePoints)
		}
	}

	return results
}

func calculatePoints(data []CardNamePoints) int64 {
	var totalPoints int

	for _, cardNamePoint := range data {
		point, err := strconv.Atoi(cardNamePoint.Points)
		if err != nil {
			return -1
		}
		totalPoints = totalPoints + point
	}

	return int64(totalPoints)
}
