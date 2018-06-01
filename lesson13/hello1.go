package main

import (
  "fmt"
)

func main() {
  if score := 43; score > 80 { //ifの中だけで使う変数を定義
    fmt.Println("great")
  } else if score > 60 {
    fmt.Println("good")
  } else {
    fmt.Println("soso")
  }
  //fmt.Println(score)　これはエラーになる
}
