package controller

import "net/http"

type AuthController struct{}

func (ctrl *AuthController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("auth controller v1 pong"))
}
