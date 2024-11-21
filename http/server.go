package http

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":3000", "http service address")
var templ = template.Must(template.New("QR").Parse(templateStr))

func StartServer() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}