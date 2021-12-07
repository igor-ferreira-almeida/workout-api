package exercisemd

import (
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/presentation/exercisedto/request"
)

type Exercise struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	MuscleGroup string `db:"muscle_group"`
	Reps        int    `db:"reps"`
	Weight      int    `db:"weight"`
}

func NewExercise(request request.ExerciseRequest) Exercise {
	return Exercise{
		Name:        request.Name,
		MuscleGroup: request.MuscleGroup,
		Reps:        request.Reps,
		Weight:      request.Weight,
	}
}
