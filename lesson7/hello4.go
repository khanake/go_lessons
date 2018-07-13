// 関数
//名前付きのreturnする場合
package main

import "fmt"

func hi(name string) (msg string) { 
  msg = "Hi!" + name // すでにmsgは定義されている
  return // msgは省略することができる
}

func main() {
  fmt.Println(hi("Tanaka"))
}
