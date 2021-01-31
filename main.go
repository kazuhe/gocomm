package main

import (
	"net/http"
)

func main() {
	// mux マルチプレクサを生成
	mux := http.NewServeMux()

	// rootURLへのリクエストをハンドラ関数'index'へリダイレクト
	// 全てのハンドラ関数は第1引数に'ResponseWriter'をとり、
	// 第2引数に'Request'をとるので改めて引数を渡す必要はない
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
