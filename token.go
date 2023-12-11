package main

import (
    "fmt"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var secretKey = []byte("THIS IS A SECRET")

func createToken(username string) (string, time.Time, error) {
    expTime := time.Now().Add(time.Hour * 24)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username, 
    })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", expTime, err
    }
    return tokenString,expTime, nil
}

func verifyToken(tokenString string) error {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {return secretKey, nil})
    if err != nil {
        return err
    }

    if !token.Valid {
        return fmt.Errorf("invalid token")
    }

    return nil
}
type MyCustomClaims struct {
    Username string `json:"username"`
}
func tokenUsername(tokenString string) string {
    token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {return secretKey, nil})
    claims := token.Claims.(jwt.MapClaims)
    return claims["username"].(string) 
}






