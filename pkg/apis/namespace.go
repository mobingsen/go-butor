package apis

import (
	"github.com/gin-gonic/gin"
	"go-topics/pkg/service"
	"net/http"
)

func GetNamespaces(c *gin.Context) {
	namespaces, err := service.GetNamespace()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, namespaces)
}