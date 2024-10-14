package apiserver

import (
	"mtg/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type server struct {
	router  *gin.Engine
	logger  *logrus.Logger
	storage store.Storage
}

func newServer(logger *logrus.Logger, storage store.Storage) *server {
	srv := &server{
		router:  gin.New(),
		logger:  logger,
		storage: storage,
	}

	srv.configureRouter()

	return srv
}

func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(w, r)
}

func (srv *server) configureRouter() {
	srv.router.POST("/data", srv.handlerGetData)
}
