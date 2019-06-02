package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ibraimgm/jolly-crane/rest"
)

type controller struct{}

func check(c *gin.Context) {
	c.JSON(200, gin.H{"message": "It works!"})
}

func (c *controller) SetupRoutes(router gin.IRouter) {
	router.GET("/check", check)
}

// Controller retorna um rest.Controller que configura as rotas disponíveis
func Controller() rest.Controller {
	return &controller{}
}
