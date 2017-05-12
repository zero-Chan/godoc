package json

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/zero-Chan/document"
)

type JsonTransCoder struct {
}

func CreateJsonTransCoder() JsonTransCoder {
	coder := JsonTransCoder{}
	return coder
}

func NewJsonTransCoder() *JsonTransCoder {
	coder := CreateJsonTransCoder()
	return &coder
}

//    To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
//       bool, for JSON booleans
//       float64, for JSON numbers
//       string, for JSON strings
//       []interface{}, for JSON arrays
//       map[string]interface{}, for JSON objects
//       nil for JSON null
// Unmarshal: make []byte -> doc
// []byte -> interface{} 		[ json.unmarshal ]
// interface{} -> doc	[ doc.marshal ]
func (coder *JsonTransCoder) Unmarshal(data []byte, v interface{}) (*document.Document, error) {
	err := json.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, fmt.Errorf("non-pointer(%s)", rv.Type().String())
	}

	vv := rv.Elem()
	switch vv.Kind() {
	case reflect.Map:
		objSec := document.NewObjectSection("")
		err := objSec.Marshal(vv.Interface())
		if err != nil {
			return nil, err
		}

		return document.NewDocument(objSec), nil

	case reflect.Struct:
		objSec := document.NewObjectSection("")
		err := objSec.Marshal(vv.Interface())
		if err != nil {
			return nil, err
		}

		return document.NewDocument(objSec), nil

	case reflect.Array:
		arrSec := document.NewArraySection("")
		err := arrSec.Marshal(vv.Interface())
		if err != nil {
			return nil, err
		}

	case reflect.String:
		strSec := document.NewStringSection("", vv.String())
		err := strSec.Marshal(vv.String())
		if err != nil {
			return nil, err
		}
		return document.NewDocument(strSec), nil

	case reflect.Bool:
		bSec := document.NewBoolSection("", vv.Bool())
		err := bSec.Marshal(vv.Bool())
		if err != nil {
			return nil, err
		}
		return document.NewDocument(bSec), nil

	default:
		return nil, fmt.Errorf("JsonTranCoder not support Unmarshal type[%s]", rv.Kind().String())
	}

	return document.NewDocument(document.NewNIlSection("")), nil
}

// Marshal: make doc to json format string
// doc -> interface{}	[ doc.unmarshal ]
// interface{} -> []byte	 [ json.marshal ]
func (coder *JsonTransCoder) Marshal(doc document.Document) ([]byte, error) {
	var v interface{}

	switch doc.Type() {
	case document.Object:
		objSec, err := doc.Object()
		if err != nil {
			return nil, err
		}

		mapv := make(map[string]interface{})
		if err = objSec.Unmarshal(&mapv); err != nil {
			return nil, err
		}

		v = mapv

	case document.Array:
		arrSec, err := doc.Array()
		if err != nil {
			return nil, err
		}

		arrv := make([]interface{}, 0)
		if err = arrSec.Unmarshal(&arrv); err != nil {
			return nil, err
		}

		v = arrv

	case document.String:
		strSec, err := doc.String()
		if err != nil {
			return nil, err
		}

		var strv string
		if err = strSec.Unmarshal(&strv); err != nil {
			return nil, err
		}

		v = strv

	case document.Bool:
		bSec, err := doc.Bool()
		if err != nil {
			return nil, err
		}

		var boolv bool
		if err = bSec.Unmarshal(&boolv); err != nil {
			return nil, err
		}

		v = boolv

	case document.Nil:
		nilSec, err := doc.Nil()
		if err != nil {
			return nil, err
		}

		var nilv interface{}
		if err = nilSec.Unmarshal(&nilv); err != nil {
			return nil, err
		}

		v = nilv

	case document.Int:
		numSec, err := doc.Number()
		if err != nil {
			return nil, err
		}

		var intv int
		if err = numSec.Unmarshal(&intv); err != nil {
			return nil, err
		}

		v = intv

	case document.Int8:
		numSec, err := doc.Number()
		if err != nil {
			return nil, err
		}
		var int8v int8
		if err = numSec.Unmarshal(&int8v); err != nil {
			return nil, err
		}

		v = int8v

	case document.Int16:
		numSec, err := doc.Number()
		if err != nil {
			return nil, err
		}

		var int16v int16
		if err = numSec.Unmarshal(&int16v); err != nil {
			return nil, err
		}

		v = int16v

	case document.Int32:
		numSec, err := doc.Number()
		if err != nil {
			return nil, err
		}

		var int32v int32
		if err = numSec.Unmarshal(&int32v); err != nil {
			return nil, err
		}

		v = int32v

	case document.Int64:
		numSec, err := doc.Number()
		if err != nil {
			return nil, err
		}

		var int64v int64
		if err = numSec.Unmarshal(&int64v); err != nil {
			return nil, err
		}

		v = int64v

	case document.Float32:
		numSec, err := doc.Number()
		if err != nil {
			return nil, err
		}

		var float32v float32
		if err = numSec.Unmarshal(&float32v); err != nil {
			return nil, err
		}

		v = float32v

	case document.Float64:
		numSec, err := doc.Number()
		if err != nil {
			return nil, err
		}

		var float64v float64
		if err = numSec.Unmarshal(&float64v); err != nil {
			return nil, err
		}

		v = float64v

	default:
		return nil, fmt.Errorf("JsonTranCoder not support Marshal type[%s]", doc.Type().String())
	}

	return json.Marshal(v)
}
