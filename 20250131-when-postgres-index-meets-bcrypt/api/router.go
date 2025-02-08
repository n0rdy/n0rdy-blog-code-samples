package api

import (
	"20250131-when-postgres-index-meets-bcrypt/service"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ApiRouter struct {
	service *service.Service
}

func NewApiRouter(service *service.Service) *ApiRouter {
	return &ApiRouter{service: service}
}

func (ap *ApiRouter) NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/users", ap.handleGetUsers)
		r.Get("/users/{ssn}", ap.handleGetUser)
	})
	return router
}

func (ap *ApiRouter) handleGetUsers(w http.ResponseWriter, req *http.Request) {
	users, err := ap.service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ap.sendJsonResponse(w, http.StatusOK, users)
}

func (ap *ApiRouter) handleGetUser(w http.ResponseWriter, req *http.Request) {
	ssn := chi.URLParam(req, "ssn")
	user, err := ap.service.GetUser(ssn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ap.sendJsonResponse(w, http.StatusOK, user)
}

func (ap *ApiRouter) sendJsonResponse(
	w http.ResponseWriter,
	httpCode int,
	payload interface{},
) {
	respBody, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(respBody)
}
