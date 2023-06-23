package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/mockPizza/internal/domain"
	"github.com/yosa12978/mockPizza/internal/services"
)

type PizzaHandler interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	CreatePizza(ctx *gin.Context)
}

type pizzaHandler struct {
	pizzaService services.PizzaService
}

func NewPizzaHandler(pizzaService services.PizzaService) PizzaHandler {
	return &pizzaHandler{
		pizzaService: pizzaService,
	}
}

func (handler *pizzaHandler) GetAll(ctx *gin.Context) {
	ctx.JSON(200, handler.pizzaService.GetAll())
}

func (handler *pizzaHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "pizza not found",
		})
		return
	}
	pizza, err := handler.pizzaService.GetById(uint(id))
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "pizza not found",
		})
		return
	}
	ctx.JSON(200, pizza)
}

func (handler *pizzaHandler) CreatePizza(ctx *gin.Context) {
	var pizzaDto domain.PizzaCreateDTO
	if err := ctx.ShouldBindJSON(&pizzaDto); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	res, err := handler.pizzaService.Create(pizzaDto.ToObj())
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"id": res.ID,
	})
}
