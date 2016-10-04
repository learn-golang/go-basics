package users

import "reflect"

type FieldRef map[string]interface{}
type FieldDef  map[string]FieldRef
type FieldsDef []FieldDef

type DataDef map[string]interface{}
type SetOfTypes []DataDef

type UserInterface interface {
	GetName() string
	GetSurname() string
	GetID() int64
	SetID(new_id int64)
}

type GenericUserInfoInterface interface {
	UserInterface
	IsAdmin() bool
}

type GenericUserTypeProcessor interface {
	GetType() reflect.Type
	GetFieldsByType(type_of reflect.Type) SetOfTypes
	GetAllFields() FieldsDef
}
