# Cool

Cool library for using better syntax in Go.

## Installation

```shell
go get -v github.com/Streamer272/cool
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/Streamer272/cool"
)

func main() {
	var s cool.String = cool.NewString("My custom string")
	fmt.Printf("%s\n", s.Replace("string", "syntax sugar library"))

	var i cool.Int = cool.NewIntFromString("69")
	i.Increment()
	fmt.Printf("%d is now equal to 70\n", i)
	i.MustBePositive() // panics if i is zero or lower

	cool.ExitSuccess()
}
```

## License

This project is licensed under [MIT](https://github.com/Streamer272/cool/blob/main/LICENSE) license.

## Contribution

Contributors are welcome.
