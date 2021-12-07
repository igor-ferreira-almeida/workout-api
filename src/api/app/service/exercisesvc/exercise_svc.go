package exercisesvc

import (
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/exercisemd"
	"github.com/igor-ferreira-almeida/workout-api/src/api/infrastructure/repository/db/exercisedb"
)

var service Service

type Service interface {
	Find(id int64) (exercisemd.Exercise, error)
	Add(exercise exercisemd.Exercise) (exercisemd.Exercise, error)
	Update(id int64, exercise exercisemd.Exercise) (exercisemd.Exercise, error)
}

func init() {
	service = newServiceImpl(exercisedb.Inject())
}

func Inject() Service {
	return service
}