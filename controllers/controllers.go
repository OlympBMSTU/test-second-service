package controllers

import (
	"github.com/OlympBMSTU/exercises/logger"
	"github.com/OlympBMSTU/test-second-service/config"
	"github.com/OlympBMSTU/test-second-service/result"
	"github.com/OlympBMSTU/test-second-service/views"
)

type IMiddleControllers interface {
	SaveAnswer(answer views.UserAnswer) result.Result
}

type MiddleControllers struct {
	logger *logger.ILogger
	// ref to storage
	// ref to mqmd
	// also logger
	// config
	// answer service
	// also check auth maybe there for now

	// answerService serviIAnswerService
	config *config.Config
}

func (self MiddleControllers) SaveAnswer() result.Result {

}

// type Controllers interface {
// 	GetExercise(id int) result.Result
// 	GetListExercise(limit, offset int) result.Result
// 	GetTags(subject string) result.Result
// 	GetSubjects() result.Result
// 	UploadSecondRoundExercise() result.Result
// }
