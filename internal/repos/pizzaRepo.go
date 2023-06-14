package repos

import (
	"github.com/yosa12978/mockPizza/internal/domain"
	"gorm.io/gorm"
)

type PizzaRepo interface {
	Create(pizza domain.Pizza) error
	FindAll() []domain.Pizza
	FindById(id uint) (domain.Pizza, error)
	Update(pizza domain.Pizza) (domain.Pizza, error)
	Delete(id uint) (domain.Pizza, error)
}

type pizzaRepo struct {
	db *gorm.DB
}

func NewPizzaRepo(db *gorm.DB) PizzaRepo {
	return &pizzaRepo{db: db}
}

func (repo *pizzaRepo) Create(pizza domain.Pizza) error {
	return repo.db.Create(pizza).Error
}

func (repo *pizzaRepo) FindAll() []domain.Pizza {
	var pizzas []domain.Pizza
	repo.db.Find(&pizzas)
	return pizzas
}

func (repo *pizzaRepo) FindById(id uint) (domain.Pizza, error) {
	var pizza domain.Pizza
	err := repo.db.First(&pizza, "id = ?", id).Error
	return pizza, err
}

func (repo *pizzaRepo) Update(pizza domain.Pizza) (domain.Pizza, error) {
	err := repo.db.Save(&pizza).Error
	return pizza, err
}

func (repo *pizzaRepo) Delete(id uint) (domain.Pizza, error) {
	var pizza domain.Pizza
	if err := repo.db.First(&pizza, "id = ?", id).Error; err != nil {
		return pizza, err
	}
	return pizza, repo.db.Delete(&pizza).Error
}
