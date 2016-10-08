package users

import (
	"reflect"
)

func (type_processor *TypeProcessor) GetType() reflect.Type {
	return reflect.TypeOf(type_processor.Generic).Elem()
}

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
