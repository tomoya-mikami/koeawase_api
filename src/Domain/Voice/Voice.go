package Voice

import (
	"fmt"
	"strconv"
	"strings"
)

type Voice struct {
	ID            string
	Name          string    `firestore:"name"`
	PowerSpectrum []float64
}

type VoiceFireStore struct {
	Name string `firestore:"name"`
	PowerSpectrum string `firestore:"power_spectrum"`
}

const LOW_FREQUENCY = 100
const HIGH_FREQUENCY = 2000

func ConvertFloatArrayToString(array []float64) string {
	result := ""
	array_len := len(array)
	for i := 0; i < array_len; i++ {
		result += fmt.Sprintf("%f", array[i])
		if i != array_len - 1 {
			result += ","
		}
	}

	return result
}

func ConvertStringToFloatArray(s string) []float64 {
	split := strings.Split(s, ",")
	array := make([]float64, 0)
	for i := 0; i < len(split); i++ {
		f, _ := strconv.ParseFloat(split[i], 64)
		array = append(array, f)
	}

	return array
}

