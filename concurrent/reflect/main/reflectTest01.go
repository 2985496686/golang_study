package main

import (
	"fmt"
	"reflect"
)

type user struct {
	id    int64 `ss`
	name  string
	score int32
}

func (u user) getNameAndScore() (string, int32) {
	return u.name, u.score
}

func (u user) getId() int64 {
	return u.id
}

func main() {
	u := user{
		name:  "张三",
		id:    1,
		score: 100,
	}
	userType := reflect.TypeOf(u)
	/*
		Field(i int) StructField
		----- 返回结构体类型指定下标处的字段信息,若参数不是一个结构体类型会引发恐慌：panic: reflect: Field of non-struct type string
		StructField中有如下信息：
		Name string           ------- 字段名称
		PkgPath string        ------- 结构体所在的包
		Type      Type        ------- 字段类型
		Tag       StructTag   -------字段的标记，若没有标记，默认为0
		Offset    uintptr     -------字段在结构体中的偏移量(结构体对其)
		Index     []int       -------字段在结构体中的下标
		Anonymous bool        -------是否是嵌入式字段
	*/
	field1 := userType.Field(0)
	fmt.Printf("%v\n", field1) //打印结果为 {id main int64 ss 0 [0] false}

	//NumField() int 返回结构体中的字段数量
	numField := userType.NumField()
	fmt.Println(numField)

	/*
			FieldByName(name string)(StructField,bool)
		    根据字段名获取字段信息，第二个返回值表示是否找到
	*/
	fieldName, isFind := userType.FieldByName("id")
	fmt.Printf("%v  %v \n", fieldName, isFind)

	/*
		FieldByIndex(index []int) StructField
		--- 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息，没有找到时返回零值。当类型不是结构体或索引超界时发生恐慌
	*/

	/*
		FieldByNameFunc(match func(string) bool) (StructField, bool)
		返回满足条件的字段信息
	*/
	field, _ := userType.FieldByNameFunc(func(s string) bool {
		if s == "name" {
			return true
		}
		return false
	})
	fmt.Println(field) //{name main string  8 [1] false}
}
