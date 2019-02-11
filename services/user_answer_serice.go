package services

import (
	"github.com/OlympBMSTU/exercises/db/result"
	"github.com/OlympBMSTU/test-second-service/entities"
	"github.com/jackc/pgx"
)

type IUserAnswerService interface {
	SaveAnswer(ans entities.UserAnswer) result.DbResult
	//GetAnswer(userID, ) DbResult
	// /GetUserAnswers()
	// Think about
	// GetList
}

// its a good variant select t_n, id from tags where t_n in names;
// then get ids that exist and not exist create -> get new ids -> insert into mm table
type UserAnswerService struct {
	db *pgx.ConnPool
}

func SaveAnswer(ans entities.UserAnswer) result.DbResult {
	return result.OkResult(nil, nil)
}
