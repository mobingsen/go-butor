package apis

import (
	"github.com/gin-gonic/gin"
	"go-topics/pkg/service"
	"net/http"
)

func GetPods(c *gin.Context) {
	namespaceName := c.Param("namespaceName")
	pods, err := service.GetPods(namespaceName)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, pods)
}
