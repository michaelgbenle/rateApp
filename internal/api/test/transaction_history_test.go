package test

import (
	"github.com/golang/mock/gomock"
	"github.com/michaelgbenle/rateApp/cmd/server"
	"github.com/michaelgbenle/rateApp/internal/api"
	"github.com/michaelgbenle/rateApp/internal/repository/mocks"
	"testing"
)

func TestTransactionsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	h := api.NewHTTPHandler(mockRepo)

	router := server.SetupRouter(h, mockRepo)

}
