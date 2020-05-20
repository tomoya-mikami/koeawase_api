//+build wireinject

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/wire"

	Handler "local.packages/handler"
	Similarity "local.packages/similarity"
	Src "local.packages/src"
	Task "local.packages/task"
	Voice "local.packages/voice"
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
		Similarity.NewRepository,
		Similarity.NewService,
		Handler.NewVoiceHandler,
		Src.NewServer,
	)

	return nil, nil
}
