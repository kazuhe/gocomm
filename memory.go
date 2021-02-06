package main

import "fmt"

// Post 投稿を表現する構造体（メモリ中に保管するデータ）
type Post struct {
	ID      int
	Content string
	Author  string
}

// PostByID ユニークなIDと実際の構造体Postを対応付けるために
// キー(ユニークなID): バリュー(Postへのポインタ)を設定する
var PostByID map[int]*Post

// PostsByAuthor 著者名と実際の構造体Postを対応付けるために
// キー(著者名): バリュー(Postへのポインタ)を設定する
var PostsByAuthor map[string][]*Post

// store 構造体Postへのポインタを'PostByID'と'PostsByAuthor'へ格納
func store(post Post) {
	PostByID[post.ID] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func callMemory() {
	PostByID = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{ID: 1, Content: "Hello Post", Author: "kazuhe1"}
	post2 := Post{ID: 2, Content: "Post2", Author: "kazuhe2"}
	post3 := Post{ID: 3, Content: "Post3", Author: "kazuhe3"}

	store(post1)
	store(post2)
	store(post3)

	fmt.Println(PostByID[1])
	fmt.Println(PostByID[2])

	for _, post := range PostsByAuthor["kazuhe1"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["kazuhe2"] {
		fmt.Println(post)
	}
}
