Go reflection
=============

Type creation
-------------

Here's the thing, let's go back to where we end and see what we had:

    Fields that are matching to search criteria [map[name:Denis] map[surname:Makogon]]
    All data fileds of a type [map[name:map[type:string value:Denis]] map[surname:map[type:string value:Makogon]] map[id:map[type:int64 value:1]]]
    Fields that are matching to search criteria [map[admin:true]]
    All data fileds of a type [map[User:map[type:users.User value:<users.User Value>]] map[admin:map[type:bool value:true]]]

As you can see one of returned fields has type `users.User` and its value was not interpreted in correct way. But we'd like to get its representation similar to other fields rather than string representation of `reflect.Value` struct.
For this we need to modify our `GetValue` function

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
		case reflect.Struct:
			new_v_struct := reflect.New(v.Type()).Elem()
			for i := 0; i < v.NumField(); i++ {
				field := v.Type().Field(i)
				fmt.Println(fmt.Sprintf(
					"Attempting to set field `%s` to structure `%s`",
					field.Name, new_v_struct.Type()))
				fmt.Println(new_v_struct.Type())
				value := reflect.Indirect(v).FieldByName(field.Name)
				if !new_v_struct.FieldByName(field.Name).CanSet() {
					panic(fmt.Sprintf(
						"It may appear that field `%s`" +
						" in lower case, so it is read-only mode on structure `%s`",
						field.Name, new_v_struct.Type()))
				}
				new_v_struct.FieldByName(field.Name).Set(value)
			}
			return new_v_struct
	}
	return nil
}
```
Each time this function will be called it will panic (in our particular use case) because each field of `User` structure is lowercase so Go thinks that those fields are not exportable.
So, the thing is type can't be initialized properly if its fields are lowercase, it is necessary to have a workaround, and that's what i did.

Since our goal is to convert an object to a list map of its fields (type and value) we just need to use proper `GetValue` function
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
		case reflect.Struct:
			fieldsdefs := FieldsDef{}
			for i := 0; i < v.NumField(); i++ {
				def := make(FieldDef)
				field := v.Type().Field(i)
				value := reflect.Indirect(v).FieldByName(field.Name)
				def[field.Name] = FieldRef{"type": field.Type, "value": GetValue(value)}
				fieldsdefs = append(fieldsdefs, def)
			}
			return fieldsdefs
	}
	return nil
}
```
So, this function will react on each structure by recursively decomposing it into desired format.

For this lesson we'd amend attempts to create new structures `User` and `AdminUser` due to its complexity.
But the main reason of showing how object gets converted into list of maps or map, or whatever it will be is to should how Go can work with type creation.


Think of JSON/XML
-----------------

Since you are aware what is necessary to do for type decomposition, how would you implement convertion to a mapping?

For such use case we need to transform `GetValue` from

    func GetValue(v reflect.Value) interface{};
to

    func GetValue(v reflect.Value, value_processor func(v reflect.Value) interface{}) interface{};

`value_processor` stand for transforming structure to a mapping (code listed below)

    func (type_processor *TypeProcessor) ToMap() map[string]interface{} {
        map_ := map[string]interface{}{}
        for i := 0; i < type_processor.GetType().NumField(); i++ {
            field := type_processor.GetType().Field(i)
            value := reflect.Indirect(reflect.ValueOf(
                type_processor.Generic)).FieldByName(field.Name)
            map_[field.Name] = GetValue(value, JSONStructValueProcessor)
        }
        return map_
    }

So, it collects all values recursively by decomposing complex types like structure to their primitive types.


Think for SQL`ing
-----------------

What do i need to extract SQL query from structure definition? Easy, collect fields, us structure as table representation

    func (type_processor *TypeProcessor) ToSQL() string {
        buffer := bytes.NewBufferString("")
        for i := 0; i < type_processor.GetType().NumField(); i++ {
            field := type_processor.GetType().Field(i)
            if i == 0 {
                buffer.WriteString("SELECT ")
            } else {
                buffer.WriteString(", ")
            }
            buffer.WriteString(field.Name)
        }
        if buffer.Len() > 0 {
            fmt.Fprintf(buffer, " FROM %s;", type_processor.GetType().Name())
        }
        return buffer.String()
    }

Looks good, easy and useful. But, stop, again. Remember how reflection represents embedded structures? Now take a look this code

    admin := new(users.AdminUser)
    admin = admin.Create(name, surname)
    type_processor := users.TypeProcessor{Generic: admin}
    type_processor.ToSQL()

Output will have next value:

    SQL query: SELECT User, admin FROM AdminUser;
    
Uugh, again, embedded `User` structure. How to omit this one? Or how to expand it to useful sub-query?
My answer here - TAG THEM ALL!
