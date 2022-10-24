package api

import (
	"encoding/json"
	"example/go/internal/services/transactionsvc"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

type TransactionController struct {
	log            *zap.SugaredLogger
	validate       *validator.Validate
	transactionSvc transactionsvc.IServive
}

func NewTransactionController(server *HttpServer, validator *validator.Validate, ts transactionsvc.IServive) {
	c := &TransactionController{
		log:            server.Logger,
		validate:       validator,
		transactionSvc: ts,
	}

	server.Router.Group(func(r chi.Router) {
		r.Post("/transactions", c.handleCreateTransaction)
		r.Get("/transactions", c.handleGetAllTransactions)
	})
}

func (c *TransactionController) handleCreateTransaction(writer http.ResponseWriter, req *http.Request) {
	var payload transactionsvc.CreatePld

	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		c.log.Errorf("Malformed body")
		RenderError(req.Context(), writer, err)
		return
	}

	//if err := c.validate.Struct(payload); err != nil {
	//	c.log.Errorf("payload not valid %v", err)
	//	RenderError(req.Context(), writer, err)
	//	return
	//}

	c.transactionSvc.Create(req.Context(), &payload)
	RenderJSON(req.Context(), writer, http.StatusCreated, payload)
}

func (c *TransactionController) handleGetAllTransactions(writer http.ResponseWriter, req *http.Request) {
	res, err := c.transactionSvc.GetAll(req.Context())

	if err != nil {
		return
	}

	RenderJSON(req.Context(), writer, http.StatusCreated, res)
}
