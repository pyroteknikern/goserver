package main

import (
    "strconv"
    "html/template"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
    reddit "goserver/reddit"
)

func memeHandler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path[6:] {
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
