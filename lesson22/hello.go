// go言語でwebサーバーを作ってみる

package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // ほとんど決まり文句
  fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:]) // 詳しいことはドキュメントを見る
}

func main() {
  http.HandleFunc("/", handler) // ルートにアクセスされたらハンドラーを起動
  http.ListenAndServe(":8080", nil)
}

