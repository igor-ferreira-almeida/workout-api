package trainingsvc

import (
	"fmt"
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/trainingmd"
	"github.com/igor-ferreira-almeida/workout-api/src/api/infrastructure/repository/db/trainingdb"
)

type serviceImpl struct {
	serviceDB trainingdb.Service
}

func newServiceImpl(serviceDB trainingdb.Service) Service {
	return serviceImpl{
		serviceDB: serviceDB,
	}
}

func (service serviceImpl) Find(id int64) (trainingmd.Training, error) {
	exercise, err := service.serviceDB.Find(id)

	if err != nil {
		fmt.Println("error")
	}

	return exercise, nil
}
func (service serviceImpl) Add(training trainingmd.Training) (trainingmd.Training, error) {
	//training, err := service.serviceDB.Add(training)

	//if err != nil {
	//	fmt.Println("error")
	//}

	return training, nil
}
func (service serviceImpl) Update(id int64, training trainingmd.Training) (trainingmd.Training, error) {
	return trainingmd.Training{}, nil
}