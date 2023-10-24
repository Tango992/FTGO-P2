package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secretKey = []byte("12345")

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	claims := jwt.MapClaims{
		"username": "Tugas Meilyanto",
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	//Method NewWithClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Failed create token")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path

		if path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Token Not Found", http.StatusUnauthorized)
			return
		}

		// Method Parse
		parsedToken, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Algoritma tidak valid")
			}

			return secretKey, nil
		})

		if parsedToken.Valid {
			fmt.Println(parsedToken)
		}

		if parseErr != nil || !parsedToken.Valid {
			fmt.Println("Error while decode token : ", parseErr)
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Halaman ini butuh otentikasi"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	r.Use(AuthMiddleware)
	r.HandleFunc("/home", HomeHandler).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)

}
