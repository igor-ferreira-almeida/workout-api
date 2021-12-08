package exercisesvc

import (
	"fmt"
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/exercisemd"
	"github.com/igor-ferreira-almeida/workout-api/src/api/infrastructure/repository/db/exercisedb"
)

type serviceImpl struct {
	serviceDB exercisedb.Service
}

func newServiceImpl(serviceDB exercisedb.Service) Service {
	return serviceImpl{
		serviceDB: serviceDB,
	}
}

func (service serviceImpl) Find(id int64) (exercisemd.Exercise, error) {
	exercise, err := service.serviceDB.Find(id)

	if err != nil {
		fmt.Println("error")
	}

	return exercise, nil
}
func (service serviceImpl) Add(exercise exercisemd.Exercise) (exercisemd.Exercise, error) {
	exercise, err := service.serviceDB.Add(exercise)

	if err != nil {
		fmt.Println("error")
	}

	return exercise, nil
}
func (service serviceImpl) Update(id int64, exercise exercisemd.Exercise) (exercisemd.Exercise, error) {
	return exercisemd.Exercise{}, nil
}