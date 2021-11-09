package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) changeName(newName string) {
	fmt.Printf("バリュー：%v\n", p)
	fmt.Printf("ポインター：%p\n", p)
	p.name = newName
}

func (p *Person) getName() string {
	return p.name
}
func main() {
	p := &Person{"kj", 28}
	fmt.Printf("バリュー：%v\n", p)
	fmt.Printf("ポインター：%p\n", p)
	p.changeName("NEW")
	fmt.Println(p.getName())

}
