package main

import (
	"fmt"
	collection "github.com/MytsV/list-go"
)

func main() {
	list := collection.List{}
	list.Append('a')
	fmt.Printf("Length after appending: %d", list.Length())
}