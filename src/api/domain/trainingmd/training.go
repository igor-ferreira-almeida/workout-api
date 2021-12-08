package trainingmd

import (
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/presentation/trainingdto/request"
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/setmd"
	"time"
)

type Training struct {
	MuscleGroups []string              `db:"muscle_groups"`
	Sets         []setmd.Set           `db:"sets"`
	Date         time.Time             `db:"date"`
}

func NewTraining(request request.TrainingRequest) Training {
	return Training{}
}
