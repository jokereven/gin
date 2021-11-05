package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 渲染模板
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 解析模板之前自定义一个函数
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("parse template err,failed%v", err)
		return
	}
	// 渲染模板
	str1 := "<script>alert('Hello');</script>"
	str2 := "<a href='http://code520.com.cn'>JokerEven</a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("listen serve err,failed%v", err)
		return
	}
}
