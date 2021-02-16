package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kazuhe/gocomm/data"
)

// TestHandleGet GETリクエストで正常にJSONを返しているかテスト
func TestHandleGet(t *testing.T) {
	// 各テストケースは独立して実行されるため
	// テスト関数内でマルチプレクサを生成
	mux := http.NewServeMux()
	// テスト対象のハンドルを付加
	mux.HandleFunc("/post/", handleRequest)

	// レスポンス受け取るHTTPレスポンスレコーダーを定義
	writer := httptest.NewRecorder()
	// テスト対象へのリクエストを作成
	request, _ := http.NewRequest("GET", "/post/1", nil)

	// マルチプレクサにリクエストを送信
	mux.ServeHTTP(writer, request)

	// エラーチェック
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	// テスト対象から受け取ったコンテンツ(JSON)を構造体Postに組み換える
	var post data.Post
	json.Unmarshal(writer.Body.Bytes(), &post)

	// JSONを取得できているかテスト
	if post.ID != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}
