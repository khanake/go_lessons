//map
// キーと値のペア、ハッシュのようなもの
// 値を指定して宣言することもできる

package main

import (
  "fmt"
)
func main() {
  m := map[string]int{"aaa":100, "bbb":200}
  fmt.Println(m)
  fmt.Println(len(m)) //要素の数を出力
  delete(m, "bbb")
  fmt.Println(m)
  a ,b := m["aaa"] //値を返すのと、その要素が存在しているかをbooleanで返す
  fmt.Println(a)
  fmt.Println(b)
}

