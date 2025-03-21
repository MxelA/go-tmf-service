package handler_v4_2

import (
	"github.com/MxelA/tmf-service-go/pkg/repository"
)

type ServiceOrderHandler struct {
	repo repository.ServiceOrderRepository
}

func NewServiceOrderHandler(repo repository.ServiceOrderRepository) *ServiceOrderHandler {
	return &ServiceOrderHandler{repo: repo}
}
