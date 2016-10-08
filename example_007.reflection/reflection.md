Go reflection
=============

Definition
----------

Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming. It's also a great source of confusion.
So, basically, reflection helps to work object type and value. In Go, reflection is used in various frameworks and internal parts of Go itself.

Things that i didn't know it became some sort of a problems
-----------------------------------------------------------

The Go runtime doesn't exposes a list of types built in the program.
And there is a reason: you never have to build all types available, but instead just a subset.


Go reflect
----------

In Go standard library, there is a lib called `reflect` that provides an API to work with data type within runtime.
Most of articles regarding reflection would tell you to avoid use of reflections since it is not always easy.
And in most cases, according to those articles, if you're using reflection you probably building ORM framework.
But nevertheless, reflection is a feature that can be very useful while you need to work with type processing in runtime, similar to Python `type` function.

Reflection for type and value processing
----------------------------------------

In our example we have an interface that was designed to provide an API to work with type
```
type GenericUserTypeProcessor interface {
	GetType() reflect.Type
	GetFieldsByType(type_of reflect.Type) SetOfTypes
	GetAllFields() FieldsDef
}
```

where
```
type FieldRef map[string]interface{}
type FieldDef  map[string]FieldRef
type FieldsDef []FieldDef

type DataDef map[string]interface{}
type SetOfTypes []DataDef
```

Let's examine those methods. First of them is `GetType() reflect.Type`
```
func (type_processor *TypeProcessor) GetType() reflect.Type {
	return reflect.TypeOf(type_processor.Generic).Elem()
}
```
Note in case of pointer type we should retrieve the underlying actual type by calling `Elem` on object of `reflect.Type`.
So, this method actually retrieves a Type of an object by reflecting it.

Next is `GetFieldsByType(type_of reflect.Type) SetOfTypes`
```
func (type_processor *TypeProcessor) GetFieldsByType(type_of reflect.Type) SetOfTypes {
	list_of_data_defs := SetOfTypes{}
	for i := 0; i < type_processor.GetType().NumField(); i++ {
		item := make(DataDef)
		field := type_processor.GetType().Field(i)
		if field.Type == type_of {
			value := reflect.Indirect(reflect.ValueOf(
				type_processor.Generic)).FieldByName(field.Name)
			item[field.Name] = GetValue(value)
			list_of_data_defs = append(list_of_data_defs, item)
		}
	}
	return list_of_data_defs
}
```
Actual semantic of this method is - find every data field that type matches to `type_of`, so here what it does:
* loop over type data fields
* retrieve data field
* if type match, retrieve data field value and build a map {`name`:`value`}, append map to a list.
* otherwise, do nothing

In this method you may see certain complexity in data field value retrieving, that because we need to access value of an `type_processor.Generic` rather than its pointer.
And there's a difference between `reflect.Indirect` and `reflect.ValueOf(...).Elem()`, if the value is an interface then `reflect.Indirect(v)` will return the same value, while `reflect.ValueOf(...).Elem()` will return the contained dynamic value, and if the value is something else, then `v.Elem()` will panic.

It doesn't look complicated, right? But it does, and value reflection can cause pain, hare are couple examples.
What would this function print?
```
process(user, reflect.TypeOf(""))
```

```
Fields that are matching to search criteria [map[name:Denis] map[surname:Makogon]]
All data fileds of a type [map[name:map[type:string value:Denis]] map[surname:map[type:string value:Makogon]] map[id:map[type:int64 value:<int64 Value>]]]
```
So, as you can see, for our search criteria (lookup for every attribute which type is `string`) object `user` of a type `User` has two fields `name` and `surname` and their corresponding values are `Denis` and `Makogon`.
But if we'd take a look at second line, you will see that there's yet one data field of `int64` type and its value is not represented in regular way.

For most of built-in types like int(any its modification), float, bool, by use of reflection will give value object, but not value (like we've used to).
So, here's a tricky part of value reflection, for types listed above it is necessary to run type-specific function to retrieve a result, for example:

    int ---> v := reflect.ValueOf(int(100)) ---> v.Int()
    float ---> v := reflect.ValueOf(float(1.0)) ---> v.Float()
    bool ---> v := reflect.ValueOf(false) ---> v.Bool()

In our particular case we need to have a switch function that would retrieve value from its reflection

```
func GetValue(v reflect.Value) interface{} {
	switch v.Kind() {
		case reflect.Int, reflect.Int8,
			reflect.Int16, reflect.Int32, reflect.Int64:
			return v.Int()
		case reflect.Bool:
			return v.Bool()
		case reflect.String:
			return v.String()
	}
	return v
}
```

Apply this function wherever it is required with the assumption that incoming value type is not a complex type like structure, slices, etc.

And the last method is `GetAllFields() FieldsDef`
```
func(type_processor *TypeProcessor) GetAllFieldsDef() FieldsDef {
	fieldsdefs := FieldsDef{}
	for i := 0; i < type_processor.GetType().NumField(); i++ {
		def := make(FieldDef)
		field := type_processor.GetType().Field(i)
		field_name, field_type := field.Name, field.Type
		value := reflect.Indirect(reflect.ValueOf(
			type_processor.Generic)).FieldByName(field_name)
		def[field_name] = FieldRef{"type": field_type, "value": GetValue(value)}
		fieldsdefs = append(fieldsdefs, def)
	}
	return fieldsdefs
}
```
It is similar to previous one, but gets all fields without type filtering.

And let's review [main.go](main.go). Main funtion will print out line:

    Fields that are matching to search criteria [map[name:Denis] map[surname:Makogon]]
    All data fileds of a type [map[name:map[type:string value:Denis]] map[surname:map[type:string value:Makogon]] map[id:map[type:int64 value:1]]]
    Fields that are matching to search criteria [map[admin:true]]
    All data fileds of a type [map[User:map[type:users.User value:<users.User Value>]] map[admin:map[type:bool value:true]]]
