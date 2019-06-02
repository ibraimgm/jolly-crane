package rest

import (
	"github.com/gin-gonic/gin"
)

// Controller define as rotas dentro de um determinado contexto
// As rotas geradas devem ser montadas em '/', para que o usuário da API
// tenha mais liberdade em decidir como o endereço final será estruturado.
type Controller interface {
	SetupRoutes(router gin.IRouter)
}
