package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webssh/pkg/service"
)

func GetNamespaces(c *gin.Context) {
	namespaces, err := service.GetNamespace()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, namespaces)
}