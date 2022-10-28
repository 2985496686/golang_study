package main

import (
	"GoCodeb/src/concurrent/reflect/util"
	"fmt"
	"reflect"
)

type user struct {
	Id    int64 `ss`
	Name  string
	Score int32
}

func (u user) GetNameAndScore() (string, int32) {
	return u.Name, u.Score
}

func (u user) GetId() int64 {
	return u.Id
}

func main() {
	u := user{
		Name:  "张三",
		Id:    1,
		Score: 100,
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

	//获取一个值的Value
	userValue := reflect.ValueOf(u)
	//由反射类型转换为接口类型
	u2 := userValue.Interface().(user)
	fmt.Printf("%T", u2)

	//通过反射修改变量的值
	/*
		1. 被修改的变量必须是可以被寻址的
		2.若修改的是结构体字段，被修改的字段必须是可以被导出的
	*/
	/*
		id := userValue.Field(1) //获取第一个字段的值
		id.SetInt(30)
		执行上面代码会报panic，因为这样的值是不可以被修改的。
		我们在获取userValue时 :reflect.ValueOf(u)，这里接收到的u只是一个副本，对它进行修改并不会改变外部的u，所以这里直接设置为了不可修改
	*/
	userValue2 := reflect.ValueOf(&u).Elem()
	id := userValue2.Field(0)
	id.SetInt(40)
	fmt.Println(u)

	/*
		通过反射获取方法,执行方法
	*/
	methodType := userType.Method(0)
	fmt.Println(methodType)

	methodValue := userValue.Method(0)
	call := methodValue.Call([]reflect.Value{})
	fmt.Printf("%v", call[0])
	util.Add(1, 2)
}
