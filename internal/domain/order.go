package domain

type Order struct {
	BaseModel
	Pizzas   []Pizza `gorm:"foreignKey:ID"`
	Price    float32
	Finished bool
	Paid     bool
	Desc     string
}
