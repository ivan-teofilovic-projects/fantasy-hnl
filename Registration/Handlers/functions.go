package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	models "example.com/greetings/Models"
	utility "example.com/greetings/Utility"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

func SignIn(r *http.Request, db *sql.DB, w http.ResponseWriter) {
	var user models.Login
	var storedUser models.Login
	fmt.Println("inside of /signin")
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	sqlStatement := `SELECT Password FROM Users WHERE username=?;`
	result := db.QueryRow(sqlStatement, user.Username)
	fmt.Println(result)
	result.Scan(&storedUser.Password)
	fmt.Println(storedUser.Password)
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &models.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Register(r *http.Request, db *sql.DB, w http.ResponseWriter) {
	var user models.User
	sqlStatement := `INSERT INTO Users (Firstname, Lastname, Username, Email, Password) VALUES (?, ?, ?, ?, ?);`
	json.NewDecoder(r.Body).Decode(&user)
	hash, _ := utility.HashPassword(user.Password)
	_, err := db.Exec(sqlStatement, user.Firstname, user.Lastname, user.Username, user.Email, hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
