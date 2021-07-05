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
	var board = openBoard("go-blog-board.json")

	completeListID := getCompleteListID(board)
	customFieldPointID := getCustomFieldPointsID(board)

	res := getCardNameandPoints(board, completeListID, customFieldPointID)

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "OK",
		"data":    res,
	})
}
