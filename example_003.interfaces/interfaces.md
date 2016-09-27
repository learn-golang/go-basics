Go interfaces
=============

Embedding
---------

Having a capability to embed structures into other structures to mimic inheritance allows to process different types with defined API in base types.
But the thing is, Go is Python and duck typing does not work here.

In our particular case, what is the way to implement generic data type API to process both `User` and `AdminUser` structures?


Interfaces
----------

Since we already have `User` structure that defines all necessary API like:

    GetName() string;
    GetSurname() string;
    GetID() int64;

We can use its definition to create an interfaces `UserInterface`:

    type UserInterface interface {
        GetName() string
        GetSurname() string
        GetID() int64
        SetID(new_id int64)
    }

Actual differences between `User` (or `AdminUser`) and `UserInterface` are:

 * interface doesn't have data fields
 * interface does have function declaration
 * interface does have interface embedding

As it can be seen, interface and structure declarations are slightly different and this difference comes from interface disability to hold data field, so it was necessary to have a function that will be able to modify data field of actual structure that follows corresponding interface.

Now let's go back to our example. [Here](users/interfaces.go) you can find interface type declaration.
Actual differences between [Example 2](../example_002.embedding/main.go) and [Example 3](main.go) is hidden in `processor.go` files.

User list modifications
-----------------------

First of all, function `func (users *Users) Append(user UserInterface) error;` was modified in order to handle generic type
    
```
func (users *Users) Append(user UserInterface) error {
	users.index += 1
	switch v := user.(type) {
		case *User, *AdminUser:
			add(users, v)
			return nil
		default:
			return errors.New("Unknown type")
	}
}
```
As you can see, now `Append` accepts parameters of a type `UserInterface` interface.
And here comes magic of interface processing, in Go it is called - **type assertion**.

Type assertion
--------------

In Go, type assertion helps to identify a type of an object while it is being processed via interface.
In order to identify an object type Go provides next capability:

```
obj, ok := obj.(t Type)
```

But this code works only if it is known for sure that, through this interface, object has **only one type**.
In other cases when object can be a multi-type type (because of embedding or mora than one structure follows defined interface) Go documentation suggests to use `switch-case` syntax to identify type correctly.

So, now our `Append` function can work with interface instead of exact type and since our code has two entities: `User` and `AdminUser` it was necessary to use `switch-case`.
Remarkable thing here is that interface does not really care if you pass object by its value or its reference, it is all about how would you define cases.
Since we use only references to work with objects than our cases are: `*User` and `*AdminUser`.
Along with exact cases `switch-case` syntax provides an ability to define `default case` for those situations when none of other cases worked.

	switch v := user.(type) {
		case *User, *AdminUser:
			add(users, v)
			return nil
		default:
			return errors.New("Unknown type")
	}

For our own case, `default case` would be an error since we expect only two exact types.

And the last changes in our code:

* a type of user holder - from `[]*Users` to `[]UserInterface`
* a return value of `GetUsers()` function - from `[]*Users` to `[]UserInterface`
