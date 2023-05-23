package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zhayt/student-service/model"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) SingIn(e echo.Context) error {
	ctx, cancel := context.WithTimeout(e.Request().Context(), _defaultContextTime)
	defer cancel()

	var studentDTO model.StudentDTO

	if err := e.Bind(&studentDTO); err != nil {
		h.l.Error("Bind error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Bind error"})
	}

	student, err := h.service.Student.GetStudentByEmail(ctx, studentDTO)
	if err != nil {
		h.l.Error("GetStudentByEmail error", zap.Error(err))
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Get student error"})
	}

	return e.JSON(http.StatusOK, student)
}

func (h *Handler) SignUp(e echo.Context) error {
	ctx, cancel := context.WithTimeout(e.Request().Context(), _defaultContextTime)
	defer cancel()

	var student model.Student

	if err := e.Bind(&student); err != nil {
		return e.JSON(http.StatusBadRequest, errorResponse{Message: "Bind error"})
	}

	studentID, err := h.service.Student.CreateStudent(ctx, student)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorResponse{Message: "Create student error"})
	}

	return e.JSON(http.StatusOK, successResponse{ID: studentID})
}
