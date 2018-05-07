package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"html/template"
	"errors"
	"log"
)
//这个是是干哈的
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	//API？
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "Wiki/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	//注意这个路径问题
	filename := "Wiki/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	//也是读取save
	//title := r.URL.Path[len("/save/"):]
	//读取这个页面的内容 body的内容
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	if err := p.save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	//读取view后面的路径，比如说test就读取了test

	//这个被什么代替了？？被正则代替了么
	//title := r.URL.Path[len("/view/"):]

	//从我们的test.txt里面读取东西存在p里面
	//如果load的文件不存在，那么久要redirect。 我们此时只存在test
	p, err := loadPage(title);
	if  err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	//title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)

}
////阅读下源码，看看这样做的好处是什么。他的好处应该是刚开始一次性Parsefiles所有的可能的files， 然后后面就不用一遍遍重复parsefile了。???
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, "Wiki/" + tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
var templates = template.Must(template.ParseFiles("Wiki/edit.html", "Wiki/view.html"))

//func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
//	t, err := template.ParseFiles("Wiki/" + tmpl + ".html")
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	if err := t.Execute(w, p); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//https://golang.org/doc/articles/wiki/
