package main

import "fmt"
import "www.wuyupei.top/http/util"
import "net/http"

func main() {
	fmt.Printf("\"a\": %v\n", "a")
	util.Foo()
	util.Bar()
	fmt.Printf("\"开始。。。\": %v\n", "开始。。。")

	http.HandleFunc("/go", handleGo)

	http.ListenAndServe("127.0.0.1:8000", nil)
}

func handleGo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("连接成功....")
	fmt.Printf("methouds: %v\n", r.Method)
	fmt.Printf("Path: %v\n", r.URL.Path)
	fmt.Printf("Header: %v\n", r.Header)
	fmt.Printf("Body: %v\n", r.Body)

	hi := "hi"

	w.Write([]byte(hi))
}
