package utils

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
)

func RawSQLScanner(query *gorm.DB, source interface{}) {
	rows, err := query.Rows()
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

//
//type Ordered interface {
//	uint64 | string | time.Time | float32
//}
//
//func GenericTypeCast[T](inputVal any, ref T) {
//	val, err := inputVal.(T)
//	if !err {
//		ref = val
//	}
//	//ref = nil
//}
