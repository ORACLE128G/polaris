package pipline

func ArraySource(ints ... int) <-chan int {

	out := make(chan int, len(ints))

	go func() {
		for _, v := range ints {
			out <- v
		}
		close(out)
	}()
	return out
}
