package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type PostType struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func parseJSON() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post PostType
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}
}

func genJSON() {
	post := PostType{
		ID:      1,
		Content: "JSON生成コンテンツ",
		Author: Author{
			ID:   2,
			Name: "kazuhe",
		},
		Comments: []Comment{
			{
				ID:      3,
				Content: "コメント",
				Author:  "Adam",
			},
			{
				ID:      4,
				Content: "コメント2",
				Author:  "Betty",
			},
		},
	}

	jsonFile, err := os.Create("genpost.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}
