package zabbix

import (
	"fmt"
	"ormtest/base"
	"ormtest/action"
	"reflect"
)

func Discovery_db(){
	monitor := new(action.MySQLInfo)
	myref := reflect.ValueOf(monitor).Elem()
	typeOfType := myref.Type()
	fmt.Printf("\n{")
	fmt.Printf("\t\"data\":[\n")
	length := len(base.Cdb)

	begin := 0
	for db,_ := range base.Cdb {
		if begin < length {
			for i := 0; i < myref.NumField(); i++ {
				fmt.Printf("\t{\n")
				fmt.Printf("\t\t\"{#DBNAME}\":\"%s\",\n", db)
				fmt.Printf("\t\t\"{#ITEM}\":\"%s\"\n", typeOfType.Field(i).Name)
				if begin == length-1 && i == myref.NumField()-1 {
					fmt.Printf("\t}\n")
				} else {
					fmt.Printf("\t},")
				}
			}
			begin++
		}
	}
	fmt.Printf("\t]\n")
	fmt.Printf("}\n")
}

