package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func saveDataToFile() {
	// ioutilを使ってファイルを読み書きする例
	data := []byte("Hello World!!!\n")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	// 構造体Fileを使ってフィルを読み書きする例
	file1, _ := os.Create("data2")
	defer file1.Close()
	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	// 構造体Fileを使ってフィルを読み込む例
	file2, _ := os.Open("data2")
	defer file2.Close()
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println("read2", string(read2))
}
