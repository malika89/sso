package utils

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"reflect"
)

func ParseField(v reflect.Value) (keyList, valList []string, tbn string) {
	for i := 0; i < v.Type().NumField(); i++ {
		tableName := v.Type().Field(i).Tag.Get("tbn")
		if tableName != "" {
			tbn = tableName
		}
		tag := v.Type().Field(i).Tag.Get("json")
		if tag != "" {
			vl := parseNil(v.Field(i))
			if len(vl) != 0 {
				keyList = append(keyList, fmt.Sprintf(`"%s"`, tag))
				valList = append(valList, vl...)
			}
		}
	}
	return
}

func parseNil(v reflect.Value) (valList []string) {
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		if v.Len() != 0 {
			valList = parseType(v)
		}
	case reflect.Ptr:
		if !v.IsNil() {
			valList = parseType(v.Elem())
		}
	case reflect.Struct:
		if reflect.New(v.Type()).Elem().Interface() != v.Interface() {
			valList = parseType(v)
		}
	case reflect.Bool:
		valList = parseType(v)
	default:
		if v.Interface() != reflect.Zero(v.Type()).Interface() {
			valList = parseType(v)
		}
	}
	return
}

func parseType(v reflect.Value) (valList []string) {
	switch v.Type() {
	case reflect.TypeOf(sql.NullString{}):
		v1 := v.Interface()
		v2 := v1.(sql.NullString)
		valList = append(valList, fmt.Sprintf(`'%v'`, v2.String))
	case reflect.TypeOf(sql.NullFloat64{}):
		v1 := v.Interface()
		v2 := v1.(sql.NullFloat64)
		valList = append(valList, fmt.Sprintf(`'%v'`, v2.Float64))
	case reflect.TypeOf(sql.NullInt64{}):
		v1 := v.Interface()
		v2 := v1.(sql.NullInt64)
		valList = append(valList, fmt.Sprintf(`'%v'`, v2.Int64))
	case reflect.TypeOf(pq.StringArray{}):
		v1 := v.Interface()
		v2 := v1.(pq.StringArray)
		v3, _ := v2.Value()
		valList = append(valList, fmt.Sprintf(`'%v'`, v3))
	default:
		valList = append(valList, fmt.Sprintf(`'%v'`, v))
	}
	return
}

