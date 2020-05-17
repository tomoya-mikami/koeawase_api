package Voice

import (
	"os"
	"io"
	"math/cmplx"
	"math"

	"github.com/mjibson/go-dsp/fft"
	"github.com/youpy/go-wav"
)

type ServiceInterface interface {
	CalculatePowerSpectrum(*os.File) []float64
	CosSimilarity(sample []float64, training []float64) float64
	Add(name string, powerSpectrum []float64) (*Voice, error)
	Get(id string) (Voice, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) ServiceInterface {
	service := new(Service)
	service.repository = repository
	return service
}

func (s Service) Add(name string, powerSpectrum []float64) (*Voice, error) {
	voice := new(Voice)
	voice.Name = name
	voice.PowerSpectrum = powerSpectrum
	return s.repository.Add(voice)
}

func (s Service) Get(id string) (Voice, error) {
	return s.repository.Get(id)
}

func (s Service) CalculatePowerSpectrum(file *os.File) []float64 {
	// Wavファイルの読み込み 
	reader := wav.NewReader(file)

	defer file.Close()

	signals := make([]float64, 0)

	// 解析はモノラルチャンネルで行うので片方のチャンネルだけ取得
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}

		for _, sample := range samples {
			signals = append(signals, reader.FloatValue(sample, 0))
		}
	}

	// 高速フーリエ変換
	fftResult := fft.FFTReal(signals)

	// フーリエ変換の結果は複素数なので絶対値を取得し, パワースペクトルに変換する
	powerSpectrum := make([]float64, 0)
	for _, r := range fftResult {
		powerSpectrum = append(powerSpectrum, cmplx.Abs(r))
	}

	return powerSpectrum
}

func (s Service) CosSimilarity(sample []float64, training []float64) float64 {
	sample_len := len(sample)
	training_len := len(training)
	target_len := minInt(sample_len, training_len)

	// target_sample
	ts := sample[0:target_len]
	// target_traingin
	tt := training[0:target_len]

	return calculateDot(ts, tt) / (math.Sqrt(calculateDot(ts, ts)) * math.Sqrt(calculateDot(tt, tt)))
}

func minInt(a int, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}

func calculateDot(x []float64, y []float64) float64 {
	result := 0.0;
	for i := 0; i < len(x); i++ {
		result += x[i] * y[i]
	}
	return result
}
