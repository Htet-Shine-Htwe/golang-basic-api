package user

import (
	"fmt"
	"net/http"

	"github.com/dede182/revesion/service/auth"
	"github.com/dede182/revesion/types"
	"github.com/dede182/revesion/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", h.handleLogin)
	r.HandleFunc("/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	jsonHandler := utils.NewHandler()

	var loginPayload types.LoginUserPayload
	if err := jsonHandler.ValidateBody(r, loginPayload); err != nil {
		jsonHandler.WriteError(w, http.StatusBadRequest, err)
	}

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	jsonHandler := utils.NewHandler()

	var registerPayload types.RegisterUserPayload
	if err := jsonHandler.ValidateBody(r, &registerPayload); err != nil {
		jsonHandler.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validator.Struct(registerPayload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		jsonHandler.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", validationErrors))
		return
	}

	// check user with same email already exists
	_, err := h.store.GetUserByEmail(registerPayload.Email)
	if err == nil {
		jsonHandler.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with %s already exists", registerPayload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(registerPayload.Password)

	if err != nil {
		jsonHandler.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user with %s already exists", registerPayload.Email))
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: registerPayload.FirstName,
		LastName:  registerPayload.LastName,
		Email:     registerPayload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		jsonHandler.WriteError(w, http.StatusBadRequest, err)
	}

	jsonHandler.WriteJson(w, http.StatusCreated, nil)

}
