import qrcode "github.com/skip2/go-qrcode"

Create a PNG image:

  var png []byte
  png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
Create a PNG image and write to a file:

  err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")