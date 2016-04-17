package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	length   = uint(7)
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func main() {
	println("anvilward, dwarf spy")

	for token := range Enumerate(length, alphabet) {
		time.Sleep(100 * time.Millisecond)

		res, err := http.Head(fmt.Sprintf("http://bit.ly/%s", token))
		if err != nil {
			log.Println(token, err)
			continue
		}

		if res.StatusCode != 301 {
			log.Println(token, res.Status)
			continue
		}

		log.Println(token, res.Header.Get("Location"))
	}
}
