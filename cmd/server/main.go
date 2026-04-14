package main
import (
	"github.com/wanxsky/kitchen-api/internal/domain"
	"github.com/wanxsky/kitchen-api/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"os"
)

func main() {
	var db *gorm.DB
	var err error 

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "production" {
		dsn := os.Getenv("DATABASE_URL")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open("kitchen.db"), &gorm.Config{})
	}

	if err != nil {
		panic("Gagal Konek ke DB")
	}

	db.AutoMigrate(&domain.Product{})
	productHandler := handler.ProductHandler{DB: db}
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	
	app.Get("/products", productHandler.GetAll)
	app.Get("/products/:id", productHandler.Get)
	app.Post("/products", productHandler.Create)
	app.Put("/products/:id", productHandler.Update)
	app.Delete("/products/:id", productHandler.Delete)
	app.Listen(":3000")
}
