package user

import (
	"net/http"

	"github.com/dede182/revesion/types"
	"github.com/dede182/revesion/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", h.handleLogin)
	r.HandleFunc("/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	jsonHandler := utils.NewHandler()

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	jsonHandler := utils.NewHandler()
	jsonHandler.SetupHeader(w)

	var user types.RegisterUserPayload
	if err := jsonHandler.ValidateBody(r, user); err != nil {
		jsonHandler.WriteError(w, http.StatusBadRequest, err)
	}

}
