package Similarity

type Similarity struct {
	ID string
	Name1ID string `firestore:"name1_id"`
	Name1 string `firestore:"name1"`
	Name2 string `firestore:"name2"`
	Similarity float64 `firestore:"similarity"`
}
