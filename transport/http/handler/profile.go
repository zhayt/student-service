package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zhayt/student-service/model"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func (h *Handler) ShowProfile(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	studentName := e.Param("name")

	studentProfile := h.service.Profile.GetStudentAllProfileData(ctx, studentName)

	return e.JSON(http.StatusOK, studentProfile)
}

func (h *Handler) UpdateProfile(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), _defaultContextTime)
	defer cancel()

	var studentInfo model.StudentPersonalInfo

	if err := e.Bind(&studentInfo); err != nil {
		h.l.Error("Bind error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Bind error"})
	}

	if err := h.service.Profile.CreateOrUpdateStudentInfo(ctx, studentInfo); err != nil {
		h.l.Error("CreateOrUpdateStudentInfo error", zap.Error(err))
		return e.JSON(http.StatusInternalServerError, errorResponse{Message: "CreateOrUpdateStudentInfo error"})
	}

	h.l.Info("Student info was created or updated", zap.Int("studentID", studentInfo.StudentID))

	return e.NoContent(http.StatusOK)
}
