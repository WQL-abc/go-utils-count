# What is this?
Easy counter

# Installation
```commandline
go get -u github.com/go-utils/count
```

# Usage
```go
import (
	"fmt"

	"github.com/go-utils/count"
)

func main() {
	if cnt, err := count.Do("something", "a"); err == nil {
		fmt.Println(cnt) // cnt = 0
	}

	if cnt, err := count.Do([]string{"1", "2", "3"}, "1"); err == nil {
		fmt.Println(cnt) // cnt = 1
	}

	if cnt, err := count.Do([]int{1, 2, 3}, 1); err == nil {
		fmt.Println(cnt) // cnt = 1
	}

	if cnt, err := count.Do("12321", "1"); err == nil {
		fmt.Println(cnt) // cnt = 2
	}

	if cnt, err := count.Do([]interface{}{1, "2", 3, "4", nil, nil}, nil); err == nil {
		fmt.Println(cnt) // cnt = 2
	}
}
```

Check [Go Playground](https://play.golang.org/p/Yee5NWWEx-q)

# Support
`string`, `slice`, `array`

# License
[MIT](./LICENSE)
