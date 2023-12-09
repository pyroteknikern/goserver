package main

import (
    "fmt"
    "log"
    "html/template"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func startHandler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/":
        renderPage(w, "html/home.html")
    default:
        pageNotFound(w)
    }
}
        

func renderPage(w http.ResponseWriter, documentPath string, ) (error) {
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

func createDatabase() (error) {

    db, err := sql.Open("sqlite3", "./db.sqlite3")
    if err != nil {
        return err
    }
    statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
    if err != nil {
        return err
    }
    statement.Exec()
    return nil
}


func main() {
    port := "8080"

    err := createDatabase()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Server started\nListening at port: %v\n", port)
    http.HandleFunc("/", startHandler)
    http.HandleFunc("/auth/", loginHandler)
    http.HandleFunc("/memes/", memeHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}




