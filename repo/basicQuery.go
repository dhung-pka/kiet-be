package repo

import (
	"fmt"
	"log"
	"reflect"
	"service/config"
	"service/model"
	"service/utils"
	"time"

	"gorm.io/gorm"
)

type basicQueryRepo struct {
	db *gorm.DB
}

type BasicQueryRepo interface {
	BasicQueryInsertData(data interface{}, modelType string) (interface{}, error)
}

func (b *basicQueryRepo) BasicQueryInsertData(data interface{}, modelType string) (interface{}, error) {
	modelTypeInsert := model.TypeModel[modelType]
	dataMap := map[string]interface{}{}

	for i := 0; i < reflect.TypeOf(data).NumField(); i++ {
		key := reflect.TypeOf(data).Field(i).Name
		switch key {
		// case "Id":
		// 	continue
		case "CreatedAt":
			dataMap[key] = time.Now()
		case "UpdatedAt":
		case "DeletedAt":
			dataMap[key] = nil
		default:
			value := reflect.ValueOf(data).FieldByName(key)
			switch reflect.ValueOf(data).FieldByName(key).Kind() {
			case reflect.Int:
				dataMap[key] = int(value.Interface().(int))
			case reflect.Uint:
				dataMap[key] = uint(value.Interface().(uint))
			case reflect.Float64:
				dataMap[key] = float64(value.Interface().(float64))
			case reflect.Float32:
				dataMap[key] = float32(value.Interface().(float32))
			case reflect.Struct:
				continue
			default:
				dataMap[key] = value
			}
		}
	}

	if err := b.db.Debug().Model(&modelTypeInsert).Where("id = ?", dataMap["Id"]).Omit("Id").Updates(&dataMap).Error; err != nil {
		return nil, err
	}

	var result interface{}
	stringSql :=
		`
	UPDATE districts
	SET 
	name = CASE id 
		when 1 THEN 'heh1'
		when 2 THEN 'heh2'	
		when 3 THEN 'heh3'
		when 4 THEN 'heh4'
		else name
		end,
	code = CASE id 
		when 1 THEN 'heh1'
		when 2 THEN 'heh2'	
		when 3 THEN 'heh3'
		when 4 THEN 'heh4'
		else code
		end
	WHERE id IN(1, 2, 3, 4)
	RETURN id
	`
	b.db.Debug().Raw(stringSql).Scan(&result)

	log.Println("R: ", result)

	dataReturn := map[string]interface{}{}

	for key, value := range dataMap {
		field := utils.FirstLetterToLower(key)
		if reflect.ValueOf(value).Kind() == reflect.Struct {
			newValue := fmt.Sprintf("%v", value)
			dataReturn[field] = newValue
		} else {
			dataReturn[field] = value
		}

	}

	return dataReturn, nil
}

func NewBasicQueryRepo() BasicQueryRepo {
	db := config.GetDB()
	return &basicQueryRepo{
		db: db,
	}
}
