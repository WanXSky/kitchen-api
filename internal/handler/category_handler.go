package handler
import (
	"github.com/wanxsky/kitchen-api/internal/domain"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CategoryHandler struct {
	DB *gorm.DB
}

func (h *CategoryHandler) GetAll (c *fiber.Ctx) error {
	var categories []domain.Category 
	h.DB.Preload("Products").Find(&categories)
	return c.Status(200).JSON(fiber.Map{ "status": true, "data": categories })
}

func (h *CategoryHandler) Create (c *fiber.Ctx) error {
	category := new(domain.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{ "status": false, "error": "Bad Request"})
	}

	if err := validate.Struct(category); err != nil {
		return c.Status(400).JSON(fiber.Map{ "status": false, "error": "Bad Request", "detail": err.Error() })
	}

	if err := h.DB.Create(category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{ "status": false, "error": "Internal Server Error" })
	}

	return c.Status(201).JSON(fiber.Map{ "status": true, "data": category })
}
