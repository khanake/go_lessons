// 参照渡しのやりかた

package main

import (
  "fmt"
)

type user struct {
  name string
  score int
}

func (u user) show() { //この場合は値のコピーが行われる 値渡し
  fmt.Printf("name:%s, score:%d", u.name, u.score)
}

func (u *user) hit() { //こうすることで参照渡しとなり、アドレスを書き換えてくれる
  u.score++
}

func main() {
  u := user{name:"hanake", score:200}
  u.hit()
  u.show()
}

