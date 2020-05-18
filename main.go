package main

import (
	"context"
	"log"
	"os"
	"cloud.google.com/go/firestore"
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
		server, err := InitializeServer(client, ctx)
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	}
}
