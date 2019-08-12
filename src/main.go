package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
        <head>
          <title>Golang web testpage</title>
        </head>
        <body>
          hello world!
        </body>
      </html>
		`))
  })
  
  // Webサーバーを開始
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}