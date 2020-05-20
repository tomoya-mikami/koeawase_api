package Similarity

import (
	"log"
	"context"
	"cloud.google.com/go/firestore"
)

type RepositoryInterface interface {
	Add(similarity *Similarity) (*Similarity, error)
	Get(id string) (Similarity, error)
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

func (r Repository) Add(similarity *Similarity) (*Similarity, error) {
	docRef, _, err := r.client.Collection("Similarities").Add(r.ctx, similarity)
	if err != nil {
		return similarity, err
	}
	similarity.ID = docRef.ID

	return similarity, err
}

func (r Repository) Get(id string) (Similarity, error) {
	dsnap, err := r.client.Collection("Similarities").Doc(id).Get(r.ctx)
	if err != nil {
		log.Print(err)
		return Similarity{}, err
	}

	var similarity Similarity
	dsnap.DataTo(&similarity)
	similarity.ID = id
	return similarity, err
}
