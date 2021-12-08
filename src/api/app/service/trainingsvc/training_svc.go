package trainingsvc

import (
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/trainingmd"
	"github.com/igor-ferreira-almeida/workout-api/src/api/infrastructure/repository/db/trainingdb"
)

var service Service

type Service interface {
	Find(id int64) (trainingmd.Training, error)
	Add(training trainingmd.Training) (trainingmd.Training, error)
	Update(id int64, training trainingmd.Training) (trainingmd.Training, error)
}

func init() {
	service = newServiceImpl(trainingdb.Inject())
}

func Inject() Service {
	return service
}
