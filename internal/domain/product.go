package domain
import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name"`
	Stock int `json:"stock"`
}
