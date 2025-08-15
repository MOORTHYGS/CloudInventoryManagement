package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL == "" {
		supabaseURL = "https://opvifyrwxktzyuyskhme.supabase.co"
	}

	frontendOrigin := "https://localhost:8443"
	redirectAfterLogin := "https://localhost:9443/api/auth/callback"

	http.HandleFunc("/api/login/google", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w, r, frontendOrigin)
		loginWithGoogle(w, r, supabaseURL, redirectAfterLogin)
	})

	http.HandleFunc("/api/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w, r, frontendOrigin)
		authCallback(w, r, frontendOrigin)
	})

	http.Handle("/api/dashboard", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w, r, frontendOrigin)
		dashboardAPI(w, r)
	})))

	http.HandleFunc("/api/set-token", setTokenHandler)

	addr := ":9443"
	log.Printf("Backend API running at https://localhost%s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, "./cert.pem", "./key.pem", nil))
}

func enableCORS(w *http.ResponseWriter, r *http.Request, origin string) {
	(*w).Header().Set("Access-Control-Allow-Origin", origin)
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}

func loginWithGoogle(w http.ResponseWriter, r *http.Request, supabaseURL, redirectTo string) {
	redirectURL := fmt.Sprintf("%s/auth/v1/authorize?provider=google&redirect_to=%s", supabaseURL, redirectTo)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func authCallback(w http.ResponseWriter, r *http.Request, frontendOrigin string) {
	http.Redirect(w, r, frontendOrigin+"/dashboard.html", http.StatusFound)
}

// Add your Supabase JWT secret here (get it from your Supabase project settings)
var supabaseJWTSecret = []byte("0mPanXw9ag7z7yXLMIy/8MI6kCcD9J55nM74XYVpdogNVl+Z4m8EwHxzOmxCR/N3jP/OqKdpOQFRrMSzUuSLSw==")

// JWT validation function
func validateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return supabaseJWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

// Middleware Authentication for users by access token
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w, r, "https://localhost:8443") // Always set CORS headers

		if r.Method == "OPTIONS" { // Handle preflight
			w.WriteHeader(http.StatusOK)
			return
		}

		cookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value
		if strings.TrimSpace(tokenString) == "" {
			http.Error(w, "Unauthorized: empty token", http.StatusUnauthorized)
			return
		}
		_, err = validateJWT(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func dashboardAPI(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message": "Welcome to the dashboard!"}`))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func setTokenHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w, r, "https://localhost:8443") // <-- Add this

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req struct {
		AccessToken string `json:"access_token"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.AccessToken == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    req.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
	w.WriteHeader(http.StatusOK)
}
