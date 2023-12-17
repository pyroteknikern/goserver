package main

import (
    "fmt"
    "log"
    "html/template"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func ReadUserIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
    return IPAddress
}

func startHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path, "\n", ReadUserIP(r))
    switch r.URL.Path {
    case "/":
        homePage(w, r)
    default:
        pageNotFound(w)
    }
}

func currentUser(r *http.Request) Data {
    c, err := r.Cookie("token")
    var d Data
    if err != nil {
        d = Data{
            Auth: false,
            Username: "",
        }
        fmt.Println(err)
    } else {
        username := tokenUsername(c.Value)
        d = Data{
            Auth: true,
            Username: username,
        }
    }
    return d
}

func homePage(w http.ResponseWriter, r *http.Request) {
    d := currentUser(r) 
    d.Title = "Home"
    renderPage(w, "html/home.html", d)
}

type Data struct {
    Auth bool
    Username string
    Title string
}

func renderPage(w http.ResponseWriter, htmlPath string, d interface{}) (error) {
    tmpl, err := template.ParseFiles(htmlPath, "html/top.html")
    if err != nil {
        fmt.Println(err)
        return err
    }
    tmpl.ExecuteTemplate(w, "header", d)
    tmpl.Execute(w, d)
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
    http.HandleFunc("/auth/memes/", memeHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}




