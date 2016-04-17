package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	length = flag.Uint("length", 7, "length of the tokens")
	host   = flag.String("host", "bit.ly", "url shortener host")
	delay  = flag.String("delay", "100ms", "delay between two attemps")
	help   = flag.Bool("help", false, "display the help message")

	https = flag.Bool("https", false, "use https to connect to the host")

	lower = flag.Bool("lower", true, "use alphabetical chars in the token (lowercases)")
	upper = flag.Bool("upper", true, "use alphabetical chars in the token (uppercase)")
	num   = flag.Bool("num", true, "use numerical chars in the token")
)

const (
	alphaLower = "abcdefghijklmnopqrstuvwxyz"
	alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaNum   = "0123456789"
)

func main() {
	println("anvilward, dwarf spy")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	d, err := time.ParseDuration(*delay)
	if err != nil {
		fmt.Println("invalid delay value:", err)
		os.Exit(1)
	}

	if *length < 1 {
		fmt.Println("invalid token length: must be > 1")
		os.Exit(1)
	}

	var alphabet = ""
	if *lower {
		alphabet += alphaLower
	}
	if *upper {
		alphabet += alphaUpper
	}
	if *num {
		alphabet += alphaNum
	}

	if len(alphabet) == 0 {
		fmt.Println("must use at least one of lower, upper, num")
		os.Exit(1)
	}

	if len(*host) == 0 {
		fmt.Println("invalid host")
		os.Exit(1)
	}

	var scheme = "http"
	if *https {
		scheme = "https"
	}

	for token := range Enumerate(*length, alphabet) {
		time.Sleep(d)

		res, err := http.Head(fmt.Sprintf("%s://%s/%s", scheme, *host, token))
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
