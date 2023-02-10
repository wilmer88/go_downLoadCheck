package main

import (

    "github.com/gofiber/fiber/cli.v2"
	"os"

    cli "github.com/urfave/cli"


)


func main() {
	cli.NewApp().Run(os.Args)
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":3000")
}