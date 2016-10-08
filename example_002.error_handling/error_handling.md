Go error handling
=================

Error handling in Go is quite tricky thing, since Go follows pure OOP it doesn't have exception throwing mechanism.
Go idiom is in returning error explicitly wherever it is necessary.
So, we have three type exception handling:

* Sentinel errors
* Error types
* Opaque errors

Sentinel errors
---------------

The first category of error handling is sentinel errors:

    if err == SomeErr {...}

Using sentinel values is the least flexible error handling strategy, as the caller must compare the result to predeclared value using the equality operator.
So here are some points that must be kept in mind:

* Never inspect the output of error.Error (the Error method on the error interface exists for humans, not code)
* Sentinel errors become part of your public API (if your public function or method returns an error of a particular value then that value must be public, and of course documented)
* Sentinel errors create a dependency between packages (as an example, to check if an error is equal to `io.EOF`, your code must import the `io` package)

If someone asks you to export an error value from your package, you should politely decline and instead suggest an alternative method, such as the ones I will discuss next.

Error types
-----------

What makes type an error type? Basically, Go has an interface for errors, so each type that is considered to be an error need to implement interface `error`

    type MyError struct {
        Msg string
    }

least minimum implementation is

    func (e *MyError) Error() string {
        return fmt.Sprintf("%s", e.Msg)
    }


So, now let's adopt this to our example, we need to adjust `SayHello` to `func (h Hello) SayHello() (string, error)`
    
    func (h Hello) SayHello() (string, error) {
            if h.HelloString == "" {
                return "", &MyError{Msg:"Structure was not initialized."}
            }
            return h.HelloString, nil
    }

Now we have multiple return values:

* string
* error

Small nit here is if your error structure type (`MyError`) has only one data field of a type `string` it can be substituted with

    errors.New("Structure was not initialized.")

instead

    &MyError{Msg:"Structure was not initialized."}

It is recommended to process an errors by their types rather than values

    if err, ok := err.(SomeErrType); ok {...}

Because `MyError` error is a type, callers can use type assertion (in two ways: one described above and other one through `switch`)to extract the extra context from the error.

Don’t just check or transform errors
------------------------------------

Very good example was brought by [Dave Cheney](http://dave.cheney.net/), Donovan and Kernighan’s The Go Programming Language recommends that you add context to the error path using `fmt.Errorf`

    func AuthenticateRequest(r *Request) error {
            err := authenticate(r.User)
            if err != nil {
                    return fmt.Errorf("authenticate failed: %v", err)
            }
            return nil
    }

This pattern is incompatible with the use of sentinel error values or type assertions, because converting the error value to a string, merging it with another string, then converting it back to an error with `fmt.Errorf` breaks equality and destroys any context in the original error.


Annotating errors
-----------------

This pattern is quite good one because it doesn't exchange context of errors like we've seen before but extends it, for such thing there a lib [github.com/pkg/erros](https://godoc.org/github.com/pkg/errors)
That package has a lot useful APIs to work with errors, but i'd like to outline for now only two of them:

* `func Wrap(cause error, message string) error`
* `func Cause(err error) error`

So, it is quite easy to work find out what are those for:

* `Wraps` takes an error and new error string and wraps previous error with a new one by adding a new string to that
* `Cause` does the reverse operation to `Wraps` and checks if an error was wrapped or not

By use of origin context of an error is always stays as is.

Conclusion
----------

For maximum flexibility I recommend that you try to treat all errors as opaque.
In the situations where you cannot do that, assert errors for behaviour, not type or value.

