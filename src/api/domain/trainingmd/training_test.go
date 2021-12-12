package trainingmd

import (
	"encoding/json"
	"fmt"
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/presentation/trainingdto/request"
	"testing"
)

func TestNewTraining(t *testing.T) {
	trainingRequest := request.TrainingRequest{}

	trainingRequest.MuscleGroups = []string{"Hamstrings", "Glutes", "Quadriceps", "Calves"}
	trainingRequest.Date = "2021-12-07T00:00:00-03:00"
	trainingRequest.Exercises = []request.ExerciseRequest{
		{
			Name: "Lying Leg Curls",
			MuscleGroup: "Hamstrings",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 30, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Drop Set",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 40},
						{Reps: 10, Weight: 30},
						{Reps: 10, Weight: 20},
					},
				},
				{
					Technique: "Drop Set",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 50},
						{Reps: 10, Weight: 40},
						{Reps: 10, Weight: 30},
					},
				},
				{
					Technique: "Drop Set",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 50},
						{Reps: 9, Weight: 40},
						{Reps: 9, Weight: 30},
					},
				},
				{
					Technique: "Drop Set",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 50},
						{Reps: 8, Weight: 40},
						{Reps: 9, Weight: 30},
					},
				},
			},
		},
		{
			Name: "Romanian Deadlift",
			MuscleGroup: "Hamstrings",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 12, Rest: 60, Weight: 40},
					},
				},
			},
		},
		{
			Name: "Seated Leg Curls",
			MuscleGroup: "Hamstrings",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 18, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 13, Rest: 60, Weight: 25},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 15, Rest: 60, Weight: 25},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 30},
					},
				},
			},
		},
		{
			Name: "Twisting Hip Extension",
			MuscleGroup: "Glutes",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 12},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 13, Rest: 60, Weight: 12},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 13, Rest: 60, Weight: 20},
					},
				},
			},
		},
		{
			Name: "Sumo Deadlift",
			MuscleGroup: "Glutes",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 12, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 50},
					},
				},
			},
		},
		{
			Name: "Bulgarian Split Squat",
			MuscleGroup: "Glutes",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 1},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 12, Rest: 60, Weight: 10},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 10},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 10},
					},
				},
			},
		},
		{
			Name: "Hip Adduction Machine",
			MuscleGroup: "Glutes",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 25, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 19, Rest: 60, Weight: 30},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 15, Rest: 60, Weight: 30},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 40},
					},
				},
			},
		},
		{
			Name: "Standing Calf Press (Smith Machine)",
			MuscleGroup: "Calves",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 40},
						{Reps: 10, Weight: 1},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 40},
						{Reps: 11, Weight: 1},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 40},
						{Reps: 10, Weight: 1},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 40},
						{Reps: 10, Weight: 1},
					},
				},
			},
		},
		{
			Name: "Seated Donkey Calf Raise",
			MuscleGroup: "Calves",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 18, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 60},
					},
				},
			},
		},
		{
			Name: "Seated Calf Raise",
			MuscleGroup: "Calves",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 12, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 14, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 25},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 25},
					},
				},
			},
		},
	}
	jsonRequest, _ := json.Marshal(trainingRequest)
	fmt.Println(jsonRequest)

	training, err := NewTraining(trainingRequest)

	if err != nil {
		t.Fatal(err)
	}

	for _, exercise := range training.Exercises {
		// 400 + 450 + 500 + 400 + 600
		if exercise.Volume != 2350 {
			t.Fatal(err)
		}
	}

	fmt.Println(training)
}


func TestNewTraining_LegDay(t *testing.T) {
	trainingRequest := request.TrainingRequest{}

	trainingRequest.MuscleGroups = []string{"Hamstrings", "Glutes", "Quadriceps", "Calves"}
	trainingRequest.Date = "2021-12-07T00:00:00-03:00"
	trainingRequest.Exercises = []request.ExerciseRequest{
		{
			Name: "Lying Leg Curls",
			MuscleGroup: "Hamstrings",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 30, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Drop Set",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 40},
						{Reps: 10, Weight: 30},
						{Reps: 10, Weight: 20},
					},
				},
				{
					Technique: "Drop Set",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 50},
						{Reps: 10, Weight: 40},
						{Reps: 10, Weight: 30},
					},
				},
				{
					Technique: "Drop Set",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 50},
						{Reps: 9, Weight: 40},
						{Reps: 9, Weight: 30},
					},
				},
				{
					Technique: "Drop Set",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 50},
						{Reps: 8, Weight: 40},
						{Reps: 9, Weight: 30},
					},
				},
			},
		},
		{
			Name: "Romanian Deadlift",
			MuscleGroup: "Hamstrings",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 12, Rest: 60, Weight: 40},
					},
				},
			},
		},
		{
			Name: "Seated Leg Curls",
			MuscleGroup: "Hamstrings",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 18, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 13, Rest: 60, Weight: 25},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 15, Rest: 60, Weight: 25},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 30},
					},
				},
			},
		},
		{
			Name: "Twisting Hip Extension",
			MuscleGroup: "Glutes",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 12},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 13, Rest: 60, Weight: 12},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 13, Rest: 60, Weight: 20},
					},
				},
			},
		},
		{
			Name: "Sumo Deadlift",
			MuscleGroup: "Glutes",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 12, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 50},
					},
				},
			},
		},
		{
			Name: "Bulgarian Split Squat",
			MuscleGroup: "Glutes",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 1},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 12, Rest: 60, Weight: 10},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 10},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 10},
					},
				},
			},
		},
		{
			Name: "Hip Adduction Machine",
			MuscleGroup: "Glutes",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 25, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 19, Rest: 60, Weight: 30},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 15, Rest: 60, Weight: 30},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 40},
					},
				},
			},
		},
		{
			Name: "Standing Calf Press (Smith Machine)",
			MuscleGroup: "Calves",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 40},
						{Reps: 10, Weight: 1},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 40},
						{Reps: 11, Weight: 1},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 40},
						{Reps: 10, Weight: 1},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 40},
						{Reps: 10, Weight: 1},
					},
				},
			},
		},
		{
			Name: "Seated Donkey Calf Raise",
			MuscleGroup: "Calves",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 18, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 60},
					},
				},
			},
		},
		{
			Name: "Seated Calf Raise",
			MuscleGroup: "Calves",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 12, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 14, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 25},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 25},
					},
				},
			},
		},
	}

	training, err := NewTraining(trainingRequest)

	if err != nil {
		t.Fatal(err)
	}

	for _, exercise := range training.Exercises {
		// 400 + 450 + 500 + 400 + 600
		if exercise.Volume != 2350 {
			t.Fatal(err)
		}
	}

	fmt.Println(training)
}

func TestNewTraining_ChestDay(t *testing.T) {
	trainingRequest := request.TrainingRequest{}

	trainingRequest.MuscleGroups = []string{"Chest", "Shoulders"}
	trainingRequest.Date = "2021-12-08T00:00:00-03:00"
	trainingRequest.Exercises = []request.ExerciseRequest{
		{
			Name: "Incline Bench Press",
			MuscleGroup: "Chest",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 16, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 16, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 60},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 60},
					},
				},
				{
					Technique: "Drop Set",
					Sets: []request.SetRequest{
						{Reps: 1, Weight: 70},
						{Reps: 1, Weight: 80},
						{Reps: 7, Weight: 60},
						{Reps: 12, Weight: 40},
					},
				},
			},
		},
		{
			Name: "Incline Dumbbell Fly",
			MuscleGroup: "Chest",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 6, Rest: 60, Weight: 64},
						{Reps: 6, Weight: 44},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 5, Rest: 60, Weight: 64},
						{Reps: 5, Weight: 44},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 4, Rest: 60, Weight: 64},
						{Reps: 5, Weight: 44},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 44},
					},
				},
			},
		},
		{
			Name: "Incline Cable fly",
			MuscleGroup: "Chest",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 14, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 12, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 13, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 60},
					},
				},
			},
		},
		{
			Name: "Fly Machine",
			MuscleGroup: "Chest",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 17, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 15, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 60},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 60},
					},
				},
			},
		},
		{
			Name: "Cable Fly",
			MuscleGroup: "Chest",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 6, Rest: 60, Weight: 40},
						{Reps: 12, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 6, Rest: 60, Weight: 40},
						{Reps: 15, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 6, Rest: 60, Weight: 40},
						{Reps: 12, Rest: 60, Weight: 40},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 6, Rest: 60, Weight: 40},
						{Reps: 10, Rest: 60, Weight: 40},
					},
				},
			},
		},
		{
			Name: "Military Press",
			MuscleGroup: "Shoulders",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 15, Rest: 60, Weight: 50},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 15, Rest: 60, Weight: 20},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 26},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 26},
					},
				},
			},
		},
		{
			Name: "Machine Press",
			MuscleGroup: "Shoulders",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 11, Rest: 60, Weight: 25},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 25},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 30},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 30},
					},
				},
			},
		},
		{
			Name: "Dumbbell Lateral Raise",
			MuscleGroup: "Shoulders",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 24},
						{Reps: 10, Weight: 12},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 24},
						{Reps: 10, Weight: 12},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 10, Rest: 60, Weight: 24},
						{Reps: 10, Weight: 12},
					},
				},
				{
					Technique: "Bi Set",
					Sets: []request.SetRequest{
						{Reps: 9, Rest: 60, Weight: 24},
						{Reps: 8, Weight: 12},
					},
				},
			},
		},
		{
			Name: "One-Arm Cable Lateral Raise",
			MuscleGroup: "Shoulders",
			Executions: []request.ExecutionRequest{
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 10},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 10},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 8, Rest: 60, Weight: 10},
					},
				},
				{
					Technique: "Normal",
					Sets: []request.SetRequest{
						{Reps: 7, Rest: 60, Weight: 10},
					},
				},
			},
		},
	}

	jsonRequest, _ := json.Marshal(trainingRequest)
	fmt.Println(jsonRequest)
	training, err := NewTraining(trainingRequest)

	if err != nil {
		t.Fatal(err)
	}

	for _, exercise := range training.Exercises {
		// 400 + 450 + 500 + 400 + 600
		if exercise.Volume != 2350 {
			t.Fatal(err)
		}
	}

	fmt.Println(training)
}