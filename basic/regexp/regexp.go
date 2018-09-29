package main

import (
	"fmt"
	"regexp"
)

const text = `1302 ****123 asdas oracle128g@gmail.com
  123@qq.com  w@qq.com  w@qw.com.cn`
func main() {
	compile, _ := regexp.Compile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	match := compile.FindAllString(text, -1)
	fmt.Println(match)
}
