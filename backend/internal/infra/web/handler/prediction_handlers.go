package handler

import (
	"context"
	"net/http"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PredictionHandlers struct {
	EventDispatcher        events.EventDispatcherInterface
	PredictionRepository   entity.PredictionRepository
	PredictionCreatedEvent events.EventInterface
}

func NewPredictionHandlers(
	eventDispatcher events.EventDispatcherInterface,
	predictionRepository entity.PredictionRepository,
	predictionCreatedEvent events.EventInterface,
) *PredictionHandlers {
	return &PredictionHandlers{
		EventDispatcher:        eventDispatcher,
		PredictionRepository:   predictionRepository,
		PredictionCreatedEvent: predictionCreatedEvent,
	}
}

// CreatePredictionHandler
// @Summary Create a new Prediction
// @Description Create a new Prediction in the system
// @Tags Predictions
// @Accept json
// @Produce json
// @Param input body prediction_usecase.CreatePredictionInputDTO true "Prediction entity to create"
// @Success 200 {object} prediction_usecase.CreatePredictionOutputDTO
// @Router /prediction [post]
func (h *PredictionHandlers) CreatePredictionHandler(c *gin.Context) {
	var input prediction_usecase.CreatePredictionInputDTO
	ctx := context.Background()
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := prediction_usecase.NewCreatePredictionUseCase(
		h.PredictionCreatedEvent,
		h.PredictionRepository,
		h.EventDispatcher,
	).Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// FindPredictionByIdHandler
// @Summary Retrieve a Prediction by ID
// @Description Get details of a specific Prediction by its ID
// @Tags Predictions
// @Accept json
// @Produce json
// @Param prediction_id path string true "Prediction ID"
// @Success 200 {object} prediction_usecase.FindPredictionOutputDTO
// @Router /prediction/{prediction_id} [get]
func (h *PredictionHandlers) FindPredictionByIdHandler(c *gin.Context) {
	var input prediction_usecase.FindPredictionByIdInputDTO
	predictionId, err := uuid.Parse(c.Param("prediction_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.PredictionId = predictionId

	ctx := context.Background()
	ctx = context.WithValue(ctx, "limit", c.DefaultQuery("limit", "20"))
	ctx = context.WithValue(ctx, "offset", c.DefaultQuery("offset", "0"))
	res, err := prediction_usecase.NewFindPredictionByIdUseCase(h.PredictionRepository).Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// FindAllPredictionsByLocationIdHandler
// @Summary Retrieve all Predictions by Location ID
// @Description Get a list of all Predictions by Location ID
// @Tags Predictions
// @Accept json
// @Produce json
// @Param location_id path string true "Location ID"
// @Success 200 {array} prediction_usecase.FindAllPredictionsByLocationIdOutputDTO
// @Router /prediction/location/{location_id} [get]
func (h *PredictionHandlers) FindAllPredictionsByLocationIdHandler(c *gin.Context) {
	var input prediction_usecase.FindAllPredictionsByLocationIdInputDTO
	locationId, err := uuid.Parse(c.Param("location_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.LocationId = locationId
	ctx := context.Background()
	ctx = context.WithValue(ctx, "limit", c.DefaultQuery("limit", "20"))
	ctx = context.WithValue(ctx, "offset", c.DefaultQuery("offset", "0"))
	res, err := prediction_usecase.NewFindAllPredictionsByLocationIdUseCase(h.PredictionRepository).Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// FindAllPredictionsHandler
// @Summary Retrieve all Predictions
// @Description Get a list of all Predictions
// @Tags Predictions
// @Accept json
// @Produce json
// @Success 200 {array} prediction_usecase.FindAllPredictionsOutputDTO
// @Router /prediction [get]
func (h *PredictionHandlers) FindAllPredictionsHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "limit", c.DefaultQuery("limit", "20"))
	ctx = context.WithValue(ctx, "offset", c.DefaultQuery("offset", "0"))
	res, err := prediction_usecase.NewFindAllPredictionsUseCase(h.PredictionRepository).Execute(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdatePredictionHandler
// @Summary Update a Prediction
// @Description Update a specific Prediction entity
// @Tags Predictions
// @Accept json
// @Produce json
// @Param prediction_id path string true "Prediction ID"
// @Param input body prediction_usecase.UpdatePredictionInputDTO true "Prediction entity to update"
// @Success 200 {object} prediction_usecase.UpdatePredictionOutputDTO
// @Router /prediction/{prediction_id} [put]
func (h *PredictionHandlers) UpdatePredictionHandler(c *gin.Context) {
	var input prediction_usecase.UpdatePredictionInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	predictionId, err := uuid.Parse(c.Param("prediction_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.PredictionId = predictionId
	ctx := context.Background()
	res, err := prediction_usecase.NewUpdatePredictionUseCase(h.PredictionRepository).Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeletePredictionHandler
// @Summary Delete a Prediction
// @Description Remove a specific Prediction from the system
// @Tags Predictions
// @Accept json
// @Produce json
// @Param prediction_id path string true "Prediction ID"
// @Success 200 {string} string "Prediction deleted successfully"
// @Router /prediction/{prediction_id} [delete]
func (h *PredictionHandlers) DeletePredictionHandler(c *gin.Context) {
	var input prediction_usecase.DeletePredictionInputDTO
	predictionId, err := uuid.Parse(c.Param("prediction_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.PredictionId = predictionId
	ctx := context.Background()
	err = prediction_usecase.NewDeletePredictionUseCase(h.PredictionRepository).Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prediction deleted successfully"})
}
