package handler

import (
		"strconv"
    "github.com/wanxsky/kitchen-api/internal/domain"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type ProductHandler struct {
        DB *gorm.DB
}

func (h *ProductHandler) GetAll (c *fiber.Ctx) error {
        var products []domain.Product
        h.DB.Preload("Category").Find(&products)
        return c.JSON(products)
}

func (h *ProductHandler) Create (c *fiber.Ctx) error {
        product := new(domain.Product)
        if err := c.BodyParser(product); err != nil {
                return c.Status(400).JSON(fiber.Map{"status": false, "error": "Bad Request"})
        }

		if err := validate.Struct(product); err != nil {
		return c.Status(400).JSON(fiber.Map{ "status": false, "error": "Bad Request", "detail": err.Error() })
	}

	if err := h.DB.Create(product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{ "status": false, "error": "Internal Server Error" })
	}
	h.DB.Preload("Category").First(product, product.ID)
        return c.Status(201).JSON(fiber.Map{ "status": true, "data": product })
}

func (h *ProductHandler) Update (c *fiber.Ctx) error {
  idStr := c.Params("id")
  id, err := strconv.Atoi(idStr)
  var product domain.Product
  if err != nil {
    return c.Status(400).JSON(fiber.Map{"status": false, "error": "ID harus berupa angka" })
  }

  if err := h.DB.First(&product, id).Error; err != nil {
    return c.Status(404).JSON(fiber.Map{"status": false, "error": "Product tidak ditemukan"})
  }
  if err := c.BodyParser(&product); err != nil {
    return c.Status(400).JSON(fiber.Map{"status": false, "error": "Format tidak sesuai"})
  }
	if err := validate.Struct(product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": false,
			"error": "Bad Request",
			"detail": err.Error(),
		})
	}
	if err := h.DB.Save(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": false,
			"error": "Internal Server Error",
		})
	}

	h.DB.Preload("Category").First(&product, product.ID)

  return c.Status(200).JSON(fiber.Map{ "status": true, "data": product})
}

func (h *ProductHandler) Get (c *fiber.Ctx) error {
  idStr := c.Params("id")
  id, err := strconv.Atoi(idStr)
  var product domain.Product
  if err != nil {
    return c.Status(400).JSON(fiber.Map{"status": false, "error": "ID harus berupa angka"})
  }
  
  if err := h.DB.Preload("Category").First(&product, id).Error; err != nil {
    return c.Status(404).JSON(fiber.Map{"status": false, "error": "Product tidak ditemukan"})
  }
  return c.Status(200).JSON(fiber.Map{ "status": true, "data": product})
}

func (h *ProductHandler) Delete (c *fiber.Ctx) error {
  idStr := c.Params("id")
  id, err := strconv.Atoi(idStr)
	var product domain.Product
  if err != nil {
    return c.Status(400).JSON(fiber.Map{ "status": false, "error": "ID harus berupa angka"})
  }
	if err := h.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{ "status": false, "error": "Product tidak ditemukan"})
	}
  if err := h.DB.Delete(&product, id).Error; err != nil {
    return c.Status(500).JSON(fiber.Map{ "status": false, "error": "Internal Server Error"})
  }
  return c.Status(200).JSON(fiber.Map{ "status": true, "data": "berhasil dihapus"})
}
