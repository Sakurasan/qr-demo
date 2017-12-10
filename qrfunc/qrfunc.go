package qrfunc

import qrcode "github.com/skip2/go-qrcode"
import "fmt"

func Qrfile() {
    err := qrcode.WriteFile("http://github.com/sakurasan/qrdemo4go", qrcode.Medium, 256, "qr.png")
    if err != nil {
        fmt.Println("write error")
	}
}
