package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"strconv"

	"github.com/kazuhe/gocomm/data"
)

func main() {
	// Heroku環境変数からポート番号を取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr: ":" + port,
	}

	// /post/へのリクエストをハンドラ関数'handleRequest'へリダイレクト
	// 全てのハンドラ関数は第1引数に'ResponseWriter'をとり、
	// 第2引数に'Request'をとるので改めて引数を渡す必要はない
	http.HandleFunc("/post/", handleRequest)
	log.Println("start http listenig :" + port)
	server.ListenAndServe()
}

// handleRequest リクエストを正しい関数に振り分けるためのハンドラ
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 検証のためにリクエストに含まれる情報を出力
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))

	var err error

	// リクエストメソッドに応じてそれぞれのCRUD関数に作業を振り分ける
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	// リクエスト自体に関わるエラー処理
	// エラーがあれば詳細とステータス500を返す
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleGet /post/<id> GETリクエストに応じて投稿を読み出す関数
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// URLのパスを抽出
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	// メソッドRetriveでidを元にDBの値を取得して構造体Postを作成
	post, err := data.Retrive(id)
	if err != nil {
		return
	}

	// 構造体PostをJSONフォーマットのバイト列に変換
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}

	// バイト列をResponseWriterに書き出す
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// handlePost POSTリクエストに応じて投稿を作成する関数
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	// コンテンツの長さをサイズとしたバイト列を作成
	len := r.ContentLength
	body := make([]byte, len)

	// コンテンツ(JSON)を読み込む
	r.Body.Read(body)

	// コンテンツ(JSON)を構造体Postに組み換える
	var post data.Post
	json.Unmarshal(body, &post)

	// メソッドCreateで構造体PostをDBに保存
	err = post.Create()
	if err != nil {
		return
	}

	// ステータス200を返す
	w.WriteHeader(200)
	return
}

// handlePut /post/<id> PUTリクエストに応じて投稿を更新する関数
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	// URLのパスを抽出
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println("Get path error:", err)
		return
	}

	// メソッドRetriveでidを元にDBの値を取得して構造体Postを作成
	post, err := data.Retrive(id)
	if err != nil {
		fmt.Println("Retrive() error:", err)
		return
	}

	// コンテンツの長さをサイズとしたバイト列を作成
	len := r.ContentLength
	body := make([]byte, len)

	// リクエスト本体からコンテンツ(JSON)を読み込む
	r.Body.Read(body)

	// コンテンツ(JSON)を構造体Postに組み換える
	json.Unmarshal(body, &post)

	// メソッドUpdateでDBに保存されている値を更新
	err = post.Update()
	if err != nil {
		fmt.Println("Update() error:", err)
		return
	}

	// ステータス200を返す
	w.WriteHeader(200)
	return
}

// handleDelete /post/<id> DELETEリクエストに応じて投稿を削除する関数
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	// URLのパスを抽出
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	// メソッドRetriveでidを元にDBの値を取得して構造体Postを作成
	post, err := data.Retrive(id)
	if err != nil {
		return
	}

	// メソッドDeleteでDBに保存されている値を削除
	err = post.Delete()
	if err != nil {
		return
	}

	// ステータス200を返す
	w.WriteHeader(200)
	return
}
