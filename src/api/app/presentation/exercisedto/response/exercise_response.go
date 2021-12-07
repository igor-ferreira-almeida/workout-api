package response

import "github.com/igor-ferreira-almeida/workout-api/src/api/domain/exercisemd"

type ExerciseResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	MuscleGroup string `json:"muscular_group"`
	Reps        int    `json:"reps"`
	Rest        uint   `json:"rest"`
	Weight      int    `json:"weight"`
}

func NewExerciseResponse(exercise exercisemd.Exercise) ExerciseResponse {
	return ExerciseResponse{
		ID:          exercise.ID,
		Name:        exercise.Name,
		MuscleGroup: exercise.MuscleGroup,
		Reps:        exercise.Reps,
		Weight:      exercise.Weight,
	}
}
