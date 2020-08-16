package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Post struct {
	ID       int       `json: "id"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Comments []Comment `json:comments`
}

type Comment struct {
	ID      int    `json: "id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Author struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	return
}

func main() {
	_, err := decode("post.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
