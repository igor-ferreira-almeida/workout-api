package musclegroupmd

const (
	chest      = "chest"
	back       = "back"
	deltoids   = "deltoids"
	biceps     = "biceps"
	triceps    = "triceps"
	forearms   = "forearms"
	quadriceps = "quadriceps"
	hamstrings = "hamstrings"
	calves     = "calves"
)

var (
	MuscleGroups = map[string]MuscleGroup{
		chest:      {name: chest},
		back:       {name: back},
		deltoids:   {name: deltoids},
		biceps:     {name: biceps},
		triceps:    {name: triceps},
		forearms:   {name: forearms},
		quadriceps: {name: quadriceps},
		hamstrings: {name: hamstrings},
		calves:     {name: calves},
	}
)

type MuscleGroup struct {
	name string
}

func IsMuscleGroupValid(muscleGroup string) bool {
	_, exist := MuscleGroups[muscleGroup]
	return exist
}
