package main

import (
	// "bytes"
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	qrcode "github.com/skip2/go-qrcode"
	// "github.com/qpliu/qrencode-go/qrencode"
)

const headInfo = `
 *************************************************************
 ** a demo for go with Qrcode api .                         **
 ** by: Sakurasan                                           **
 ** http://github.com/sakurasan/                            **
 ** Please support genuine!!!                               **
 ** listen on 0.0.0.0:9000...                               **
 ** You can use http://127.0.0.1:9000 as license server     **
 ** Ctrl + C --> quit. 										**
 *************************************************************
`

var qraddr = flag.String("addr", ":9000", "http service address") // Q=17, R=18

var qrtempl = template.Must(template.New("qr").Parse(templateStr))

func main() {
	log.Println(time.Now().Format("2006-01-02 15:04:05"), "\n", headInfo)
	http.Handle("/qr", http.HandlerFunc(QR))
	http.Handle("/", http.HandlerFunc(MainController))
	err := http.ListenAndServe(*qraddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func MainController(w http.ResponseWriter, r *http.Request) {
	qrtempl.Execute(w, "")
}

func QR(w http.ResponseWriter, req *http.Request) {
	var png []byte
	qrstr := req.FormValue("s")
	if qrstr == "" {
		qrstr = " "
	}
	var size int
	var err error
	qrsize := req.FormValue("size")
	if qrsize == "" {
		size = 256
	} else {
		size, err = strconv.Atoi(req.FormValue("size"))
		if err != nil {
			size = 256
		}
	}
	png, _ = qrcode.Encode(qrstr, qrcode.Medium, size)

	// qrtempl.Execute(w, string(png))
	w.Write(png)
}

const templateStr = `
<html>
<head>
<meta charset="utf-8">
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<!--img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}"/-->
<img src="{{.}}"/>
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/qr" name=f method="GET">
<input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
<input type=submit value="Show QR" >
</form>
</body>
</html>
`
