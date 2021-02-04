package main

import (
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	// mux マルチプレクサを生成
	mux := http.NewServeMux()

	// rootURLへのリクエストをハンドラ関数'index'へリダイレクト
	// 全てのハンドラ関数は第1引数に'ResponseWriter'をとり、
	// 第2引数に'Request'をとるので改めて引数を渡す必要はない
	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/uploaded", uploaded)
	mux.HandleFunc("/write", write)
	mux.HandleFunc("/501", notImplemented)
	mux.HandleFunc("/json", jsonWriter)
	mux.HandleFunc("/set_cookie", setCookie)
	mux.HandleFunc("/get_cookie", getCookie)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	// 自己署名SSL証明書とサーバの秘密鍵の生成
	gencert()

	// HTTP/2で動作するサーバを用意
	http2.ConfigureServer(server, &http2.Server{})

	// HTTPSで運用するには'ListenAndServeTLS'を使用
	// cert.pem: SSL証明書, key.pem: サーバ用の秘密鍵
	server.ListenAndServeTLS("cert.pem", "key.pem")

}
