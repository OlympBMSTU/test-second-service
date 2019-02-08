package main

type Controllers interface {
	UploadExercise()
	GetExercise(id int)
	GetListExercise(limit, offset int)
	GetTags(subject string)
	GetSubjects()
	UploadSecondRoundExercise()
}
