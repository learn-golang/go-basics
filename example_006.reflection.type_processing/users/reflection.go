package users

import (
	"reflect"
	"bytes"
	"fmt"
)

func (type_processor *TypeProcessor) GetType() reflect.Type {
	// Note in case of pointer type,
	// we should retrieve the underlying actual type
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
			item[field.Name] = GetValue(value, StructValueProcessor)
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
		def[field_name] = FieldRef{"type": field_type, "value": GetValue(value, StructValueProcessor)}
		fieldsdefs = append(fieldsdefs, def)
	}
	return fieldsdefs
}

func StructValueProcessor(v reflect.Value) interface{} {
	fieldsdefs := FieldsDef{}
	for i := 0; i < v.NumField(); i++ {
		def := make(FieldDef)
		field := v.Type().Field(i)
		value := reflect.Indirect(v).FieldByName(field.Name)
		def[field.Name] = FieldRef{"type": field.Type, "value": GetValue(value, StructValueProcessor)}
		fieldsdefs = append(fieldsdefs, def)
	}
	return fieldsdefs
}

func JSONStructValueProcessor(v reflect.Value) interface{} {
	json_ := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := reflect.Indirect(v).FieldByName(field.Name)
		json_[field.Name] = GetValue(value, StructValueProcessor)
	}
	return json_
}


func GetValue(v reflect.Value, value_processor func(v reflect.Value) interface{}) interface{} {
	switch v.Kind() {
		case reflect.Int, reflect.Int8,
			reflect.Int16, reflect.Int32, reflect.Int64:
			return v.Int()
		case reflect.Bool:
			return v.Bool()
		case reflect.String:
			return v.String()
		case reflect.Struct:
			return value_processor(v)
	}
	return nil
}

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
