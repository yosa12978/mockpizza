package domain

type Usr struct {
	BaseModel
	Username      string
	Role          string
	Password_hash string
	Email         string
}
