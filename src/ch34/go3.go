package ch34

import (
	"fmt"
	"reflect"
)

// 在 Go 语言中，reflect 实现了运行时反射。reflect 包会帮助识别 interface{} 变量的底层具体类型和具体值。这正是我们所需要的。createQuery 函数接收 interface{} 参数，根据它的具体类型和具体值，创建 SQL 查询。

type order1 struct {
	orderId    int
	customerId int
}
type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery1(q interface{}) {

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.String:
				sprintf := "%s, \"%s\""
				if i == 0 {
					sprintf = "%s\"%s\""
				}
				query = fmt.Sprintf(sprintf, query, v.Field(i).String())
			case reflect.Int:
				sprintf := "%s , %d"
				if i == 0 {
					sprintf = "%s%d"
				}
				query = fmt.Sprintf(sprintf, query, v.Field(i).Int())
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func MainCall3() {
	o := order1{orderId: 456, customerId: 56}
	createQuery1(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery1(e)
	i := 90
	createQuery1(i)
}
