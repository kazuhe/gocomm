package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// uploaded アップロードされたファイルを受信
func uploaded(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

// write Writeメソッドを使ってレスポンスを送信するための書き込み
func write(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>gocomm</title></head>
	<body><h1>Gocomm</h1></body>
	</html>`
	w.Write([]byte(str))
}

// notImplemented WriteHeaderメソッドを使ってレスポンスヘッダの書き込み
func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "501 Not Implemented")
}

// Post ユーザーの投稿を表す構造体
type Post struct {
	User    string
	Threads []string
}

// jsonWriter ResponseWriterを直接使ってjson出力の書き込み
func jsonWriter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Kazuhe",
		Threads: []string{"1番目", "2番目", "3番目"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
