package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func jsonHadler(c *fiber.Ctx) error {
	data := struct {
		Query string `json:"query"`
		Json  string `json:"json"`
	}{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	response := make(map[string]string)
	response["query"] = data.Query
	response["json"] = data.Json
	response["result"] = `{
		"name": {"first": "Tom", "last": "Anderson"},
		"age":37,
		"children": ["Sara","Alex","Jack"],
		"fav.movie": "Deer Hunter",
		"friends": [
		  {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		  {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		  {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
		]
	  }`
	return c.JSON(response)
}

func main() {
	app := fiber.New()

	app.Post("/parse", jsonHadler)

	err := app.Listen(":3000")

	if err != nil {
		fmt.Println("[ERROR] Failed to start server")
	}
}
