package main

import "fmt"

func main() {
	greeting := greetHello("en")
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
	"en": "Hello",
	"fr": "bonjour",
	"ru": "priviyet",
}

func greetHello(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return "unsupported language"
	}
	return greeting
}
