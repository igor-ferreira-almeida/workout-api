package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/presentation/exercisedto/request"
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/presentation/exercisedto/response"
	"github.com/igor-ferreira-almeida/workout-api/src/api/app/service/exercisesvc"
	"github.com/igor-ferreira-almeida/workout-api/src/api/domain/exercisemd"
	"net/http"
	"strconv"
)

var exerciseService = exercisesvc.Inject()

func MapExerciseRoutes(router *gin.Engine) {
	exerciseController := ExerciseController{}
	router.GET("/exercises/:id", exerciseController.Find)
	router.POST("/exercises", exerciseController.Add)
}

type ExerciseController struct {}

// @Tags exercise
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Summary Find exercise by id
// @Param id path string true "Exercise ID"
// @Success 200 {object} response.ExerciseResponse
// @Router /exercises/{id} [get]
func (controller ExerciseController) Find(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	fmt.Println(id)
	fmt.Println(err)

	//if err != nil {
	//	errorResponseDTO := errordto.NewErrorResponseDTO(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, "Invalid param for id")
	//	context.JSON(http.StatusBadRequest, errorResponseDTO)
	//	return
	//}
	//
	//foodFound, err := foodService.Find(id)
	//
	//if err != nil {
	//	errorResponseDTO := dto.NewErrorResponseDTO(http.StatusText(http.StatusNotFound), http.StatusNotFound, err.Error())
	//	context.JSON(http.StatusNotFound, errorResponseDTO)
	//	return
	//}
	//
	//foodResponse := foodresponse.NewFoodResponse(foodFound)
	context.JSON(http.StatusOK, id)
}

// @Tags exercise
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Summary Create a new exercise tests
// @Param exercise body request.ExerciseRequest true "Add an exercise"
// @Success 201 {object} response.ExerciseResponse
// @Router /exercises [post]
func (controller ExerciseController) Add(context *gin.Context) {
	exerciseRequest := request.ExerciseRequest{}
	err := context.BindJSON(&exerciseRequest)

	if err != nil {

	}

	// TODO: Tratar erro posteriormente
	exercise, _ := exerciseService.Add(exercisemd.NewExercise(exerciseRequest))

	exerciseResponse := response.NewExerciseResponse(exercise)
	context.JSON(http.StatusCreated, exerciseResponse)
}