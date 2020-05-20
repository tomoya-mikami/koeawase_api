package Voice

import (
	"log"
	"context"
	"cloud.google.com/go/firestore"
)

type RepositoryInterface interface {
	Add(voice *Voice) (*Voice, error)
	Get(id string) (*Voice, error)
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
	voiceFirestore := new(VoiceFireStore)
	voiceFirestore.Name = voice.Name
	voiceFirestore.PowerSpectrum = ConvertFloatArrayToString(voice.PowerSpectrum)
	docRef, _, err := r.client.Collection("Voices").Add(r.ctx, voiceFirestore)
	if err != nil {
		return voice, err
	}
	voice.ID = docRef.ID

	return voice, err
}

func (r Repository) Get(id string) (*Voice, error) {
	dsnap, err := r.client.Collection("Voices").Doc(id).Get(r.ctx)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var voiceFireStore VoiceFireStore
	dsnap.DataTo(&voiceFireStore)
	voice := new(Voice)
	voice.Name = voiceFireStore.Name
	voice.PowerSpectrum = ConvertStringToFloatArray(voiceFireStore.PowerSpectrum)
	voice.ID = id
	return voice, err
}
