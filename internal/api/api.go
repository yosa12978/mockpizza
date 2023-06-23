package api

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/mockPizza/internal/api/handlers"
	"github.com/yosa12978/mockPizza/internal/db"
	"github.com/yosa12978/mockPizza/internal/repos"
	"github.com/yosa12978/mockPizza/internal/services"
)

func Listen(listener net.Listener) {
	router := gin.Default()
	router.Use(gin.Logger())
	pizzaService := services.NewPizzaService(repos.NewPizzaRepo(db.GetDB()))
	pizzaHandler := handlers.NewPizzaHandler(pizzaService)
	router.GET("/pizzas", pizzaHandler.GetAll)
	router.GET("/pizzas/:id", pizzaHandler.GetByID)
	router.POST("/pizzas", pizzaHandler.CreatePizza)

	router.RunListener(listener)
}
