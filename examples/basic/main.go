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
