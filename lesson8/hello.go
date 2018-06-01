// 関数２
// 返り値を複数持つことができる

package main

import "fmt"

func swap(a, b int) (int, int) {
  return b, a
}

func main() {
  fmt.Println(swap(5,2))
}
