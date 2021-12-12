package request

type TrainingRequest struct {
	MuscleGroups []string          `json:"muscle_groups" example:"Chest,Back,Deltoids,Biceps,Triceps,Forearms,Quadriceps,Hamstrings,Calves"`
	Exercises    []ExerciseRequest `json:"exercises"`
	Date         string            `json:"date" example:"2021-12-07T00:00:00-03:00"`
}

type ExerciseRequest struct {
	Name        string             `json:"name" example:"Bench Press"`
	MuscleGroup string             `json:"muscle_group" example:"Chest"`
	Executions  []ExecutionRequest `json:"executions"`
}

type ExecutionRequest struct {
	Technique string       `json:"technique" example:"Normal,Drop Set,Bi Set,Rest Pause"`
	Sets      []SetRequest `json:"sets"`
}

type SetRequest struct {
	Reps   int `json:"reps" example:"12"`
	Rest   int `json:"rest" example:"30"`
	Weight int `json:"weight" example:"20"`
}
