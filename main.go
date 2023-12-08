package main

import (
    "log"
    "strconv"
    "html/template"
    "net/http"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    users "goserver/users"
    reddit "goserver/reddit"
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

    case "/sign-up-page":
        signUpPage(w, r)
    case "/sign-up":
        signUp(w, r)

    case "/sign-in-page":
        signInPage(w, r)
    case "/sign-in":
        signIn(w, r)

    case "/meme-send":
        memeSend(w, r)
    case "/meme-page":
        memePage(w, r)

    default:
        pageNotFound(w)
    }
}

var Images []string 
type Img struct {
    Image string
    Counter int
}

func memeSend(w http.ResponseWriter, r *http.Request) {
    var err error
    if len(Images) == 0 {
        Images, err = reddit.GetPosts()
    }
    if err != nil {
        return
    }
    counterString := r.FormValue("counter")
    counter, err := strconv.Atoi(counterString)
    if err != nil {
        counter = 0
    }

    counter += 1
    if counter >= len(Images) {
        counter = 0
    }

    img := Img{
        Image: Images[counter],
        Counter: counter,
    }
    tmpl, _ := template.ParseFiles("html/template.html")
    tmpl.Execute(w, img)
}

func memePage(w http.ResponseWriter, r *http.Request) {
    var err error
    if len(Images) == 0 {
        Images, err = reddit.GetPosts()
    }
    if err != nil {
        return
    }
    img := Img{
        Image: Images[0], 
        Counter: 0,
    }
    t, _ := template.New("images.html").ParseFiles("html/images.html")
    t.Execute(w, img)
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


func homePage(w http.ResponseWriter, r *http.Request) {
    renderPage(w, "html/home.html")
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
    err := createDatabase()
    if err != nil {
        log.Fatal(err)
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}




