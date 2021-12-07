package exercisedb

import (
	"fmt"
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/exercisemd"
	"github.com/igor-ferreira-almeida/workout-api/src/api/infrastructure/repository/db"
)

const (
	findByIdQuery = `SELECT * FROM exercise e WHERE e.id = ?`
	insertQuery   = `INSERT INTO exercise (name, muscle_group, reps, weight) VALUES (?, ?, ?, ?)`
)

type serviceImpl struct{}

func newServiceImpl() serviceImpl {
	return serviceImpl{}
}


type Exercise struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	MuscleGroup string `db:"muscle_group"`
	Reps        int    `db:"reps"`
	Weight      int    `db:"weight"`
}

func (service serviceImpl) Add(exercise exercisemd.Exercise) (exercisemd.Exercise, error) {
	result, err := db.Connection.Exec(insertQuery, exercise.Name, exercise.MuscleGroup, exercise.Reps, exercise.Weight)

	//TODO: implementar
	if err != nil {
		fmt.Println(err)
	}

	id, err := result.LastInsertId()

	//TODO: implementar
	if err != nil {
		fmt.Println(err)
	}

	savedFood, err := service.Find(id)

	//TODO: implementar
	if err != nil {
		fmt.Println(err)
	}

	return savedFood, nil
}


func (service serviceImpl) Find(id int64) (exercisemd.Exercise, error) {
	var food exercisemd.Exercise
	err := db.Connection.Get(&food, findByIdQuery, id)

	//TODO: implementar
	if err != nil {
		fmt.Println(err)
	}
	return food, nil
}
