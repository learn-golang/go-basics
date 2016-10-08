Go `panic` and `recover`
========================

Error handling is a good thing, but there are some cases when you can't handle error because it is not expected.
In Go there are two terms: `panic` and `recover`. 
First one is used to escalate error to top level as fast as possible, and `recover` helps to process `panic` calls.
But please note that `panic` and `recover` are not working as `try:except` in other languages, `panic` will cause program to exit with `ExitCode 2`.

Here's an example of both:

	func () {
		st, err = Hello{}.SayHello()
		if err != nil {
			panic(fmt.Sprintf("Can't proceed, aborint. Reason %s", err.Error()))
		}
	}()

we are calling `panic` because structure field is empty.

	defer func () {
		if er := recover(); er != nil {
			fmt.Println()
		}
	}()

By use of defer we can process `panic` result right after it appears, here's full example

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

As you can see if panic meant to be processed every defer (there are might be more than one) must be declared right after function definition and before actual function logic.
So, if we'd call this function it will print out
    
    Start Panic Defer
    I'm in recovery right now. In recovery: Something went wrong: Structure was not initialized.

As you can see right after `panic` we have another function call, but it wouldn't work because only `defer` with `recover` will be executed.

Here are interesting things:

* `panic` accepts type `interface{}`, so you can define custom type and send it to `recover` for further processing
* `defer` does not have a return value, so `defer` with `recover` is only allowing to continue program without `Exit code 2`
