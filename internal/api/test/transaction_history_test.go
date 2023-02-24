package test

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/michaelgbenle/rateApp/cmd/server"
	"github.com/michaelgbenle/rateApp/internal/api"
	"github.com/michaelgbenle/rateApp/internal/middleware"
	"github.com/michaelgbenle/rateApp/internal/models"
	"github.com/michaelgbenle/rateApp/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTransactionsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)
	h := api.NewHTTPHandler(mockRepo)

	router := server.SetupRouter(h, mockRepo)

	testEmail := "kukus@yahoo.com"

	t1 := models.Transaction{
		UserEmail: testEmail,
		Balance: models.Balance{
			NGN: 5000,
			USD: 200,
		},
		TransactionType: "NGN to USD",
		Success:         true,
		CreatedAt:       time.Now(),
	}
	t2 := models.Transaction{
		UserEmail: testEmail,
		Balance: models.Balance{
			NGN: 80000,
			USD: 100,
		},
		TransactionType: "USD to NGN",
		Success:         true,
		CreatedAt:       time.Now(),
	}
	testTransactions := &[]models.Transaction{t1, t2}
	bodyJSON, err := json.Marshal(testTransactions)
	if err != nil {
		t.Fail()
	}
	testUser := &models.User{
		Email:    testEmail,
		Password: "lalala4$",
		Balance: models.Balance{
			NGN: 0,
			USD: 100,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	secret := os.Getenv("JWT_SECRET")

	accClaim, _ := middleware.GenerateClaims(testEmail)
	acc, err := middleware.GenerateToken(jwt.SigningMethodHS256, accClaim, &secret)
	fmt.Println(acc)
	if err != nil {
		t.Fail()
	}
	mockRepo.EXPECT().FindUserByEmail(testEmail).Return(testUser, nil)
	t.Run("Testing for Successful Request", func(t *testing.T) {

		mockRepo.EXPECT().GetTransactions(testUser).Return(testTransactions, nil)

		rw := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/user/transactions_history", strings.NewReader(string(bodyJSON)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *acc))

		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "transaction history and balances")
	})
}
