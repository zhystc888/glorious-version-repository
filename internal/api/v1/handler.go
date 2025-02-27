package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhystc888/glorious-version-repository/internal/config"
	"github.com/zhystc888/glorious-version-repository/pkg/response"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	r.GET("/hello", HelloHandler)
	return r
}

func HelloHandler(c *gin.Context) {
	response.Success(c, "Hello, World!")
}
