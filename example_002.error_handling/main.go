package main

import (
	"fmt"
	"github.com/pkg/errors"
)


type MyError struct {
	Msg string
}


func (e *MyError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}


type  Hello struct {
	HelloString string
}

func SayHello() string {
	return "Hello.\n"
}

func (h Hello) SayHello() (string, error) {
	if h.HelloString == "" {
		return "", errors.Wrap(&MyError{Msg:"Structure was not initialized."}, "Something went wrong")
	}
	return h.HelloString, nil
}

func DoPanicAndRecover(initialize bool) {
	defer func () {
		if r := recover(); r != nil {
			fmt.Println("Start Panic Defer")
			fmt.Println("I'm in recovery right now.", r)
		}
	}()

	h := Hello{}
	if initialize {
		h.HelloString = "Hello"
	}
	_, err := h.SayHello()
	if err != nil {
		panic(errors.Wrap(err, "In recovery"))
		fmt.Println("This will never run.")
	}
}

func main() {
	fmt.Println(SayHello())

	st, err := Hello{}.SayHello()
	fmt.Println("Result? - ", st)
	fmt.Println("Error? - ", err)

	h := Hello{HelloString:"Hello.\n"}
	st, err = h.SayHello()
	fmt.Println("Result? - ", st)
	fmt.Println("Error? - ", err)

	fn := func (say_hello func() string) string {
		return  say_hello()
	}
	fmt.Println(fn(SayHello))

	DoPanicAndRecover(true)
	fmt.Println("This will run.\n")
	DoPanicAndRecover(false)
}
