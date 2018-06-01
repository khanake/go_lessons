//map
// キーと値のペア、ハッシュのようなもの

package main

import (
  "fmt"
)
func main() {
  m := make(map[string]int) //stringの空のMapを作る
  m["aaa"] = 100
  m["bbbb"] = 200
  fmt.Println(m)
}

