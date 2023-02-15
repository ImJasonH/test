package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	resp, err := http.Get("https://cgr.dev/v2")
	if err != nil {
		log.Fatal(err)
	}
	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, bytes.NewReader(b))
}
