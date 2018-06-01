// 定数
package main

import "fmt"

func main() {
  const name = "aaa" // [const] [定数名]
  //name = "hanaie" 挿入不可能
  fmt.Println(name)

  //iota の使い方
  const (
    sun = iota
    mon
    tue
  )
  fmt.Println(sun, mon, tue)
}
