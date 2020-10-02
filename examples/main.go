package main

import (
	"fmt"
	"github.com/ltoddy/spinner"
	"time"
)

func main() {
	now := time.Now()
	spin := spinner.NewSpinner()
	spin.Start()
	time.Sleep(time.Second * 2)
	spin.Stop()
	fmt.Println("Hello world")
	fmt.Println(time.Since(now))
}
