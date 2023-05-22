package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zhayt/student-service/model"
	"net/http"
)

func (h *Handler) ShowProfile(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), _defaultContextTime)
	defer cancel()

	studentName := e.Param("name")

	studentProfile, err := h.service.PersonalInfo.GetStudentAllProfileData(ctx, studentName)
	if err != nil {
		return e.JSON(http.StatusNotFound, errorResponse{Message: "GetStudentAllProfileData error"})
	}

	return e.JSON(http.StatusOK, studentProfile)
}

func (h *Handler) UpdateProfile(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), _defaultContextTime)
	defer cancel()

	var studentInfo model.StudentPersonalInfo

	if err := e.Bind(&studentInfo); err != nil {
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Bind error"})
	}

	if err := h.service.PersonalInfo.CreateOrUpdateStudentInfo(ctx, studentInfo); err != nil {
		return e.JSON(http.StatusInternalServerError, errorResponse{Message: "CreateOrUpdateStudentInfo error"})
	}

	return e.NoContent(http.StatusOK)
}
