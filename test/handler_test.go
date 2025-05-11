package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/Udehlee/healthHub-System/internals/api/handlers"
	"github.com/Udehlee/healthHub-System/internals/models"
	"github.com/Udehlee/healthHub-System/utility"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Save(user *models.User) error {
	args := m.Called(*user)
	return args.Error(0)
}

func (m *MockDB) CheckEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if user, ok := args.Get(0).(models.User); ok {
		return &user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDB) GetAllUsers() ([]*models.User, error) {
	args := m.Called()
	users := args.Get(0).([]*models.User)
	return users, args.Error(1)
}

func (m *MockDB) SaveAppointment(appointment *models.Appointment) error {
	args := m.Called(*appointment)
	return args.Error(0)
}

func (m *MockDB) AssignStaff(appointmentID int64, appointment *models.Appointment) error {
	args := m.Called(appointmentID, *appointment)
	return args.Error(0)
}

func (m *MockDB) GetAssignedAppointments() ([]*models.Appointment, error) {
	args := m.Called()
	appointments := args.Get(0).([]*models.Appointment)
	return appointments, args.Error(1)
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(m *MockDB)
		expectedCode   int
		expectedOutput string
	}{
		{
			name: "successful register",
			requestBody: models.User{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "pita@email.com",
				Password:  "password123",
			},
			mockSetup: func(m *MockDB) {
				m.On("Save", mock.AnythingOfType("models.User")).Return(nil)
			},
			expectedCode:   http.StatusOK,
			expectedOutput: "user created successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDb := new(MockDB)
			if tt.mockSetup != nil {
				tt.mockSetup(mockDb)
			}

			handler := api.NewHandler(mockDb)

			router := gin.Default()
			router.POST("/register", handler.Register)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expectedOutput)

			mockDb.AssertExpectations(t)
		})
	}
}

func TestLogin(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(m *MockDB)
		expectedCode   int
		expectedOutput string
	}{
		{
			name: "successful login",
			requestBody: models.LoginRequest{
				Email:    "pita@email.com",
				Password: "password",
			},
			mockSetup: func(m *MockDB) {
				hashedPwd, _ := utility.HashPassword("password123")
				mockUser := models.User{
					UserID:    5,
					FirstName: "pita",
					LastName:  "Doe",
					Email:     "pita@email.com",
					Password:  hashedPwd,
					Role:      "user",
				}
				m.On("CheckEmail", "pita@email.com").Return(mockUser, nil)
			},
			expectedCode:   http.StatusOK,
			expectedOutput: "logged in successfully",
		},
		{
			name:        "user not found",
			requestBody: models.LoginRequest{Email: "notfoundlee@example.com", Password: "pass"},
			mockSetup: func(m *MockDB) {
				m.On("CheckEmail", "notfoundlee@example.com").Return(models.User{}, errors.New("user not found"))
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: "user not found",
		},
		{
			name:        "wrong password",
			requestBody: models.LoginRequest{Email: "pita@email.com", Password: "wrongpass"},
			mockSetup: func(m *MockDB) {
				hashedPwd, _ := utility.HashPassword("correctpass")
				mockUser := models.User{Email: "pita@email.com", Password: hashedPwd}
				m.On("CheckEmail", "pita@email.com").Return(mockUser, nil)
			},
			expectedCode:   http.StatusBadRequest,
			expectedOutput: "wrong password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDb := new(MockDB)
			if tt.mockSetup != nil {
				tt.mockSetup(mockDb)
			}

			handler := api.NewHandler(mockDb)
			router := gin.Default()
			router.POST("/login", handler.Login)

			body, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expectedOutput)

			mockDb.AssertExpectations(t)
		})
	}
}
