package middleware

import (
    "net/http"
    "strings"
    "rest_api/internal/apps/register/helper"
    "rest_api/internal/apps/register/model/web"
)

type AuthMiddleware struct {
    Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
    return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    // Check if the request URL is for the Swagger UI, Swagger JSON, login, or logout
    if strings.HasPrefix(request.URL.Path, "/swagger-ui") || 
       strings.HasPrefix(request.URL.Path, "/swagger/") || 
       request.URL.Path == "/api/login" || 
       request.URL.Path == "/api/logout" {
        // If it is, bypass the middleware and serve the request directly
        middleware.Handler.ServeHTTP(writer, request)
        return
    }

    // Otherwise, perform the usual authentication check
    if "RAHASIA" == request.Header.Get("X-API-Key") {
        // ok
        middleware.Handler.ServeHTTP(writer, request)
    } else {
        // error
        writer.Header().Set("Content-Type", "application/json")
        writer.WriteHeader(http.StatusUnauthorized)

        webResponse := web.WebResponse{
            Code:   http.StatusUnauthorized,
            Status: "UNAUTHORIZED",
        }

        helper.WriteToResponseBody(writer, webResponse)
    }
}

