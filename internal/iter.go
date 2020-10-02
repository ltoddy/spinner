package internal

func Repeat(times int, s string) []string {
	xs := make([]string, 0)
	for i := 0; i < times; i++ {
		xs = append(xs, s)
	}
	return xs
}

func Cycle(xs []string) <-chan string {
	ch := make(chan string, 16)

	go func() {
		for {
			for _, x := range xs {
				ch <- x
			}
		}
	}()

	return ch
}
