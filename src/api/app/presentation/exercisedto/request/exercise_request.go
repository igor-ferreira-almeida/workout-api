package request

type ExerciseRequest struct {
	Name        string `json:"name" example:"Bench Press"`
	MuscleGroup string `json:"muscle_group" example:"Chest"`
	Reps        int    `json:"reps" example:"12"`
	Rest        uint   `json:"rest" example:"60"`
	Weight      int    `json:"weight" example:"20"`
}
