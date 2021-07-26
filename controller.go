package main

import (
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "200",
		"message": "Up and running",
	})
}

func getCards(c *gin.Context) {
	var board = openBoard("jsonfiles/NTMYxAndrea.json")

	completeListID := getCompleteListID(board)
	customFieldPointID := getCustomFieldPointsID(board)

	// fmt.Println(completeListID)
	// fmt.Println(customFieldPointID)

	res := getCardNameandPoints(board, completeListID, customFieldPointID)
	totalPoint := calculatePoints(res)

	data := map[string]interface{}{
		"cards":       res,
		"totalPoints": totalPoint,
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "OK",
		"data":    data,
	})
}
