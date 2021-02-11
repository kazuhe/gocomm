package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	var post PostType
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
