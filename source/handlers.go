package source

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Input  string
	Banner string
	Result string
}

var UserData Data

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art-web" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	path := "./templates/"
	UserData.Input = r.Form.Get("input")
	UserData.Banner = r.Form.Get("banner")
	bannerFile := ReadTemplate(path + UserData.Banner + ".txt")
	fontParse := getTemplate(bannerFile, UserData.Banner)
	UserData.Result = "\n" + GenerateAsciiArt(UserData.Input, fontParse)
	
	http.Redirect(w, r, "/", http.StatusFound)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles("index.html")
	CheckError(err)

	t.Execute(w, UserData)
}
