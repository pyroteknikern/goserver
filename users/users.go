package users

import (
    "database/sql"
    "errors"
)

type User struct {
    Username string
    Password string
}

func CreateUser(newUser User) (error) {
    db, err := sql.Open("sqlite3", "./db.sqlite3")
    if err != nil {
        return err
    }

    truth, err := FindUser(newUser.Username)
    if err != nil {
        return err
    }

    if truth {
        return errors.New("User already exists")
    }

    statement, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
    if err != nil {
        return err
    }

    statement.Exec(newUser.Username, newUser.Password)
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
    db, _ := sql.Open("sqlite3", "./db.sqlite3")

    var qPassword string
    err := db.QueryRow("SELECT password FROM users WHERE username=?", user.Username).Scan(&qPassword)
    if err != nil {
        return false, err
    }

    if user.Password == qPassword {
        return true, nil
    }
    return false, nil
}

