package data

import (
	"database/sql"
	"fmt"

	// PostgreSQLのデータベースドライバ
	_ "github.com/lib/pq"
)

// Post ユーザーの投稿を表す構造体
type Post struct {
	ID      int
	Content string
	Author  string
}

// DB データベースへのハンドルであり、データベース接続のプールを表す
var DB *sql.DB

// init 初期化関数でデータベースのハンドルを生成
func init() {
	var err error
	// 'sql.Open'は単にその後のDBへの接続に必要になる構造体を設定するだけでデータベースに接続する訳ではない
	DB, err = sql.Open("postgres", "user=kazuhe dbname=kazuhe password=kazuhe sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Posts 最新の投稿からパラメータで受け取ったn個の投稿を取得する
func Posts(limit int) (posts []Post, err error) {
	// 'Query'を使ってSQLから複数の行（*sql.Rows）を制限値付きで取得
	rows, err := DB.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	// *sql.Rowsの'Next'メソッドを使って複数行の中から単一行を取得
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// GetPostByID パラメータIDによって1件のPostを取得
func GetPostByID(id int) (post Post, err error) {
	post = Post{}
	// SQLのselectコマンドを使って取得したデータ（id, content, author）をpostに参照渡し
	err = DB.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

// Create 新規投稿の生成
func (post *Post) Create() (err error) {
	// SQLのプリペアドステートメント（レコード作成時に特定の値を当てはめることができる）の定義
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	// ステートメントをプリペアドステートメントとして作成するためにDB.Prepareに渡す
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// プリペアドステートメントを実行
	// 'QueryRow'で構造体sql.Row（最初の1つだけの）を返す, 'Scan'は行の中の値を引数にコピーする
	// つまり、'post.Content'と'post.Author'をDBに挿入にした後に、SQLクエリによって返された
	// idフィールドの値（DB側で生成される自動増分値）を'&post.ID'に設定している
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

// Update 投稿内容の更新
func (post *Post) Update() (err error) {
	// 'Exec'は'QueryRow'と違って返される値に関心がない, 構造体Postを更新する必要も無いのでフィールドの値をもってupdateするだけ
	_, err = DB.Exec("update posts set content = $2, author =$3 where id = $1, post.ID, post.Content, post.Author")
	return
}

// Delete 投稿の削除
func (post *Post) Delete() (err error) {
	// Postのidをもとにdelete処理
	_, err = DB.Exec("delete from posts where id = $1, post.ID")
	return
}

// DBConnect DBの挙動確認用
func DBConnect() {
	post := Post{Content: "Hello World!", Author: "kazuhe"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPostByID(post.ID)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde"
	readPost.Author = "kazupi"
	readPost.Update()

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()
}
