package utils

import (
	"reflect"
)

func SetField(myStruct interface{}, field string, newData interface{}) {
	// truyền vào địa chỉ của biến
	if reflect.TypeOf(myStruct).Kind() == reflect.Ptr {
		// Tạo 1 biến chứa giá trị trong interface
		v := reflect.ValueOf(myStruct).Elem()

		// Tạo 1 biến tạm thời mang kiểu struct của v
		tmp := reflect.New(v.Elem().Type()).Elem()

		// copy giá trị của v sang tmp
		tmp.Set(v.Elem())

		switch reflect.ValueOf(newData).Kind() {
		case reflect.Float64:
			switch tmp.FieldByName(field).Kind() {
			case reflect.Int:
				newData = int(newData.(float64))
			case reflect.Uint:
				newData = uint(newData.(float64))
			case reflect.Float32:
				newData = float32(newData.(float64))
			case reflect.Float64:
				newData = float64(newData.(float64))
			}
		}

		// Thực hiện thay đổi giá trị của tmp theo filed
		tmp.FieldByName(field).Set(reflect.ValueOf(newData))

		// Thay đổi xong thì lấy giá trị mới của tmp set cho v
		v.Set(tmp)
	}
}

func MapInterface(dataQuery map[string]interface{}, data interface{}) {
	newData := reflect.ValueOf(data).Elem().Interface()
	for i := 0; i < reflect.ValueOf(newData).NumField(); i++ {
		keyUpper := reflect.TypeOf(newData).Field(i).Name
		keyLower := FirstLetterToLower(keyUpper)

		if dataQuery[keyLower] != nil {
			SetField(data, keyUpper, dataQuery[keyLower])
		}
	}
}
