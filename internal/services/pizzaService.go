package services

import (
	"github.com/yosa12978/mockPizza/internal/domain"
	"github.com/yosa12978/mockPizza/internal/repos"
)

type PizzaService interface {
	GetAll() []domain.Pizza
	GetById(id uint) (domain.Pizza, error)
	Create(pizza domain.Pizza) (domain.Pizza, error)
	Update(pizza domain.Pizza) (domain.Pizza, error)
	Delete(id uint) (domain.Pizza, error)
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

func (service *pizzaService) Create(pizza domain.Pizza) (domain.Pizza, error) {
	return service.pizzaRepo.Create(pizza)
}

func (service *pizzaService) Delete(id uint) (domain.Pizza, error) {
	return service.pizzaRepo.Delete(id)
}

func (service *pizzaService) Update(pizza domain.Pizza) (domain.Pizza, error) {
	return service.pizzaRepo.Update(pizza)
}
