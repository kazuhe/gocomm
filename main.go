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
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	// 自己署名SSL証明書とサーバの秘密鍵の生成
	gencert()

	// HTTPSで運用するには'ListenAndServeTLS'を使用
	// cert.pem: SSL証明書, key.pem: サーバ用の秘密鍵
	server.ListenAndServeTLS("cert.pem", "key.pem")

}
