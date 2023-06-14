package domain

type Order struct {
	BaseModel
	Pizzas   []Pizza
	Price    float32
	Finished bool
	Paid     bool
	Desc     string
}
