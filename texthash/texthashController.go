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
		input := &TextHash{}
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
		found := service.FindByHash(hash)

		if found == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found."})
			return
		}

		c.JSON(http.StatusOK, found)
	})

	router.GET("/hashes", func(c *gin.Context) {
		c.JSON(http.StatusOK, service.FindAll())
	})
}

// Controller retorna o controller necessário para setar as rotas
// do endpoint principal
func Controller() rest.Controller {
	return &controller{}
}
