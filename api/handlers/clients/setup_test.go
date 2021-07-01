package clients_test

import (
	"context"

	handler "github.com/fanfit/login/api/handlers"
	"github.com/fanfit/login/api/handlers/clients"
	"github.com/fanfit/login/models/clients/repository"
	"github.com/fanfit/login/models/clients/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type mockService struct {
	GetClientResponse    repository.GetClientByEmailRow
	CreateClientResponse repository.GetClientByIDRow

	GetClientError  error
	CreateClientErr error
}

func (m mockService) GetClientByEmail(ctx context.Context, id string) (repository.GetClientByEmailRow, error) {
	return m.GetClientResponse, m.GetClientError
}

func (m mockService) CreateClient(ctx context.Context, client repository.Client) (repository.GetClientByIDRow, error) {
	return m.CreateClientResponse, m.CreateClientErr
}

func setupRouter(s service.Service) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(handler.NoRoute)
	group := r.Group("/v1")
	clients.Routes(group, s)
	return r
}
