package texttoqr

import (
	"github.com/skip2/go-qrcode"
	"log"
)

// TextToQr Converts Text to QR code
func TextToQr(text string) {
	err := qrcode.WriteFile(text, qrcode.Medium, 256, "./output/qr.png")
	if err != nil {
		log.Fatal(err)
	}
}
