package main

import (
	"context"
	"log"
	"os"
	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "planbbs")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	if len(os.Args) > 1 {
		if os.Args[0] == "sample" {
			Sample()
		} else {
			cli, err := InitializeCLI(client, ctx)
			if err != nil {
				log.Fatal(err)
			}
			cli.Execute(os.Args)
		}
	} else {
		app := fiber.New()

		app.Get("/", func(c *fiber.Ctx) {
			c.Send("Hello, World!")
		})

		app.Listen(8080)
	}
}
