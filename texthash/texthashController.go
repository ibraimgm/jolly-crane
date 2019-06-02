package texthash

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibraimgm/jolly-crane/rest"
)

var service = NewService(NewInMemRepository())

type controller struct{}

func (c *controller) SetupRoutes(router gin.IRouter) {
	router.POST("/hash", func(c *gin.Context) {
		var input *TextHash
		var err error

		if err = c.ShouldBindJSON(input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if input, err = service.Create(input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, input)
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
