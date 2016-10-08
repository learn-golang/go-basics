Go embedding
============

Embedding
---------
Unlikely to other object-oriented languages (like Python, Java, C++), Go doesn't have explicit inheritance model. So, in Go it was decided to implement embedding via shadowed fields.
Let's take a look at `AdminUser` structure from [User](users/data_types.go).

    type AdminUser struct {
        User
        is_admin bool
    }

There's key thing, `User` structure was embedded into `AdminUser` structure.

So, embedding is a kind of inheritance, it also allows to access embedded structure data fields and every function defined on embedded structure no matter via its value or reference.
Coming from `User` structure functions definitions:

    func (user *User) GetName() string;
    func (user *User) GetSurname() string;
    func (user *User) GetID() string;

all this methods are available in `AdminUser` structure and can be accessed via dot-notation

    user := new(users.User)
    user = user.Create(name, surname)
    fmt.Println("Admin name: ", admin.GetName())
	fmt.Println("Admin surname: ", admin.GetSurname())
	fmt.Println("Admin ID: ", admin.GetID())

Along with that, `AdminUser` can have its own methods, in our particular case

    func (admin *AdminUser) IsAdmin() bool;

It is obvious that embedding works only in one direction and it is not recommended to create circular embeddings.

Standard composition
--------------------

It appears that reference to `AdminUser` can't be used for adding it to a users list because `AdminUser` type is not similar to `User` type and it's obvious.
So, initially implemented method

    func (users *Users) Append(user *User);

has incorrect parameter type because it is limited to `User` type only. So what are we need to do in order to accomplish required goal?
Answer is not really obvious or simple - we need to rewrite code, but how?
Answers are:

    * rewrite AdminUser structure to have a dedicated data field which declaration follows standard composition
    * rewrite Append function to accept more generic type rather than User

And more generic type here is [Go interface](../example_005.interfaces/interfaces.md)!
