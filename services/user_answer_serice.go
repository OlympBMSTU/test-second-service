package services

type IExerciseService interface {
	CreateExercise() DbResult
	GetExercise(id int) DbResult
	GetListExercises() DbResult
	GetTags() DbResult
	GetSubjects() DbResult
	CreateSubject(subject string) DbResult
}

// its a good variant select t_n, id from tags where t_n in names;
// then get ids that exist and not exist create -> get new ids -> insert into mm table
