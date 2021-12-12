package trainingmd

import (
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/presentation/trainingdto/request"
	"time"
)

type Training struct {
	MuscleGroups         []string       `json:"muscle_groups"`
	Date                 time.Time      `json:"date"`
	Exercises            []Exercise     `json:"exercises"`
	Volume               int            `json:"volume"`
	VolumePerMuscleGroup map[string]int `json:"volume_per_muscle_group"`
}

type Exercise struct {
	Name        string      `json:"name"`
	MuscleGroup string      `json:"muscle_group"`
	Executions  []Execution `json:"executions"`
	Volume      int         `json:"volume"`
}

type Execution struct {
	Technique string `json:"technique"`
	Sets      []Set  `json:"sets"`
	Volume    int    `json:"volume"`
}

type Set struct {
	Reps   int `json:"reps"`
	Rest   int `json:"rest"`
	Weight int `json:"weight"`
	Volume int `json:"volume"`
}

func NewTraining(request request.TrainingRequest) (Training, error) {
	date, err := time.Parse(time.RFC3339, request.Date)

	exercises := make([]Exercise, 0, len(request.Exercises))
	var totalVolume int
	volumePerMuscleGroup := make(map[string]int)

	for _, exerciseRequest := range request.Exercises {
		exercise := newExercise(exerciseRequest)
		volumePerMuscleGroup[exercise.MuscleGroup] += exercise.Volume
		totalVolume += exercise.Volume
		volumePerMuscleGroup["Total"] += exercise.Volume
		exercises = append(exercises, exercise)
	}

	if err != nil {
		return Training{}, err
	}
	return Training{
		MuscleGroups: request.MuscleGroups,
		Date:         date,
		Exercises:    exercises,
		Volume:       totalVolume,
		VolumePerMuscleGroup: volumePerMuscleGroup,
	}, nil
}

func newExercise(request request.ExerciseRequest) Exercise {
	techniques := make([]Execution, 0, len(request.Executions))
	var volume int
	for _, requestTechnique := range request.Executions {
		technique := newTechnique(requestTechnique)
		volume += technique.Volume
		techniques = append(techniques, technique)
	}
	return Exercise{
		Name:        request.Name,
		MuscleGroup: request.MuscleGroup,
		Executions:  techniques,
		Volume:      volume,
	}
}

func newTechnique(request request.ExecutionRequest) Execution {
	sets := make([]Set, 0, len(request.Sets))
	var volume int

	for _, setRequest := range request.Sets {
		set := newSet(setRequest)
		volume += set.Volume
		sets = append(sets, set)
	}
	return Execution{
		Technique: request.Technique,
		Volume:    volume,
		Sets:      sets,
	}
}

func newSet(request request.SetRequest) Set {
	return Set{
		Reps:   request.Reps,
		Rest:   request.Rest,
		Weight: request.Weight,
		Volume: request.Reps * request.Weight,
	}
}
