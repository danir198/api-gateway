package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type APIGateway struct {
	Router              *mux.Router
	InventoryServiceURL string
	OrderServiceURL     string
	RateLimiter         *limiter.Limiter
}

func NewAPIGateway(inventoryServiceURL, orderServiceURL string) *APIGateway {
	rateLimiter := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	gateway := &APIGateway{
		Router:              mux.NewRouter(),
		InventoryServiceURL: inventoryServiceURL,
		OrderServiceURL:     orderServiceURL, // Initialize Order Service URL
		RateLimiter:         rateLimiter,
	}
	gateway.routes()
	return gateway
}

func (g *APIGateway) routes() {

	// Inventory service routes
	g.Router.HandleFunc("/products/{id}/availability", g.rateLimit(g.authenticate(g.routeRequest(g.InventoryServiceURL)))).Methods("GET")
	g.Router.HandleFunc("/products/{id}/inventory", g.rateLimit(g.authenticate(g.routeRequest(g.InventoryServiceURL)))).Methods("PUT")
	g.Router.HandleFunc("/products/{id}", g.rateLimit(g.authenticate(g.routeRequest(g.InventoryServiceURL)))).Methods("GET")
	g.Router.HandleFunc("/products", g.rateLimit(g.authenticate(g.routeRequest(g.InventoryServiceURL)))).Methods("POST")
	g.Router.HandleFunc("/products/{id}", g.rateLimit(g.authenticate(g.routeRequest(g.InventoryServiceURL)))).Methods("DELETE")
	g.Router.HandleFunc("/products", g.rateLimit(g.authenticate(g.routeRequest(g.InventoryServiceURL)))).Methods("GET")
	g.Router.HandleFunc("/products/search", g.rateLimit(g.authenticate(g.routeRequest(g.InventoryServiceURL)))).Methods("GET")
	g.Router.HandleFunc("/health", g.rateLimit(http.HandlerFunc(g.HealthCheckHandler)))

	// Order service routes
	g.Router.HandleFunc("/orders", g.rateLimit(g.authenticate(g.routeRequest(g.OrderServiceURL)))).Methods("POST")
	g.Router.HandleFunc("/orders", g.rateLimit(g.authenticate(g.routeRequest(g.OrderServiceURL)))).Methods("GET")
	g.Router.HandleFunc("/orders/{id}", g.rateLimit(g.authenticate(g.routeRequest(g.OrderServiceURL)))).Methods("GET")
	g.Router.HandleFunc("/orders/{id}", g.rateLimit(g.authenticate(g.routeRequest(g.OrderServiceURL)))).Methods("PUT")
	g.Router.HandleFunc("/orders/{id}/cancel", g.rateLimit(g.authenticate(g.routeRequest(g.OrderServiceURL)))).Methods("POST")
}

func (g *APIGateway) routeRequest(serviceURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, _ := url.Parse(serviceURL)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	}
}

func (g *APIGateway) authenticate(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Unauthorized request")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			log.Println("Invalid token")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("Invalid token claims")
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			http.Error(w, "Invalid role", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userRole", role)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (g *APIGateway) rateLimit(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpError := tollbooth.LimitByRequest(g.RateLimiter, w, r)
		if httpError != nil {
			log.Println("Rate limit exceeded")
			http.Error(w, httpError.Message, httpError.StatusCode)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// HealthCheckHandler responds with a simple health check message
func (g *APIGateway) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API Gateway is healthy"))
}
