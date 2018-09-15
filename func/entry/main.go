package main

import "polaris/func"

// Adder Main func.
/*func main() {
	adder := _func.Adder()
	for i := 0; i < 100; i++ {
		fmt.Printf("0 + 1 + ...%d = %d \n", i, adder(i))
	}

	bot := _func.AdderBot(0)
	for i := 0; i < 100; i++ {
		// Initial AdderBot
		var s int
		s, bot = bot(i)
		fmt.Printf("0 + 1 + ...%d = %d \n", i, s)

	}
}*/

func main () {

	init := _func.AutoInit()
	_func.Travers(init)
}
