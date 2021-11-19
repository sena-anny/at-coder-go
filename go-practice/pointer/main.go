package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v is %v years old", p.name, p.age)
}

func (p *Person) changeName(newName string) {
	fmt.Printf("バリュー：%v\n", p)
	fmt.Printf("ポインター：%p\n", p)
	p.name = newName
}

func (p *Person) getName() string {
	return p.name
}

type GitHubResponse []struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type customWriter struct{}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

func main() {
	p := &Person{"kj", 28}
	fmt.Printf("バリュー：%v\n", p)
	fmt.Printf("ポインター：%p\n", p)
	p.changeName("NEW")
	fmt.Println(p.getName())
	fmt.Printf("%s\n", *p)
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	writer := customWriter{}
	io.Copy(writer, resp.Body)
}
