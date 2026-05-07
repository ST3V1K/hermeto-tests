package main

import (
	"fmt"

	"github.com/test/no-gosum/api"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println("ID:", api.NewID())
	fmt.Println("Lang:", language.English)
}
