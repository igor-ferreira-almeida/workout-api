package exercisedto

type ExerciseRequest struct {
	Name string `json:"name" example:"Bench Press"`
	MuscularGroup string `json:"muscular_group" example:"Chest"`
	Reps uint `json:"reps" example:"12"`
}