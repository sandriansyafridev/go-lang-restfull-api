package middleware

import (
	"errors"
	"golangapi/helper"
	"golangapi/model/response"
	"net/http"
)

type authorizationMiddleware struct {
	Handler http.Handler
}

func (authorizationMiddleware *authorizationMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if AuthorizationHeader := r.Header.Get("Authorization"); AuthorizationHeader != "" && AuthorizationHeader == "RAHASIA" {
		authorizationMiddleware.Handler.ServeHTTP(w, r) //next request
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		err := errors.New("Unauthorization")
		responseError := response.BuildResponseError("Not allow access this page", err.Error(), response.EmptyObject{})
		helper.WriteRequestBody(w, responseError)
		return
	}

}

func NewAuthorizationMiddleware(handler http.Handler) http.Handler {
	return &authorizationMiddleware{
		Handler: handler,
	}
}
