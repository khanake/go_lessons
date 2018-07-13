// newと同時に値を入れることもできる

package main

import (
  "fmt"
)

type user struct { //userの構造体 Railsのモデルのようなもの
  name string
  score int
}

func main() {
  //u := user{"hanake",200}
  u := user{name:"hanake",score:200}
  fmt.Println(u)
}

