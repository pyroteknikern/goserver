package main

import (
    "log"
    "html/template"
    "net/http"
    "database/sql"
    "home/arvyd/repos/goserver/users"
)


func getUser(r *http.Request) (users.User) {
    username := r.FormValue("username")
    password := r.FormValue("password")
    return users.User{
        Username: username,
        Password: password,
    }
}


func handler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/home":
        homePage(w, r)
    case "/sign-up":
        signUp(w, r)
    case "/sign-in":
        signIn(w, r)
    default:
        pageNotFound(w)
    }
}

func signUp(w http.ResponseWriter, r *http.Request) {
    newUser := getUser(r)
}


func signIn(w http.ResponseWriter, r *http.Request) {

}

func homePage(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "html/home.html")
}

func renderPage(w http.ResponseWriter, documentPath string) (error) {
    t, err := template.ParseFiles(documentPath)
    if err != nil {
        return err
    }
    t.Execute(w, nil)
    return nil
}

func pageNotFound(w http.ResponseWriter) {
    t, _ := template.ParseFiles("html/404.html")
    t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}




