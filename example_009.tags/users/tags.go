package users

import (
	"reflect"
	"strings"
	"strconv"
	"fmt"
)

const validate = "validate"

type validation_func func(reflect.StructField, reflect.Value, string) error


func _string_length_validator(field_def reflect.StructField,
			      field_value_def reflect.Value,
			      tag_definition string) error {
	limits := strings.Split(strings.Trim(tag_definition, "length:[]"), ",")
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
	} else {
		return fmt.Errorf("Unable to validate string literal " +
			"due to malformed tag %s", tag_definition)
	}
}


func _string_case_validator(field_def reflect.StructField,
			    field_value_def reflect.Value,
			    tag_definition string) error {

	fmt.Println("Attempting to validate case.")
	err_pattern := "Field `%s` value " +
		       "`%s` is not in `%s` case"
	case_def := strings.Split(tag_definition, ":")
	_v := field_value_def.String()
	if len(case_def) == 2 {
		switch case_def[1] {
		case "upper_only":
			fmt.Println("Checking if string is in UPPER case")
			if _v != strings.ToUpper(_v) {
				fmt.Println("Case validation was not successful.")
				return fmt.Errorf(err_pattern, field_def.Name, _v, "upper")
			}
			fmt.Println("Case validation was successful.")
			return nil
		case "lower_only":
			fmt.Println("Checking if string is in LOWER case")
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
}


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



func typed_validation(validate_tag string, typed_validation_map map[string]validation_func,
		      field_def reflect.StructField, field_value_def reflect.Value) error {
	fmt.Println(fmt.Sprintf("Attempting to validate typed field `%s`", field_def.Name))
	validators := strings.Split(validate_tag, ";")
	for key, function := range typed_validation_map {
		if strings.Contains(validate_tag, key) {
			fmt.Println(fmt.Sprintf("Tag found `%s` in `%s`", key, validate_tag))
			_underlying_tag := func() string {
				for i := 0; i < len(validators); i++ {
					if strings.Contains(validators[i], key) {
						return validators[i]
					}
				}
				return ""
			}()

			if _underlying_tag == "" {
				return fmt.Errorf("Malformed field `%s` tag(s)", validate_tag)
			}

			return function(field_def, field_value_def, _underlying_tag)
		}
	}
	return nil
}


func generic_validate(st interface{}) error {
	fmt.Println("Attempting to validate fields.")

	string_validation_map := map[string]validation_func{
		"length": _string_length_validator,
		"case": _string_case_validator,
	}

	int_validation_map := map[string]validation_func{
		"value": _int_value_validator,
	}

	value := reflect.ValueOf(st).Elem()
	_type := value.Type()
	for i := 0; i < value.NumField(); i++ {

		vField, tField := value.Field(i), _type.Field(i)
		validate_tag := tField.Tag.Get(validate)
		if validate_tag == "" {
			fmt.Println(fmt.Sprintf("No validation defined for field `%s`", tField.Name))
			continue
		}

		switch vField.Kind() {
		case reflect.String:
			err := typed_validation(validate_tag, string_validation_map, tField, vField)
			if err != nil {
				return err
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			err := typed_validation(validate_tag, int_validation_map, tField, vField)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (user *User) Validate() error {
	return generic_validate(user)
}
