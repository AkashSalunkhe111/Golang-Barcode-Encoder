package texttoqr

import "github.com/skip2/go-qrcode"

import "fmt"

// TextToQr Converts Text to QR code
func TextToQr(text string) {
	err := qrcode.WriteFile(text, qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println(err)
	}
}
