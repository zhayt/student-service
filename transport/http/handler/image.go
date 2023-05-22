package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zhayt/student-service/model"
	"net/http"
	"strconv"
)

func (h *Handler) CreateImage(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), _defaultContextTime)
	defer cancel()

	var image model.Image

	if err := e.Bind(&image); err != nil {
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Bind error"})
	}

	imageID, err := h.service.Image.CreateImage(ctx, image)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorResponse{Message: "CreateImage error"})
	}

	return e.JSON(http.StatusOK, successResponse{ID: imageID})
}

func (h *Handler) UpdateImage(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), _defaultContextTime)
	defer cancel()

	var image model.Image

	if err := e.Bind(&image); err != nil {
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Bind error"})
	}

	if err := h.service.Image.UpdateImage(ctx, image); err != nil {
		return e.JSON(http.StatusInternalServerError, errorResponse{Message: "UpdateImage error"})
	}

	return e.NoContent(http.StatusOK)
}

func (h *Handler) DeleteImage(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), _defaultContextTime)
	defer cancel()

	studentID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Param error"})
	}

	if err = h.service.Image.DeleteImage(ctx, studentID); err != nil {
		return e.JSON(http.StatusInternalServerError, errorResponse{Message: "DeleteImage error"})
	}

	return e.NoContent(http.StatusOK)
}
