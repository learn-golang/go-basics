Go structures
=============


Structures and functions
------------------------
Let's talk about Go structures.
In Go structures are similar to concept of classes in Python, but by itself structure holds only its data fields.
So, subsequently structure does not hold methods in its declaration.

Let's take a look at structure [User](users/data_types.go):

    type User struct {
        name string
        surname string
        id int64
    }

This type describes `User` structure type with fields:

    name
    surname
    id

And each of this data field declarations holds its type, for example:

    name is a string holder
    surname is a string holder
    id is an int32 holder


So, basically, this structure holds simple data types.
But [Users](users/data_types.go) structure 

    type Users struct {
        users []*User
        index int64
    }

holds custom type - a list of references to multiple User structures.
Eventually it may be easy to use simpler type declaration - a list of User structures by their values.
Main purpose of such definition  - memory consumption, it is always easier 
to process entities by their references rather than values.


Probably you wonder how those structures are can be useful without defining methods for them.

Go provides an ability to create a functions that are pinned to a specific type:

    func [(name Type)] funcName

This type of functions are working in exact the same as class object instance methods in Python.
But there's couple differences, one of them is - functions can be called on structure object or on its reference.

    func (user *User) GetName()
    func (user User) GetName()  

The difference between this two is that for non-pointer functions its receives a copy of an object and every modification will be local-only, 
but pointer functions actually receives an address to an object and every modification will take its place.


Visibility
----------

Assume defined structure has number of functions on it, in Go terms it is possible to distinguish two types of methods and attributes:

    visible
    shadowed

In Go distinguishing factor can be accomplished by letter case. So, here are two rules:

 * If structure data field or function name starts with lower-case it is defined by compiler as private.
   Visibility scope of such fields and functions is limited to public methods.

 * If structure data fields or functions name starts with upper-case it is defined by compiler as public.
   Visibility scope of such fields and function is not limited.

Let's go back to our [Users](users/data_types.go) structure

    type User struct {
        name string
        surname string
        id int64
    }

If developer would consider to define user structure within main function

    name, surname := "Denis", "Makogon"
    user := new(users.User)
    users.User{name: name, surname:surname}

compiler will raise an error

    unknown users.User field 'name' in struct literal
    unknown users.User field 'surname' in struct literal

So, that's why it is necessary to have some sort of constructors similar to Python classes, that's why

    func (user *User) Create(name, surname string) *User;

was defined of the sake of creating an object instance
