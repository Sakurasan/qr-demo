/*
  https://golang.org/doc/effective_go.html#web_server
*/

package qrfunc

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"
)

const headInfo = `
 *************************************************************
 ** a demo for go with Qrcode api .                         **
 ** by: Sakurasan                                           **
 ** http://github.com/sakurasan/                            **
 ** 										                **
 ** Please support genuine!!!                               **
 ** listen on 0.0.0.0:9000...                               **
 ** You can use http://127.0.0.1:9000 as license server     **
 *************************************************************
`

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func test() {
	flag.Parse()
	log.Println(time.Now().Format("2006-01-02 15:04:05"), "\n", headInfo)
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}"/>
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
<input maxLength=1024 size=70 name=s value="" title="Text to QR Encode"><input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
