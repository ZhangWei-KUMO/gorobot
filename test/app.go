package main

import (
	"fmt"

	"github.com/ZhangWei-KUMO/gorobot"
)

func main() {
	message := gorobot.Hello("zhangwei")
	fmt.Println(message)
}
