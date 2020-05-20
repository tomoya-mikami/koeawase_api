package Similarity

import (
	Voice "local.packages/voice"
	"math"
)

type ServiceInterface interface {
	CalcurateSimilarity(sample *Voice.Voice, training *Voice.Voice) (*Similarity, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) ServiceInterface {
	service := new(Service)
	service.repository = repository
	return service
}

func (s Service) CalcurateSimilarity(sample *Voice.Voice, training *Voice.Voice) (*Similarity, error) {
	cosSimilarity := cosSimilarity(sample.PowerSpectrum, training.PowerSpectrum)
	result := new(Similarity)
	result.Name1ID = sample.ID
	result.Name1 = sample.Name
	result.Name2 = training.Name
	result.Similarity = cosSimilarity
	return s.repository.Add(result)
}

func cosSimilarity(sample []float64, training []float64) float64 {
	// target_sample
	ts := sample[Voice.LOW_FREQUENCY:Voice.HIGH_FREQUENCY]
	// target_traingin
	tt := training[Voice.LOW_FREQUENCY:Voice.HIGH_FREQUENCY]

	return calculateDot(ts, tt) / (math.Sqrt(calculateDot(ts, ts)) * math.Sqrt(calculateDot(tt, tt)))
}

func calculateDot(x []float64, y []float64) float64 {
	result := 0.0;
	for i := 0; i < len(x); i++ {
		result += x[i] * y[i]
	}
	return result
}
