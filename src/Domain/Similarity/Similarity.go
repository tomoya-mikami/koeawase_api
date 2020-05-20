package Similarity

type Similarity struct {
	ID string
	Name1 string `firestore:"name2"`
	Name2 string `firestore:"name1"`
	Similarity float64 `firestore:"similarity"`
}
