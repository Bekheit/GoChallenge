package api

import (
	"bytes"
	"encoding/json"
	"example/go/internal/services/transactionsvc"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	var tr = transactionsvc.CreatePld{
		Amount:   500,
		Currency: "EGP",
	}
	body, _ := json.Marshal(tr)

	req, err := http.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	var tc TransactionController
	handler := http.HandlerFunc(tc.handleCreateTransaction)
	handler.ServeHTTP(rr, req)

	if http.StatusCreated != rr.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusCreated, rr.Code)
	}

	var transaction map[string]interface{}

	json.Unmarshal(rr.Body.Bytes(), &transaction)

	if transaction["amount"] != 500 {
		t.Fatalf("Expected transaction amount to be '500'. Got '%v'", transaction["amount"])
	}

	if transaction["currency"] != "EGP" {
		t.Fatalf("Expected transaction currency to be 'EGP'. Got '%v'", transaction["currency"])
	}

}

//func TestCreateTransaction(t *testing.T) {
//
//	var tr = transactionsvc.CreatePld{
//		Amount:   500,
//		Currency: "EGP",
//	}
//	body, _ := json.Marshal(tr)
//	w := httptest.NewRecorder()
//	r := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
//
//	rctx := chi.NewRouteContext()
//	rctx.URLParams.Add("key", "value")
//
//	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
//	var tc TransactionController
//	handler := func(w http.ResponseWriter, r *http.Request) {
//		tc.handleCreateTransaction(w, r)
//	}
//	handler(w, r)
//}
