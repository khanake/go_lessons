// 関数２
// 関数はデータ型なので変数に代入することもできる

package main

import "fmt"

func main() {
  f := func(a, b int) (int, int) { //関数名の省略が可能
    return b, a
  }
  fmt.Println(f(2, 3))
}
