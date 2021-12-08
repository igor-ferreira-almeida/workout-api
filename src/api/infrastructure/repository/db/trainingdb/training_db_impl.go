package trainingdb

import (
	"fmt"
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/trainingmd"
	"github.com/igor-ferreira-almeida/workout-api/src/api/infrastructure/repository/db"
)

const (
	findByIdQuery = `SELECT * FROM training t WHERE t.id = ?`
	insertQuery   = `INSERT INTO training (name, muscle_group, reps, weight) VALUES (?, ?, ?, ?)`
)

type serviceImpl struct{}

func newServiceImpl() serviceImpl {
	return serviceImpl{}
}

func (service serviceImpl) Add(training trainingmd.Training) (trainingmd.Training, error) {
	result, err := db.Connection.Exec(insertQuery, training)

	//TODO: implementar
	if err != nil {
		fmt.Println(err)
	}

	id, err := result.LastInsertId()

	//TODO: implementar
	if err != nil {
		fmt.Println(err)
	}

	savedTraining, err := service.Find(id)

	//TODO: implementar
	if err != nil {
		fmt.Println(err)
	}

	return savedTraining, nil
}


func (service serviceImpl) Find(id int64) (trainingmd.Training, error) {
	var training trainingmd.Training
	err := db.Connection.Get(&training, findByIdQuery, id)

	//TODO: implementar
	if err != nil {
		fmt.Println(err)
	}
	return training, nil
}
