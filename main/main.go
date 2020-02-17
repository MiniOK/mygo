package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	// 注意，如果没有调用 ParseFrom 方法，下面无法获取表单的数据
	fmt.Println(r.Form) // 这些信息都是输出到服务器端打印的信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	// 这个写入到 w 的是输入到客户端的
	// Fprintf() 函数是将格式化字符串输出到一个文件设备中
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求方式
	if r.Method == "GET" {
		fmt.Println("have error here")
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		// 请求是否登录数据，那么执行登录的逻辑判断
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func main() {
	// 设置访问的路由
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
