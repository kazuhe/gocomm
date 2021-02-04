package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// setCookie クッキーをセット
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "first cookie value",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "second cookie value",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

// getCookie クッキーゲット
func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

// setMessage クッキーを使ったフラッシュメッセージ（セット）
func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	// 空白を使っているためヘッダ内ではURLエンコーディングする必要がある
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

// showMessage クッキーを使ったフラッシュメッセージ（ゲット/解析）
func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No Message!")
		}
	} else {
		// クッキー'flash'が見つかった場合はMaxAge（およびExpires）に
		// 負の数値を設定してSetCookieで上書きすることによって事実上の削除となる
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
