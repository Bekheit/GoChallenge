package api

import (
	"bytes"
	"encoding/json"
	"example/go/internal/services/transactionsvc"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	var tc TransactionController
	handler := http.HandlerFunc(tc.handleCreateTransaction)
	handler.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseBody(t *testing.T, body map[string]interface{}) {
	if body["amount"] != 500 {
		t.Fatalf("Expected transaction amount to be '500'. Got '%v'", body["amount"])
	}

	if body["currency"] != 500 {
		t.Fatalf("Expected transaction currency to be 'EGP'. Got '%v'", body["currency"])
	}
}

func TestCreateTransaction(t *testing.T) {
	var tr = transactionsvc.CreatePld{
		Amount:   500,
		Currency: "EGP",
	}
	body, _ := json.Marshal(tr)
	req, err := http.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	res := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, res.Code)

	var transaction map[string]interface{}

	json.Unmarshal(res.Body.Bytes(), &transaction)

	checkResponseBody(t, transaction)
}
