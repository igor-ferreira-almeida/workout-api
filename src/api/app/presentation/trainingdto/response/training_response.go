package response

import (
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/trainingmd"
	"time"
)

type TrainingResponse struct {
	MuscleGroups         []string           `json:"muscle_groups"`
	Date                 time.Time          `json:"date"`
	Exercises            []ExerciseResponse `json:"exercises"`
	VolumePerMuscleGroup map[string]int     `json:"volume_per_muscle_group"`
}

type ExerciseResponse struct {
	Name        string              `json:"name"`
	MuscleGroup string              `json:"muscle_group"`
	Executions  []ExecutionResponse `json:"executions"`
	Volume      int                 `json:"volume"`
}

type ExecutionResponse struct {
	Technique string        `json:"technique"`
	Sets      []SetResponse `json:"sets"`
	Volume    int           `json:"volume"`
}

type SetResponse struct {
	Reps   int `json:"reps"`
	Rest   int `json:"rest"`
	Weight int `json:"weight"`
	Volume int `json:"volume"`
}

func NewTrainingResponse(training trainingmd.Training) TrainingResponse {
	exercisesResponse := make([]ExerciseResponse, 0, len(training.Exercises))
	for _, exercise := range training.Exercises {
		exercisesResponse = append(exercisesResponse, newExerciseResponse(exercise))
	}
	return TrainingResponse{
		MuscleGroups:         training.MuscleGroups,
		Date:                 training.Date,
		Exercises:            exercisesResponse,
		VolumePerMuscleGroup: training.VolumePerMuscleGroup,
	}
}

func newExerciseResponse(exercise trainingmd.Exercise) ExerciseResponse {
	executionsResponse := make([]ExecutionResponse, 0, len(exercise.Executions))
	for _, execution := range exercise.Executions {
		executionsResponse = append(executionsResponse, newExecutionResponse(execution))
	}
	return ExerciseResponse{
		Name: exercise.Name,
		MuscleGroup: exercise.MuscleGroup,
		Volume: exercise.Volume,
		Executions: executionsResponse,
	}
}

func newExecutionResponse(execution trainingmd.Execution) ExecutionResponse {
	setsResponse := make([]SetResponse, 0, len(execution.Sets))
	for _, set := range execution.Sets {
		setsResponse = append(setsResponse, newSetResponse(set))
	}
	return ExecutionResponse{
		Technique: execution.Technique,
		Volume: execution.Volume,
		Sets: setsResponse,
	}
}

func newSetResponse(set trainingmd.Set) SetResponse {
	return SetResponse{
		Reps: set.Reps,
		Rest: set.Rest,
		Weight: set.Weight,
		Volume: set.Volume,
	}
}
