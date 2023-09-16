package main

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/walkin/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//go:embed build/index.html
var indexHtmlBytes []byte
//go:embed build/index.js
var indexJsBytes []byte

type ErrorResponse struct {
	Error string `json:"error"`
}
func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return ctx.Status(code).JSON(ErrorResponse { Error: "Internal Sever Error" })
}

func main() {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("running..\n\n")
	fmt.Printf("workdir: %s\n", workdir)

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(recover.New())

	queryCtl := controller.QueryController{}
	app.Post("/query", queryCtl.Query)
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(string(indexHtmlBytes))
	})
	app.Get("/index.js", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextJavaScript)
		return c.SendString(string(indexJsBytes))
	})

	app.Listen(":3000")
}
