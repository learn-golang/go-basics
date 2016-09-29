Go embedded interfaces
======================

Interface embedding
-------------------

Remind how Go allows to embed structures into structures, isn't cool, right?
And even more, Go is capable of doing different types of embedding, such as:

* structure into structure
* interface into interface
* interface into structure

But first of all let's understand how it works. Remind our initial interface - `UserInterface`.
It was implemented according to capabilities of `User` structure, but `AdminUser` structure extends it with a new data field - `is_admin bool`.
So, our interface doesn't actually covers whole set of capabilities from both types. In this case we have two options:

* modify `UserInterface` to have new function `IsAdmin`
* extend `UserInterface` in some way to have new function `IsAdmin`

In our example we would use 2nd option - extending (in Go terms - embedding). So, we would have a new interface `GenericUserInterface` that will have embedded interface

```
type GenericUserInfoInterface interface {
	UserInterface
	IsAdmin() bool
}
```

No difference with structure embedding, but you don't need to initialize interface.
Each new interfaces has all functions from embedded interfaces (probably such capability addresses OOP).

Initially we've defined function `IsAdmin` on `*AdminUser` type, and if we want to process both `*User` and `*AdminUser` in the same manner we need to have the same function on `*User` type.

```
func (user *User) IsAdmin() bool {
	return false
}
```

So, starting this point we can process both `User` and `AdminUser` in two ways - as `UserInterface` dictates or as `GenericUserInterface` dictates.


Anonymous interfaces
--------------------

In Go it is possible to embed an interface into a structure and the main purpose for such capability is to provide an ability to initialize a structure with any data type that implements anonymous interface.
In our particular example we have `GenericUserInterface` interfaces
```
type GenericUserInfoInterface interface {
	UserInterface
	IsAdmin() bool
}
```
that is being embedded into `GenericUser` structure
```
type GenericUser struct {
	GenericUserInfoInterface
}
```

So, from `GenericUser` structure interface `GenericUserInterface` is an embedded interface.
Since both `User` and `AdminUser` are implementing that interface so we can initialize `GenericUser` with both types.
In order to demonstrate that we need to have a function that talks to an object that implements an interface
```
func (generic *GenericUser) GetUserInfo() (string, string, int64, bool) {
	return generic.GetName(), generic.GetSurname(), generic.GetID(), generic.IsAdmin()
}
```

So, now let's see how it works
```
admin := new(users.AdminUser)
admin = admin.Create(name, surname)
generic := users.GenericUser{admin}
fmt.Println(generic.GetUserInfo())
```

As already told, structure `GenericUser` accepts any other structure that implements anonymous interface.
As the result you'll see

```
Denis Makogon 2 true
```

It's it awesome?
