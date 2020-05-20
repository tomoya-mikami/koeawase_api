//+build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"cloud.google.com/go/firestore"

	Voice "local.packages/voice"
	Similarity "local.packages/similarity"
	Task "local.packages/task"
	Src "local.packages/src"
	Handler "local.packages/handler"
)

func InitializeCLI(client *firestore.Client, ctx context.Context) (*Src.CLI, error) {
	wire.Build(
		Voice.NewRepository,
		Voice.NewService,
		Similarity.NewRepository,
		Similarity.NewService,
		Task.NewVoiceTask,
		Src.NewCLI,
	)

	return nil, nil
}

func InitializeServer(client *firestore.Client, ctx context.Context) (*Src.Server, error) {
	wire.Build(
		Voice.NewRepository,
		Voice.NewService,
		Handler.NewVoiceHandler,
		Src.NewServer,
	)

	return nil, nil
}
