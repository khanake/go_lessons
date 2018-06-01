// 関数
// 返り値をもらいたい場合
package main

import "fmt"

func hi(name string) string { // 返り値がある場合は型の宣言が必須
  msg := "Hi!" + name
  return msg
}

func main() {
  fmt.Println(hi("Tanaka"))
}
