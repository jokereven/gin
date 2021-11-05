package template

import (
	"fmt"
	"html/template"
	"net/http"
)

func f(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数
	kua := func(name string) (string, error) {
		return name + "真帅", nil
	}
	// 定义模板
	// 解析模板
	t := template.New("index.tmpl")
	t.Funcs(template.FuncMap{
		"kua996": kua,
	})
	_, err := t.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("解析模板失败:", err)
		return
	}
	// 渲染模板
	t.Execute(w, "Joker")
}

func demo(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("解析模板失败:", err)
		return
	}
	// 渲染模板
	t.Execute(w, "陌生人")
}

func main() {
	http.HandleFunc("/", f)
	http.HandleFunc("/demo", demo)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("http serve start failed:", err)
		return
	}
}
