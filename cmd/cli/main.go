package main

import (
	"flag"
	"github.com/AkashSalunkhe111/Golang-Barcode-Encoder/internal/texttoqr"
)

func main() {
	textPtr := flag.String("text", "Hello World", "a string")
	flag.Parse()
	texttoqr.TextToQr(*textPtr)
}
