package main

import (
	"github.com/ltoddy/spinner"
	"time"
)

func main() {
	spin := spinner.NewSpinnerBuilder().
		WithText("Spinner !").
		WithType(spinner.Dots12).
		Build()

	spin.Start()
	time.Sleep(10 * time.Second)
	spin.Stop()
}
