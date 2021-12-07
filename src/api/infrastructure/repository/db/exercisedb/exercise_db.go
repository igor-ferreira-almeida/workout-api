package exercisedb

import "github.com/igor-ferreira-almeida/workout-api/src/api/domain/exercisemd"

var service Service

type Service interface {
	Add(exercise exercisemd.Exercise) (exercisemd.Exercise, error)
	Find(id int64) (exercisemd.Exercise, error)
}

func init() {
	service = newServiceImpl()
}

func Inject() Service {
	return service
}