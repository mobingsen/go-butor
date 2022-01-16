package router

import (
	"github.com/gin-gonic/gin"
	"go-topics/pkg/apis"
	"go-topics/pkg/middleware"
)

func InitRouter(r *gin.Engine) {
	middleware.InitMiddleware(r)
	r.GET("/ping", apis.Ping)
	r.GET("/namespaces", apis.GetNamespaces)
	r.GET("namespace/:namespaceName/pods", apis.GetPods)
	r.GET("/namespace/:namespaceName/pod/:podName/container/:containerName", apis.ExecContainer)
}
