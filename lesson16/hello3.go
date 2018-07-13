// 省略形

package main

import (
  "fmt"
)

func main() {
  m := map[string]int{"hanake":200, "tanaka":100}
  for k, v := range m {
    fmt.Println(k,v)
  }
}

