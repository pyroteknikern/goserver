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

    err = FindUser(newUser)
    if err == nil {
        return errors.New("user exists")
    }

    statement, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
    if err != nil {
        return err
    }

    statement.Exec(newUser.Username, newUser.Password)
    return nil
}

func FindUser(user User) (error) {
    db, err := sql.Open("sqlite3", "./db.sqlite3")
    if err != nil {
        return err
    }

    var username string
    err = db.QueryRow("SELECT username FROM users WHERE username=?", user.Username).Scan(&username)
    if err != nil {
        return err
    }

    return nil
}

