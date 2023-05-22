package handler

import (
	"github.com/zhayt/student-service/service"
	"go.uber.org/zap"
	"time"
)

type Handler struct {
	service *service.Service
	l       *zap.Logger
}

func NewHandler(service *service.Service, l *zap.Logger) *Handler {
	return &Handler{service: service, l: l}
}

const _defaultContextTime = 5 * time.Second

type errorResponse struct {
	Message string `json:"message"`
}

type successResponse struct {
	ID int `json:"id"`
}
