package texthash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibraimgm/jolly-crane/rest"
)

type controller struct{}

func (c *controller) SetupRoutes(router gin.IRouter) {
	router.POST("/hash", func(c *gin.Context) {
		c.String(http.StatusOK, "save?")
	})

	router.GET("/hashes/:hash", func(c *gin.Context) {
		hash := c.Param("hash")
		c.String(http.StatusOK, "find %s?", hash)
	})

	router.GET("/hashes", func(c *gin.Context) {
		c.String(http.StatusOK, "list?")
	})
}

// Controller retorna o controller necessário para setar as rotas
// do endpoint principal
func Controller() rest.Controller {
	return &controller{}
}
