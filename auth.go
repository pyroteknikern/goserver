package main

import (
    "net/http"
    _ "github.com/mattn/go-sqlite3"
    users "goserver/users"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path[5:] {
    case "/sign-up-page":
        signUpPage(w, r)
    case "/sign-up":
        signUp(w, r)

    case "/sign-in-page":
        signInPage(w, r)
    case "/sign-in":
        signIn(w, r)
    default:
        pageNotFound(w)
    }
}

func signInPage(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "html/sign-in.html")
}

func signIn(w http.ResponseWriter, r *http.Request) {
    user := getUser(r)
    err := users.FindUser(user)
    if err != nil {
        http.Redirect(w, r, "/sign-in-page", http.StatusFound)
        return
    }
    http.Redirect(w, r, "/home", http.StatusFound) 
}


func signUpPage(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "html/sign-up.html")
}

func signUp(w http.ResponseWriter, r *http.Request) {
    newUser := getUser(r)
    err := users.CreateUser(newUser)
    if err != nil {
        http.Redirect(w, r, "/sign-up-page", http.StatusFound)
        return
    }
    http.Redirect(w, r, "/home", http.StatusFound) 
}


func getUser(r *http.Request) (users.User) {
    username := r.FormValue("username")
    password := r.FormValue("password")
    return users.User{
        Username: username,
        Password: password,
    }
}
