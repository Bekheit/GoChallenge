package transactionsvc

import (
	"context"
	"testing"
)

func TestDefaultService_Create(t *testing.T) {
	//log, closer := resources.NewLogger()
	//defer closer()

	//configs := config.LoadConfig(log)

	//conn := db.NewDatabaseConnection(log, configs.Database)

	model := CreatePld{
		Amount:   500,
		Currency: "EGP",
	}

	var service DefaultService

	result, _ := service.Create(context.Background(), &model)
	var expectedAmount = float32(500)
	var expectedCurrency = "EGP"
	if result.Amount != expectedAmount {
		t.Errorf("Create transaction failed, expected amount %v, got %v", expectedAmount, result.Amount)
	}
	if result.Currency != expectedCurrency {
		t.Errorf("Create transaction failed, expected currency %v, got %v", expectedCurrency, result.Currency)
	}

}
