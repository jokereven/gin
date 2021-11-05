package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err", err)
		return
	}
	// 渲染模板
	t.Execute(w, "Index")
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err%#v", err)
		return
	}
	// 解析模板
	t.Execute(w, "Home")
	// 渲染模板
}

func tmpindex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err%v", err)
		return
	}
	name := "Index"
	t.ExecuteTemplate(w, "index.tmpl", name)
}

func tmphome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err%v", err)
		return
	}
	name := "Home"
	t.ExecuteTemplate(w, "home.tmpl", name)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/tmpindex", tmpindex)
	http.HandleFunc("/tmphome", tmphome)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Http Serve Start Failed:", err)
		return
	}
}
