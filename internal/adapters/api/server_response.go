package api

import (
	"context"
	"encoding/json"
	"net/http"
)

func RenderJSON(ctx context.Context, writer http.ResponseWriter, httpStatusCode int, payload interface{}) {
	writer.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(payload)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(httpStatusCode)
	_, _ = writer.Write(js)
}

func RenderError(ctx context.Context, writer http.ResponseWriter, err error) {
	payload := map[string]error{
		"error": err,
	}
	RenderJSON(ctx, writer, 500, payload)
}
