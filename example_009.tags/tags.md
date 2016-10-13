Go structure tags
=================

Tags
----

So, Go tags are very useful feature of structure processing. Almost all encoding frameworks are designed to use tags to convert structures to XML or JSON, or whatever.
Basically, tag is nothing else that helper of metadata to actual field.
In order to access tags it is necessary to use `reflect` to access tag definitions.
Let's a bit transform our `User` structure:
    
    type User struct {
        name string `validate:"length:[0,255];case:upper_only"`
        surname string `validate:"length:[0,255];case:upper_only"`
        id int64 `validate:"value:zero_only"`
    }

From `reflect` lib

    By convention, tag strings are a concatenation of optionally 
    space-separated key:"value" pairs. Each key is a non-empty string consisting 
    of non-control characters other than space (U+0020 ' '), quote (U+0022 '"'), 
    and colon (U+003A ':'). Each value is quoted using U+0022 '"' 
    characters and Go string literal syntax.


Value validation
----------------

In our particular case we have tag `validate`, it stands for value validation (obviously) of length and case for `string` fields and zero equity for `int` string.
So, let's start with our code that reads tags and runs validation


    func (user *User) Validate() error {
        fmt.Println("Attempting to validate fields.")
        
First part is - `string` validation map

        string_validation_map := map[string]validation_func{
            "length": _string_length_validator,
            "case": _string_case_validator,
        }
        int_validation_map := map[string]validation_func{
            "value": _int_value_validator,
        }

This map contains every validation methods pinned to its tag field, similar for `int` validation map

        value := reflect.ValueOf(user).Elem()
        _type := value.Type()
        for i := 0; i < value.NumField(); i++ {
    
            vField, tField := value.Field(i), _type.Field(i)

From `StructField` definition (from package `reflect`) each field has attribute `Tag`

            validate_tag := tField.Tag.Get(validate)
            if validate_tag == "" {
                fmt.Println(fmt.Sprintf("No validation defined for field `%s`", tField.Name))
                continue
            }

If no tags defined than no validation will happen

            switch vField.Kind() {
            case reflect.String:

For each `String` field apply validation according to given tags

                err := typed_validation(validate_tag, string_validation_map, tField, vField)
                if err != nil {
                    return err
                }
            case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

For every kind of `Int` field apply validation according to given tags

                err := typed_validation(validate_tag, int_validation_map, tField, vField)
                if err != nil {
                    return err
                }
            }
        }
In our particular case we have only two data types: `String` and kinds of `Int`

        return nil
    }

Now let's take a look at function `typed_validation`

    func typed_validation(validate_tag string, typed_validation_map map[string]validation_func,
                  field_def reflect.StructField, field_value_def reflect.Value) error {
        fmt.Println(fmt.Sprintf("Attempting to validate typed field `%s`", field_def.Name))
        validators := strings.Split(validate_tag, ";")

In validation tag we use `;` to split nested tags

        for key, function := range typed_validation_map {
            if strings.Contains(validate_tag, key) {
            
If provided tag is in validation mapping than proceed to its execution

                fmt.Println(fmt.Sprintf("Tag found `%s` in `%s`", key, validate_tag))
                _underlying_tag := func() string {
                    for i := 0; i < len(validators); i++ {
                        if strings.Contains(validators[i], key) {
                            return validators[i]
                        }
                    }
                    return ""
                }()

Extracting nested tag, for example `case:upper_only` from `validation:"case:upper_only"`
    
                if _underlying_tag == "" {
                    return fmt.Errorf("Malformed field `%s` tag(s)", validate_tag)
                }

If tag there but it was not extracted properly, it map appear that tag is malformed.

                return function(field_def, field_value_def, _underlying_tag)
            }
        }
        return nil
    }


String validations
------------------


    func _string_length_validator(field_def reflect.StructField,
                                  field_value_def reflect.Value,
                                  tag_definition string) error {
        limits := strings.Split(strings.Trim(tag_definition, "length:[]"), ",")

Length range are defined by two limits: upper and lower (`length:[0,255]`)

        if len(limits) == 2 {
            _v := field_value_def.String()
            lower, err := strconv.Atoi(limits[0])
            if err != nil {
                return err
            }
            upper, err := strconv.Atoi(limits[1])
            if err != nil {
                return err
            }
            if lower < len(_v) && len(_v) <= upper {
                fmt.Println("LEN validation successful")
                return nil
            } else {
                return fmt.Errorf("string `%s` liternal length " +
                    "is not valid for field %s", _v, field_def.Name)
            }

Check if given string value meets tag-defined limits

        } else {
            return fmt.Errorf("Unable to validate string literal " +
                "due to malformed tag %s", tag_definition)
        }
    }

If range was malformed, than we need to abort validation with corresponding error.
 

    func _string_case_validator(field_def reflect.StructField,
                                field_value_def reflect.Value,
                                tag_definition string) error {
    
        fmt.Println("Attempting to validate case.")
        err_pattern := "Field `%s` value " +
                   "`%s` is not in `%s` case"
        case_def := strings.Split(tag_definition, ":")

Attempting to split `case:upper_only` into `[case upper_only]`

        _v := field_value_def.String()
        if len(case_def) == 2 {
            switch case_def[1] {
            case "upper_only":
                fmt.Println("Checking if string is in UPPER case")

Attempting to check if string value is in upper case

                if _v != strings.ToUpper(_v) {
                    fmt.Println("Case validation was not successful.")
                    return fmt.Errorf(err_pattern, field_def.Name, _v, "upper")
                }
                fmt.Println("Case validation was successful.")
                return nil
            case "lower_only":
                fmt.Println("Checking if string is in LOWER case")

Attempting to check if string value is in lower case

                if strings.ToLower(_v) != _v {
                    fmt.Println("Case validation was not successful.")
                    return fmt.Errorf(err_pattern, field_def.Name, _v, "lower")
                }
                fmt.Println("Case validation was successful.")
                return nil
            case "-":
                return nil
            }
        }
        return fmt.Errorf("Unable to validate string literal " +
                  "due to malformed tag %s", tag_definition)

If `case` nested tag is malformed - abort.

    }


Int validation
--------------

Similar to previous checks, in this validation we need to figure out if `int` field value is truly equal to zero.

    func _int_value_validator(field_def reflect.StructField,
                  field_value_def reflect.Value,
                  tag_definition string) error {
        fmt.Println(fmt.Sprintf("Attempting to validate integer field `%s`", field_def.Name))
        _v := field_value_def.Int()
        switch strings.Split(tag_definition, ":")[1] {
        case "zero_only":
            if _v != 0 {
                return fmt.Errorf("Only ZERO value allowed for " +
                    "int field `%s`, actual `%s`", field_def.Name, _v)
            }
            fmt.Println("Validation was successful.")
            return nil
        }
        return nil
    }

In-depth validation
-------------------

Here's the thing, `User` struct is simple one, because it has native types with no embedding. But `AdminUser` is more complex because it has embedding.
So, in order to keep simplicity we of value validation i'd recommend to do the same thing as i did for `AdminUser` creation

    func (admin *AdminUser) Create(name, surname string) (*AdminUser, error) {
        _new_user := &User{name: name, surname:surname}
        err := _new_user.Validate()
        if err != nil {
            return nil, err
        } else {
            return &AdminUser{User:_new_user, admin:true}, nil
        }
    }

No need to have `Validation` function implemented for type `*AdminUser` but have it for those structures that are being embedded, that would allow you to avoid need to recursive traversing through N fields and if this field embedded or structure see if it can be validated.