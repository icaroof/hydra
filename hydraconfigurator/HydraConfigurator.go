package hydraconfigurator

import (
	"errors"
	"reflect"
)

//Configuration file types
const (
	CUSTOM int8 = iota
	JSON
	XML
)

//error to return if object is not a pointer to a struct
var errWrongType = errors.New("Type must be a pointer to struct")

//GetConfiguration reads the supplied file then fills the supplied struct with configuration parameters
func GetConfiguration(confType int8, obj interface{}, filename string) (err error) {
	//check if the obj is of pointer type
	objRValue := reflect.ValueOf(obj)

	if objRValue.Kind() != reflect.Ptr || objRValue.IsNil() {
		return errWrongType
	}

	//get and confirm struct value
	objRValue = objRValue.Elem()

	if objRValue.Kind() != reflect.Struct {
		return errWrongType
	}

	switch confType {
	case CUSTOM:
		err = marshalCustomConfig(objRValue, filename)
	case JSON:
		err = decodeJSONConfig(obj, filename)
	case XML:
		err = decodeXMLConfig(obj, filename)
	}

	return
}
