package controllers

import (
	"context"
	"example/go/configs"
	. "example/go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTransaction() gin.HandlerFunc {

	return func(c *gin.Context) {

		conn := configs.ConnectDB()
		defer conn.Close(context.Background())

		var transaction Transaction

		if err := c.BindJSON(&transaction); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}
		if _, err := conn.Exec(context.Background(), "insert into Transactions (amount, currency) values ($1, $2)", transaction.Amount, transaction.Currency); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "transaction created successfully")
		return
	}
}

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		conn := configs.ConnectDB()
		defer conn.Close(context.Background())

		var transaction Transaction
		var transactions []Transaction

		rows, err := conn.Query(context.Background(), "select * from transactions;")
		if err != nil {
			c.JSON(http.StatusBadRequest, "query execution fail")
			return
		}

		for rows.Next() {
			err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Currency, &transaction.CreatedAt)
			if err != nil {
				c.JSON(http.StatusBadRequest, "binding fail")
				return
			}
			transactions = append(transactions, transaction)
		}
		c.JSON(http.StatusOK, transactions)
		return
	}
}
