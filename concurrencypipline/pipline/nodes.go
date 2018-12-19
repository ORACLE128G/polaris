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

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2

		for ok1 || ok2 {
			// v1 has value
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				// continue take out a next value
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out

}
