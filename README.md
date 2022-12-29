# go-enum

In golang, there is not enum package, whatever official or community. So I write this package based on my own
experience.

## Usage

```go
package main

import (
	"github.com/leiyang23/go-enum"
	"encoding/json"
	"fmt"
)

// ColorEnum 自定义枚举类型
type ColorEnum struct {
    RED      enum.Enum
    BLUE     enum.Enum
    BLACK    enum.Enum
    YELLOW   enum.Enum
    BigGreen enum.Enum
}

func main() {
    Color := new(ColorEnum)
    err := enum.MakeEnum(Color)
    if err != nil {
        panic(err)
    }

    fmt.Println(enum.Validate(Color, 3))
    fmt.Println(enum.Validate(Color, 8))

    buf, err := json.Marshal(Color.BLUE)
    if err != nil {
        panic(err)
    }
    fmt.Printf("marshal json: %s\n", buf)
}
```

## License

[MIT](LICENSE)
