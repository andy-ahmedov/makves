package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andy-ahmedov/makves/internal/domain"
	"github.com/gin-gonic/gin"
)

type MakvesService interface {
	GetData(ctx context.Context, ids []int64) ([]domain.Makves, error)
}

type errResponse struct {
	Message string
}

type Handler struct {
	service MakvesService
}

func NewHandler(service MakvesService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) InitGinRouter() *gin.Engine {
	router := gin.Default()

	router.Use(loggingMiddleware)

	router.GET("/get-items", h.Get)
	router.Run()

	return router
}

func getIDFromRequest(idStrings []string) ([]int64, error) {
	var ids []int64
	for _, idStr := range idStrings {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (h *Handler) Get(c *gin.Context) {
	idStrings := c.QueryArray("id")
	ids, err := getIDFromRequest(idStrings)
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Errorf("invalid id format %s", err),
		})
		return
	}

	makvesData, err := h.service.GetData(context.TODO(), ids)
	if err != nil {
		if errors.Is(err, domain.ErrDataNotFound) {
			logError("get-items", "there are no rows with this identifier", err)
			c.JSON(http.StatusBadRequest, errResponse{Message: fmt.Sprintf("StatusBadRequest: %s", err.Error())})
			return
		}
		logError("get-items", "reading data from a database", err)
		c.JSON(http.StatusInternalServerError, errResponse{Message: fmt.Sprintf("InternalServerError: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, makvesData)
}
