package setmd

import "github.com/igor-ferreira-almeida/workout-api/src/api/domain/exercisemd"

type Set struct {
	Description string                `db:"description"`
	Exercises   []exercisemd.Exercise `db:"exercises"`
}
