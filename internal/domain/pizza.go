package domain

type Pizza struct {
	BaseModel
	Name  string
	Price float32
	Desc  string
}

type PizzaCreateDTO struct {
	Name  string
	Price float32
	Desc  string
}

func (p PizzaCreateDTO) ToObj() Pizza {
	return Pizza{
		Name:  p.Name,
		Desc:  p.Desc,
		Price: p.Price,
	}
}
