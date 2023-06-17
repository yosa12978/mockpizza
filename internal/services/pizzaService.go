package services

import (
	"github.com/yosa12978/mockPizza/internal/domain"
	"github.com/yosa12978/mockPizza/internal/repos"
)

type PizzaService interface {
	GetAll() []domain.Pizza
}

type pizzaService struct {
	pizzaRepo repos.PizzaRepo
}

func NewPizzaService(pizzaRepo repos.PizzaRepo) PizzaService {
	return &pizzaService{pizzaRepo: pizzaRepo}
}

func (service *pizzaService) GetAll() []domain.Pizza {
	return service.pizzaRepo.FindAll()
}

func (service *pizzaService) GetById(id uint) (domain.Pizza, error) {
	return service.pizzaRepo.FindById(id)
}
