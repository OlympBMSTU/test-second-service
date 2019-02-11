package entities

type UserAnswer struct {
	ID         int
	VariantID  int
	ExerciseID int
	UserID     int
	Answer     []ExericseAnswer //string //
	Score      int
	FileParh   string
	Created    string
}
