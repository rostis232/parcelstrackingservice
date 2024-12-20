package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rostis232/parcelstrackingservice/models"
	"net/http"
)

type Service interface {
	GetInfo(trackNumber string) (*models.Data, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Track(c echo.Context) error {
	request := struct {
		TrackingNumber string `json:"tracking_number"`
	}{}

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(request.TrackingNumber) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "tracking_number is required")
	}

	data, err := h.service.GetInfo(request.TrackingNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
