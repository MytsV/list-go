package main

import (
	"fmt"
	collection "github.com/MytsV/list-go"
)

func main() {
	list := collection.List{}
	list.Append('a')
	list.Append('b')
	list.Append('c')
	value, _ := list.Get(2)
	fmt.Printf("%c %d\n", value, list.Length())
}