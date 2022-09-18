package api

import (
	"example/go/config"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type HttpServer struct {
	Logger *zap.SugaredLogger
	sc     config.ServerConfigurations
	Router *chi.Mux
}

func NewHttpServer(logger *zap.SugaredLogger, serverConf config.ServerConfigurations) *HttpServer {
	router := chi.NewRouter()

	return &HttpServer{
		Logger: logger,
		sc:     serverConf,
		Router: router,
	}
}

func (r *HttpServer) Start() {
	listeningAddr := ":" + strconv.Itoa(r.sc.Port)

	r.Logger.Infof("Server listening on port %s", listeningAddr)

	err := http.ListenAndServe(listeningAddr, r.Router)

	if err != nil {
		r.Logger.Errorf("Failed to start http server. %v", err)
	}
}
