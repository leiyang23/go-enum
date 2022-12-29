package enum

import (
	"bytes"
	"fmt"
	"reflect"
)

// Enum 定义枚举类型基础结构
type Enum struct {
	Name  string
	Value int
}

func (e Enum) String() string {
	return fmt.Sprintf("%d-%s", e.Value, e.Name)
}

func (e Enum) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	obj := fmt.Sprintf("{\"value\": %d, \"name\": \"%s\"}", e.Value, e.Name)
	_, err := buf.WriteString(obj)
	return buf.Bytes(), err
}

// ErrorEnum 统一错误的枚举值
var ErrorEnum = Enum{Value: 1000, Name: "ErrorEnum"}

func MakeEnum(eb interface{}) error {
	var structValue reflect.Value
	var structType reflect.Type

	structType = reflect.TypeOf(eb)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
		structValue = reflect.ValueOf(eb).Elem()
	} else {
		return fmt.Errorf("only receive pointer")
	}

	for num := 0; num < structType.NumField(); num++ {
		_, ok := structValue.Field(num).Interface().(Enum)
		if !ok {
			continue
		}

		fieldName := structType.Field(num).Name

		// 字段名到字段值的映射方法
		// 通过定义结构体的 Convert 方法来重写默认的映射方式
		name := fieldName
		if structValue.MethodByName("Convert").IsValid() {
			name = structValue.MethodByName("Convert").Call([]reflect.Value{reflect.ValueOf(fieldName)})[0].String()
		} else {
			name = Convert(fieldName)
		}

		// 防止枚举值为 0，在索引值上加 1
		e := Enum{Name: name, Value: num + 1}
		if structValue.FieldByName(fieldName).CanSet() {
			structValue.FieldByName(fieldName).Set(reflect.ValueOf(e))
		} else {
			return fmt.Errorf("can not set value of '%s'", fieldName)
		}
	}
	return nil
}

func Validate(eb interface{}, v int) Enum {
	var structValue reflect.Value
	var structType reflect.Type

	structType = reflect.TypeOf(eb)
	if structType.Kind() == reflect.Ptr {
		structValue = reflect.ValueOf(eb).Elem()
	} else {
		structValue = reflect.ValueOf(eb)
	}

	for num := 0; num < structValue.NumField(); num++ {
		e, ok := structValue.Field(num).Interface().(Enum)
		if !ok {
			continue
		}
		if e.Value == v {
			return e
		}
	}
	return ErrorEnum
}

func List(eb interface{}) []Enum {
	s := make([]Enum, 0)
	var structValue reflect.Value
	var structType reflect.Type

	structType = reflect.TypeOf(eb)
	if structType.Kind() == reflect.Ptr {
		structValue = reflect.ValueOf(eb).Elem()
	} else {
		structValue = reflect.ValueOf(eb)
	}

	for num := 0; num < structValue.NumField(); num++ {
		e, ok := structValue.Field(num).Interface().(Enum)
		if !ok {
			continue
		}
		s = append(s, e)
	}
	return s
}
