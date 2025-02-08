package main

import (
	"20250131-when-postgres-index-meets-bcrypt/common"
	"20250131-when-postgres-index-meets-bcrypt/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Route("/third-party/api/v1", func(r chi.Router) {
		r.Get("/user-info/{ssn}", handleUserInfo)
	})

	server := &http.Server{
		Addr:    "localhost:8081",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server shutdown")
		} else {
			fmt.Println("server failed")
			panic(err)
		}
	}
}

func handleUserInfo(w http.ResponseWriter, req *http.Request) {
	userInfo := common.ThirdPartyApiUserInfo{
		UserInfo: utils.GenRandomUserInfo(),
	}
	respBody, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
