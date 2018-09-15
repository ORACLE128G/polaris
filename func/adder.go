package _func

func Adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func AdderBot(base int) iAdder {
	return func (v int) (int, iAdder) {
		return base + v, AdderBot(base + v)
	}
}
