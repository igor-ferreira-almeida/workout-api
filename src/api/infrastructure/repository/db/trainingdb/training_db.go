package trainingdb

import (
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/trainingmd"
)

var service Service

type Service interface {
	Add(training trainingmd.Training) (trainingmd.Training, error)
	Find(id int64) (trainingmd.Training, error)
}

func init() {
	service = newServiceImpl()
}

func Inject() Service {
	return service
}