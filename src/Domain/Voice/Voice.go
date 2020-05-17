package Voice

type Voice struct {
	ID            string
	Name          string    `firestore:"name"`
	PowerSpectrum []float64 `firestore:"power_spectrum"`
}