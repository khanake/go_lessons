//空のインターフェースを定義する
// interface{}
// 型Switch

package main

import (
  "fmt"
)

func show(t interface{}) { // あらゆるデータ型をここで受けることが可能となる
  // 型Switch
  switch t.(type) { // 構造体が何者かを返す
  case japanese:
    fmt.Println("私は日本人です")
  default :
    fmt.Println("私は日本人ではありません")
  }

}

type greeter interface {
  greet()
}

type japanese struct {
}

type american struct {
}

func (j japanese) greet() {
  fmt.Println("こんにちは!")
}

func (j american) greet() {
  fmt.Println("Hello!")
}

// Japanese と American を同じスライスに入れたい場合にインターフェースを活用する

func main() {
  greeters := []greeter{japanese{}, american{}}
  for _, greeter := range greeters { // range は要素数だけ処理を行うもの, '_' はインデックスの省略, greeterは繰り返しの中で使われるgreetersの要素の変数
    greeter.greet()
    show(greeter)
  }
}
