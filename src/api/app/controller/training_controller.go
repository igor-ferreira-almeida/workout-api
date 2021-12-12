package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/presentation/trainingdto/request"
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/presentation/trainingdto/response"
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/service/trainingsvc"
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/trainingmd"
	"net/http"
)

var trainingService = trainingsvc.Inject()

func MapTrainingRoutes(router *gin.Engine) {
	trainingController := TrainingController{}
	router.POST("/training", trainingController.Add)
}

type TrainingController struct{}

// @Tags training
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Summary Create a new training
// @Param training body request.TrainingRequest true "Add an training"
// @Success 201 {object} response.TrainingResponse
// @Router /training [post]
func (controller TrainingController) Add(context *gin.Context) {
	trainingRequest := request.TrainingRequest{}
	err := context.BindJSON(&trainingRequest)

	if err != nil {

	}

	// TODO: Tratar erro posteriormente
	training, err := trainingmd.NewTraining(trainingRequest)
	if err != nil {

	}

	training, err = trainingService.Add(training)
	if err != nil {

	}

	trainingResponse := response.NewTrainingResponse(training)
	context.JSON(http.StatusCreated, trainingResponse)
}
