package domain
import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name", validate:"required,min=5"`
	Products []Product `json:"products,omitempty`
}
