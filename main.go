package main

import (
	"flag"
	"fmt"
)

var (
	length   uint
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	flag.UintVar(&length, "l", 6, "length of the tokens")
}

func main() {
	println("anvilward, dwarf spy")

	for url := range Resolve(Enumerate(length, alphabet)) {
		fmt.Println(url)
	}
}
