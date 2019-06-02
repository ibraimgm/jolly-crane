package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibraimgm/jolly-crane/admin"
	"github.com/ibraimgm/jolly-crane/texthash"
)

func main() {
	r := gin.Default()

	admin.Controller().SetupRoutes(r.Group("/admin"))
	texthash.Controller().SetupRoutes(r)

	r.Run()
}
