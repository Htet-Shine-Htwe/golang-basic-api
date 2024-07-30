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

type mockUserStore struct {
}

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	testCases := []struct {
		name         string
		payload      types.RegisterUserPayload
		expectedCode int
	}{
		{
			name: "should fail if the payload is invalid",
			payload: types.RegisterUserPayload{
				FirstName: "test",
				LastName:  "last",
				Email:     "asdfgmail.com",
				Password:  "asdffdsa",
			},
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
		}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			marshalled, _ := json.Marshal(tt.payload)
			req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
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

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
