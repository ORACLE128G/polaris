package pipline

import "sort"

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

func InMemSort(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		var all []int

		for v := range in {
			all = append(all, v)
		}

		// Sort
		sort.Ints(all)

		// Output
		for _, v := range all {
			out <- v
		}

		close(out)

	}()
	return out

}
