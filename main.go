package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "koeawase")
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
