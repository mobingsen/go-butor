package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webssh/pkg/service"
)

func GetPods(c *gin.Context) {
	namespaceName := c.Param("namespaceName")
	pods, err := service.GetPods(namespaceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, pods)
}

func ExecContainer(c *gin.Context) {
	namespaceName := c.Param("namespaceName")
	podName := c.Param("podName")
	containerdName := c.Param("containerName")
	method := c.DefaultQuery("action", "sh")
	err := service.WebSSH(namespaceName, podName, containerdName, method, c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
}
