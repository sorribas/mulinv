mulinv
======

Multiplicative inverse for Go's big integers.

Usage
-----

```go
package main

import (
  "github.com/sorribas/mulinv"
  "math/big"
  "fmt"
)

func main() {
  result := mulinv.MultiplicativeInverse(big.NewInt(42), big.NewInt(2017))
  fmt.Printf("%v\n", result) // 1969
}
```

License
-------

MIT
