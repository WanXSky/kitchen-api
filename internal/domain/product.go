package domain
import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name" validate:"required,min=3"`
	Stock int `json:"stock" validate:"required,min=0"`
	CategoryID uint `json:"category_id"`
	Category Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	CreatedBy uint `json:"created_by"`
}
