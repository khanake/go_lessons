//配列について
package main

import (
  "fmt"
)

func main() {
  var a [5]int //一個づつ宣言をしているこの場合は５個入れている
  a[2] = 1
  fmt.Println(a[2])
  fmt.Println(a)
}
