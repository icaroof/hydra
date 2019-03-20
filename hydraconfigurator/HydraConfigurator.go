package hydraconfigurator

import (
	"errors"
	"reflect"
)

const (
	CUSTOM int8 = iota
)

var wrongTypeError error = errors.New("Type must be a pointer to struct")

func GetConfiguration(confType int8, obj interface{}, filename string) (err error) {
	//check if the obj is of pointer type
	objRValue := reflect.ValueOf(obj)

	if objRValue.Kind() != reflect.Ptr || objRValue.IsNil() {
		return wrongTypeError
	}

	//get and confirm struct value
	objRValue = objRValue.Elem()

	if objRValue.Kind() != reflect.Struct {
		return wrongTypeError
	}

	switch confType {
	case CUSTOM:
		err = MarshalCustomConfig(objRValue, filename)
	}

	return
}
