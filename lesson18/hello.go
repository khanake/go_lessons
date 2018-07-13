// メソッドについて
// 構造体などのデータ型に紐付いた関数

package main

import (
  "fmt"
)

type user struct {
  name string
  score int
}

func (u user) show() { //構造体を紐付けるにはレシーバーが必要
  fmt.Printf("name:%s, score:%d", u.name, u.score)
}

func main() {
  u := user{name:"hanake", score:200}
  u.show()
}

