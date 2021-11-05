package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	// t, err := template.ParseFiles("./index.tmpl")
	// if err != nil {
	// 	fmt.Printf("parse template err,failed:%v", err)
	// 	return
	// }
	// 自定义模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template err,faile%v", err)
		return
	}
	// 渲染模板
	name := "Joker"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("listen serve err,failed:%v", err)
		return
	}
}
