package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var fileName string
	// fmt.Print("Please enter the file name without empty space:")
	// fmt.Scan(&fileName)

	fmt.Print("Please enter a message:")
	text := bufio.NewReader(os.Stdin)
	inputText, _ := text.ReadString('\n')
	p1 := &Page{Titile: "TestPage", Body: []byte(inputText)}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

type Page struct {
	Titile string
	Body   []byte
}

func (p *Page) save() error {
	filename := p.Titile + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Titile: title, Body: body}, nil
}
