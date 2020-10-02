# spinner

Beautiful spinners for terminal.

### Usage:

> go get github.com/ltoddy/spinner

*basic.go*

```go
package main

import (
	"github.com/ltoddy/spinner"
	"time"
)

func main() {
	spin := spinner.NewSpinner()
	spin.Start()
	time.Sleep(time.Second * 2)
	spin.Stop()
}
```

*advance.go*

```go
package main

import (
	"github.com/ltoddy/spinner"
	"time"
)

func main() {
	spin := spinner.NewSpinnerBuilder().
		WithText("Spinner !").
		WithType(spinner.BoxBounce2).
		Build()

	spin.Start()
	time.Sleep(2 * time.Second)
	spin.Stop()
}
```
