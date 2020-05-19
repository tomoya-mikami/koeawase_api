package Voice

import (
	"log"
	"context"
	"cloud.google.com/go/firestore"
)

type RepositoryInterface interface {
	Add(voice *Voice) (*Voice, error)
	Get(id string) (Voice, error)
}

type Repository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewRepository(client *firestore.Client, ctx context.Context) RepositoryInterface {
	repository := new(Repository)
	repository.client = client
	repository.ctx = ctx
	return repository
}

func (r Repository) Add(voice *Voice) (*Voice, error) {
	docRef, _, err := r.client.Collection("Voices").Add(r.ctx, voice)
	if err != nil {
		log.Fatal(err)
		return voice, err
	}
	voice.ID = docRef.ID

	return voice, err
}

func (r Repository) Get(id string) (Voice, error) {
	dsnap, err := r.client.Collection("Voices").Doc(id).Get(r.ctx)
	if err != nil {
		log.Print(err)
		return Voice{}, err
	}

	var voice Voice
	dsnap.DataTo(&voice)
	voice.ID = id
	return voice, err
}
