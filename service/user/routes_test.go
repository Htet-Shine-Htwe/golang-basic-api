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

	t.Run("should fail if the payload is invalid", func(t *testing.T) {

		payload := types.RegisterUserPayload{
			FirstName: "test",
			LastName:  "last",
			Email:     "asdfgmail.com",
			Password:  "asdffdsa",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d,got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("test user can correctly registered", func(t *testing.T) {

		payload := types.RegisterUserPayload{
			FirstName: "test",
			LastName:  "last",
			Email:     "asdf@gmail.com",
			Password:  "asdffdsa",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d,got %d", http.StatusCreated, rr.Code)
		}
	})
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
