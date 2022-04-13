package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/tidwall/gjson"
)

func jsonHadler(c *fiber.Ctx) error {
	data := struct {
		Query string `json:"query"`
		Json  string `json:"json"`
	}{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	value := gjson.Get(data.Json, data.Query)
	println(value.String())

	response := make(map[string]string)
	response["query"] = data.Query
	response["json"] = data.Json
	response["result"] = value.String()
	return c.JSON(response)
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/parse", jsonHadler)

	err := app.Listen(":3000")

	if err != nil {
		fmt.Println("[ERROR] Failed to start server")
	}
}
