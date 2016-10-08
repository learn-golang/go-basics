package main

import (
	"fmt"
)

type  Hello struct {
	HelloString string
}

func SayHello() string {
	return "Hello.\n"
}

func (h Hello) SayHello() string {
	return h.HelloString
}

func main() {
	fmt.Println(SayHello())

	h := Hello{HelloString:"Hello.\n"}
	fmt.Println(h.SayHello())

	fn := func (say_hello func() string) string {
		return  say_hello()
	}
	fmt.Println(fn(SayHello))
}
