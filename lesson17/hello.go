// 構造体について
// 複数のまとまりを意味のあるまとまりとして定義することができる

package main

import (
  "fmt"
)

type user struct { //userの構造体 Railsのモデルのようなもの
  name string
  score int
}

func main() {
  u := new(user) //user構造体の初期化
  u.name = "hanake"
  u.score = 20
  fmt.Println(u)
}

