package main

import (
    "time"
    "fmt"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
    users "goserver/users"
)
func SignedIn(r *http.Request) bool {
    c, err := r.Cookie("token")
    if err != nil {
        if err == http.ErrNoCookie {
            return false
        }
        return false
    }
    if verifyToken(c.Value) != nil {
        return false
    }
    return true
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(ReadUserIP(r), " ", r.URL.Path)
    switch r.URL.Path[5:] {
    case "/home":
        homePage(w, r)
    case "/sign-up-page":
        signUpPage(w, r)
    case "/sign-up":
        signUp(w, r)

    case "/sign-in-page":
        signInPage(w, r)
    case "/sign-in":
        signIn(w, r)
    case "/log-out":
        logOut(w, r)
    default:
        pageNotFound(w)
    }
}

func logOut(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name: "token",
        Expires: time.Now(),
    })

    http.Redirect(w, r, "/", http.StatusFound) 
}

func signInPage(w http.ResponseWriter, r *http.Request) {
    d := currentUser(r) 
    d.Title = "SignIn"
    renderPage(w, "html/sign-in.html", d)
}

func signIn(w http.ResponseWriter, r *http.Request) {
    user := getUser(r)
    truth, err := users.FindUser(user.Username)
    if err != nil {
        http.Redirect(w, r, "/auth/sign-in-page", http.StatusFound) 
        fmt.Println(err)
        return
    }
    if !truth {
        http.Redirect(w, r, "/auth/sign-in-page", http.StatusFound) 
        return
    }
    comp, err := users.ComparePassword(user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Println(err)
        return
    }
    if !comp {
        http.Redirect(w, r, "/auth/sign-in-page", http.StatusFound) 
        return
    }

    tokenString, expTime, err := createToken(user.Username)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Println(err)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name: "token",
        Value: tokenString,
        Expires: expTime,
    })
    http.Redirect(w, r, "/auth/home", http.StatusFound) 
}


func signUpPage(w http.ResponseWriter, r *http.Request) {
    d := currentUser(r) 
    d.Title = "SignUp"
    renderPage(w, "html/sign-up.html", d)
}

func signUp(w http.ResponseWriter, r *http.Request) {
    newUser := getUser(r)
    err := users.CreateUser(newUser)
    if err != nil {
        fmt.Println(err)
        http.Redirect(w, r, "/auth/sign-up-page", http.StatusFound)
        return
    }
    http.Redirect(w, r, "/auth/home", http.StatusFound) 
}


func getUser(r *http.Request) (users.User) {
    username := r.FormValue("username")
    password := r.FormValue("password")
    return users.User{
        Username: username,
        Password: password,
    }
}
