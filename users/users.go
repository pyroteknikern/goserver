package users

import (
    "database/sql"
    "errors"
    "golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
    passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(passwordHash), err
}

func checkPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

type User struct {
    Username string
    Password string
}

func CreateUser(newUser User) (error) {
    db, err := sql.Open("sqlite3", "./db.sqlite3")
    if err != nil {
        return err
    }

    truth, _ := FindUser(newUser.Username)

    if truth {
        return errors.New("User already exists")
    }

    statement, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
    if err != nil {
        return err
    }

    hash, err := hashPassword(newUser.Password)
    if err != nil {
        return err
    }

    statement.Exec(newUser.Username, hash)
    return nil
}

func FindUser(username string) (bool, error) {
    db, err := sql.Open("sqlite3", "./db.sqlite3")
    if err != nil {
        return false, err
    }

    var qUsername string
    err = db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&qUsername)
    if err != nil {
        return false, err
    }
    
    if username == qUsername {
        return true, nil
    }
    return false, nil
}

func ComparePassword(user User) (bool, error) {
    db, err := sql.Open("sqlite3", "./db.sqlite3")
    if err != nil {
        return false, err
    }

    var queryPasswordHash string
    err = db.QueryRow("SELECT password FROM users WHERE username=?", user.Username).Scan(&queryPasswordHash)
    if err != nil {
        return false, err
    }

    if checkPasswordHash(user.Password, queryPasswordHash) {
        return true, nil
    }
    return false, nil
}

