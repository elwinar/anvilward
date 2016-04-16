package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func Resolve(tokens <-chan string) <-chan string {
	var out = make(chan string)

	// The connection used to communicate with the server. I intentionnaly use
	// the HTTP protocol over the (now standard) HTTPS one to speed-up the
	// exchange, as I really don't care about encryption of this.
	conn, err := net.Dial("tcp", "goo.gl:80")
	if err != nil {
		panic(err)
	}

	// The buffer will be used to read the responses.
	var buf = bufio.NewReader(conn)

	go func() {
		for token := range tokens {
			fmt.Fprintf(conn, "HEAD /%s HTTP/1.1\r\nHost: goo.gl\r\n\r\n", token)

			res, err := http.ReadResponse(buf, nil)
			if err != nil {
				continue
			}

			if res.StatusCode != 301 {
				continue
			}

			out <- res.Header.Get("Location")
		}
		close(out)
	}()

	return out
}
