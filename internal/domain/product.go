package domain
import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name" validate:"required,min=3"`
	Stock int `json:"stock" validate:"required,min=0"`
}
