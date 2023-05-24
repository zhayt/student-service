package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zhayt/student-service/model"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (h *Handler) CreateOrUpdateImage(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), _defaultContextTime)
	defer cancel()

	var image model.Image

	if err := e.Bind(&image); err != nil {
		h.l.Error("Bind error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Bind error"})
	}

	if err := h.service.Image.CreateOrUpdateImage(ctx, image); err != nil {
		h.l.Error("CreateOrUpdateImage error", zap.Error(err))
		return e.JSON(http.StatusInternalServerError, errorResponse{Message: "UpdateImage error"})
	}

	h.l.Info("Image updated", zap.Int("studentID", image.StudentID))

	return e.NoContent(http.StatusOK)
}

func (h *Handler) DeleteImage(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), _defaultContextTime)
	defer cancel()

	studentID, err := strconv.Atoi(e.Param("studentID"))
	if err != nil {
		h.l.Error("Param error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Param error"})
	}

	if err := h.service.Image.DeleteImage(ctx, studentID); err != nil {
		h.l.Error("DeleteImage error", zap.Error(err))
		return e.JSON(http.StatusInternalServerError, errorResponse{Message: "DeleteImage error"})
	}

	h.l.Info("Image deleted", zap.Int("studentID", studentID))
	return e.NoContent(http.StatusOK)
}
