package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dede182/revesion/types"
	"github.com/gorilla/mux"
)

type mockedUserStore struct{}

func TestUserServiceHandlerStore(t *testing.T) {

	userStore := &mockedUserStore{}
	handler := NewHandler(userStore)

	testCases := []struct {
		name         string
		payload      types.RegisterUserPayload
		expectedCode int
	}{
		{
			name:         "should fail if the payload is invalid",
			payload:      types.RegisterUserPayload{},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "test user can correctly register",
			payload: types.RegisterUserPayload{
				FirstName: "test",
				LastName:  "last",
				Email:     "asdf@gmail.com",
				Password:  "asdffdsa",
			},
			expectedCode: http.StatusCreated,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			marshelled, _ := json.Marshal(tt.payload)

			req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshelled))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			router := mux.NewRouter()

			router.HandleFunc("/register", handler.handleRegister)
			router.ServeHTTP(rr, req)

			if rr.Code != tt.expectedCode {
				t.Errorf("expected status code %d,got %d", tt.expectedCode, rr.Code)

			}
		})
	}

}

func (m *mockedUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("User not found")
}

func (m *mockedUserStore) GetUserById(id int) (*types.User, error) {
	return nil, fmt.Errorf("User not found")
}

func (m *mockedUserStore) CreateUser(types.User) error {
	return nil
}
