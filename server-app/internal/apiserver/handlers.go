package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *server) handlerGetData(c *gin.Context) {
	rep := srv.storage.Repository()
	dataArr, err := rep.GetData()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dataArr)
}
