// 関数
// オプションを与える場合
package main

import "fmt"

func hi(lastname string, firstname string) { // 引数に変数名と型を与える
  fmt.Println("Hi! " + lastname + firstname)
}

func main() {
  hi("Tanaka", "Taro")
}
