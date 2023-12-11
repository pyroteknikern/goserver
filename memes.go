package main

import (
    "fmt"
    "strconv"
    "html/template"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
    reddit "goserver/reddit"
)

func memeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)

    if !loginRequired(w, r) { http.Redirect(w, r, "/kuk", http.StatusFound); return }
    
    switch r.URL.Path[11:] {
    case "/meme-send":
        memeSend(w, r)
    case "/meme-page":
        memePage(w, r)
    default:
        pageNotFound(w)
    }
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

var Images []string 

type Img struct {
    Username string
    Auth bool
    Title string
    Image string
    Counter int
}

func memePage(w http.ResponseWriter, r *http.Request) {
    var err error
    if len(Images) == 0 {
        Images, err = reddit.GetPosts()
    }
    if err != nil {
        return
    }
    c, _ := r.Cookie("token")
    username := tokenUsername(c.Value)
    img := Img{
        Username: username,
        Auth: true,
        Title: "Memes",
        Image: Images[0], 
        Counter: 0,
    }
    renderPage(w, "html/images.html", img)
}
