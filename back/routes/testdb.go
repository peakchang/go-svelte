package routes

import (
	"fmt"
	"go-svelte/db"
	"go-svelte/models"
	"go-svelte/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterTestdbToutes(router fiber.Router) {
	router.Get("/", getTestDb)
	router.Post("/upload_test", uploadTestDb)
}

func getTestDb(c *fiber.Ctx) error {

	fmt.Println("일단 들어옴?!?!?!?!")

	var testdbs []models.DbTestdb

	err := db.DB.Select(&testdbs, "SELECT * FROM test_table")

	fmt.Println(testdbs)
	if err != nil {
		fmt.Println(err)
		errObj := models.ErrObj{ErrMessage: "DB 조회 실패"}
		return c.Status(500).JSON(errObj)
	}

	apiTestdb := utils.MapSlice(testdbs, models.ConvertToAPI)
	return c.JSON(apiTestdb)
}

func uploadTestDb(c *fiber.Ctx) error {

	fmt.Println("들어옴?!?!")
	type uploadRequest struct {
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}

	var req uploadRequest

	if err := c.BodyParser(&req); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	fmt.Println(req)

	_, err := db.DB.Exec("INSERT INTO test_table (name, age) VALUES (?, ?)", req.Name, req.Age)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error(), "message": "중복된 이름 값이 있습니다."})
	}

	return c.JSON(fiber.Map{})
}
