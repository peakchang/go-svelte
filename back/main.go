package main

import (
	"go-svelte/db"
	"go-svelte/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // 자동으로 .env 로드

	db.Init() // 첫 DB 불러오기

	// .env 불러오기
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	port := os.Getenv("PORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins, // Svelte 개발 서버
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	const backApi string = "/api/v3"

	// 테스트 라우터 "/api/v3/test" 로 시작
	testGroup := app.Group(backApi + "/test")
	routes.RegisterTestdbToutes(testGroup)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("안녕하세요 Fiber!")
	})

	app.Listen(":" + port)
}
