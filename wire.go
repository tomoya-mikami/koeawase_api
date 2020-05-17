//+build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"cloud.google.com/go/firestore"

	Voice "local.packages/voice"
	Task "local.packages/task"
	Src "local.packages/src"
)

func InitializeCLI(client *firestore.Client, ctx context.Context) (*Src.CLI, error) {
	wire.Build(
		Voice.NewRepository,
		Voice.NewService,
		Task.NewVoiceTask,
		Src.NewCLI,
	)

	return nil, nil
}
