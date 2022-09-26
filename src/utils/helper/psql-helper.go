package helper

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
	restaurant2 "sykros-pro/gopro/src/dto"
)

func RawSQLScanner(query *gorm.DB, source interface{}) {
	rows, err := query.Rows()
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)
	for rows.Next() {
		err := query.ScanRows(rows, source)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func RawlSqlHelper(in interface{}) []string {
	rawlSqlSequence := make([]string, 0)
	switch in.(type) {
	case *restaurant2.UpdateRestaurantDto:
		setVariableSql := GetFieldAndValueToUpdate(in)
		rawlSql := "UPDATE RESTAURANTS " + setVariableSql + " WHERE ID = @id;"
		rawlSqlSequence = append(rawlSqlSequence, rawlSql)
	}
	return rawlSqlSequence
}

func GetFieldAndValueToUpdate(in interface{}) string {
	rawSql := make([]string, 0)
	data := map[string]interface{}{}
	switch in.(type) {
	case *restaurant2.UpdateRestaurantDto:
		for k := 0; k < reflect.TypeOf(in).Elem().NumField(); k++ {
			valueField := reflect.
				ValueOf(in).
				Elem().Field(k)
			keyFieldBytag := detachTagFromField(reflect.
				TypeOf(in).
				Elem().Field(k).
				Tag)
			if !isEmpty(valueField) {
				fmt.Printf("key[%v]:value[%v]\n",
					detachTagFromField(reflect.TypeOf(in).Elem().Field(k).Tag),
					reflect.ValueOf(reflect.ValueOf(in).
						Elem().Field(k)))


				data[keyFieldBytag] = valueField
			}
		}
	}

	for k, v := range data {
		rawSql = append(rawSql, fmt.Sprintf("SET %s=%v", k, v))
	}
	return strings.Join(rawSql, ",")
}

func detachTagFromField(tag reflect.StructTag) string {
	arraySlice := strings.Split(strings.Split(string(tag), ":")[1], "")
	return strings.Join(arraySlice[1:len(arraySlice)-1], "")
}
