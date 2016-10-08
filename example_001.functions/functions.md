Go functions
============

Minimal unit of operations in Go is function.
Each function is defined as:

    func [(t Type)] name ([args...Type]) [Type] {};

Let's take a look at function definition parts.

func
----

    func - is reserved word for function declaration
    
type pin declaration
--------------------

    [(t Type)] - defines weather function defined as single unit or it is pinned to a specific type value or type value reference

name
----

    name - function name (here it is necessary to be very careful because Go has terms of exported and unexported objects)

parameters
----------

    [args...Type] - parameters definition (in Go you can pass through variadic number of parameters using `...`, but you it is necessary to bind those parameters to a specific type)
 
return value
------------
 
    [Type] - weather function has a return statement or not of a specific type

body
----

    {} - function body


Types of functions
------------------

As i already mentioned Go defines three types of function:

* single unit
* anonymous
* type pinned


Single unit functions
---------------------

Here's a simple example of function that prints `Hello.\n`
    
    func SayHello() string {
        return "Hello.\n"
    }

We have function that is not pinned to a type, with name `SayHello` and return statement of a type `string`.


Type pinned functions
---------------------

In order to bind a function to a specific type we need to define this type:

    type  Hello struct {
        HelloString string
    }

This is just a regular structure with `HelloString` data field of a type `string`, and now we're capable to bind function to this type

    func (h Hello) SayHello() string {
            return h.HelloString
    }

As you can see, we have name collision here, but Go is smart enough to distinguish this functions `SayHello` and `(h Hello) SayHello` as two different function.

Anonymous functions
-------------------

If you are familiar with Python you'll probably aware of `lambda` function, in functional programming they are called `in-place` function. Anonymous means that this function don't have names and can be accessed only within its definition scope.
Syntax of anonymous function in Go:

    func ([args...Type]) [Type] {};

So, as you can see, function does not have name or type binding, but it has its parameters, return value, body.
Our particular example is:

	fn := func (say_hello func() string) string {
		return  say_hello()
	}

we have function that accepts function with return statement of a type `string` as parameter and returns value of a type `string`.
In scope of it definition we have variable `fn` with will store value of a type `func(func() string)`, so now we just need to call this function with specific parameters (we'll use `SayHello` function).

    fmt.Println(fn(SayHello))



So, this is what are Golang functions, [code you find here](functions.go).