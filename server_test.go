package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/kazuhe/gocomm/data"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

// TestMain テストの共通処理
func TestMain(m *testing.M) {
	setUp()

	// 個々のテストケース関数を呼び出す
	code := m.Run()
	os.Exit(code)
}

// setUp 全てのテストケースの前処理
func setUp() {
	// 各テストケースは独立して実行されるため
	// テスト毎にマルチプレクサを生成
	mux = http.NewServeMux()
	// テスト対象のハンドルを付加
	mux.HandleFunc("/post/", handleRequest)
	// レスポンス受け取るHTTPレスポンスレコーダーを定義
	writer = httptest.NewRecorder()
}

// TestHandleGet GETリクエストで正常にJSONを返しているかテスト
func TestHandleGet(t *testing.T) {
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

// TestHandlePut PUTリクエストでエラーが返されないかテスト
func TestHandlePut(t *testing.T) {
	json := strings.NewReader(`{"content":"Updates post","autor":"kazuhe"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
