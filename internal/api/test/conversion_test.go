package test

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/michaelgbenle/rateApp/cmd/server"
	"github.com/michaelgbenle/rateApp/internal/api"
	"github.com/michaelgbenle/rateApp/internal/helper"
	"github.com/michaelgbenle/rateApp/internal/middleware"
	"github.com/michaelgbenle/rateApp/internal/models"
	"github.com/michaelgbenle/rateApp/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("./../../../.env"); err != nil {
		log.Println(err.Error())
	}
	os.Exit(m.Run())
}

func TestUsdNgnHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	h := api.NewHTTPHandler(mockRepo)

	router := server.SetupRouter(h, mockRepo)
	testEmail := "kukus@yahoo.com"
	accClaim, _ := middleware.GenerateClaims(testEmail)

	secret := os.Getenv("JWT_SECRET")
	acc, err := middleware.GenerateToken(jwt.SigningMethodHS256, accClaim, &secret)
	if err != nil {
		t.Fail()
	}
	exchange := &models.Exchange{
		Currency: "NGN",
		Amount:   50000,
	}
	bodyJSON, err := json.Marshal(exchange)
	if err != nil {
		t.Fail()
	}
	transaction := &models.Transaction{
		UserEmail: testEmail,
		Balance: models.Balance{
			NGN: 0,
			USD: 65.96,
		},
		TransactionType: "NGN to USD",
		Success:         true,
		CreatedAt:       time.Now(),
	}
	testUser := &models.User{
		Email:    testEmail,
		Password: "lalala4$",
		Balance: models.Balance{
			NGN: 50000,
			USD: 0,
		},
	}
	rates, _ := helper.GetRates()
	rate := rates.Data.Rates.Usdcngn0.Rate
	value := helper.ConvertNgnToUsd(exchange.Amount, rate)

	mockRepo.EXPECT().FindUserByEmail(testEmail).Return(testUser, nil)
	mockRepo.EXPECT().UpdateUserbalances(testUser, exchange, value).Return(transaction, nil)

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/user/usdngn", strings.NewReader(string(bodyJSON)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *acc))
	router.ServeHTTP(rw, req)
	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Contains(t, rw.Body.String(), "success")

}

func TestNgnUsdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	h := api.NewHTTPHandler(mockRepo)

	router := server.SetupRouter(h, mockRepo)
	testEmail := "kukus@yahoo.com"
	accClaim, _ := middleware.GenerateClaims(testEmail)

	secret := os.Getenv("JWT_SECRET")
	acc, err := middleware.GenerateToken(jwt.SigningMethodHS256, accClaim, &secret)
	if err != nil {
		t.Fail()
	}
	exchange := &models.Exchange{
		Currency: "USD",
		Amount:   50,
	}
	bodyJSON, err := json.Marshal(exchange)
	if err != nil {
		t.Fail()
	}
	transaction := &models.Transaction{
		UserEmail: testEmail,
		Balance: models.Balance{
			NGN: 37900,
			USD: 50,
		},
		TransactionType: "USD to NGN",
		Success:         true,
		CreatedAt:       time.Now(),
	}
	testUser := &models.User{
		Email:    testEmail,
		Password: "lalala4$",
		Balance: models.Balance{
			NGN: 0,
			USD: 100,
		},
	}
	rates, _ := helper.GetRates()
	rate := rates.Data.Rates.Usdcngn.Rate
	value := helper.ConvertUsdToNgn(exchange.Amount, rate)

	mockRepo.EXPECT().FindUserByEmail(testEmail).Return(testUser, nil)
	mockRepo.EXPECT().UpdateUserbalances(testUser, exchange, value).Return(transaction, nil)

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/user/ngnusd", strings.NewReader(string(bodyJSON)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *acc))
	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Contains(t, rw.Body.String(), "success")
}
