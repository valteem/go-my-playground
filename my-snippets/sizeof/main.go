package main

import (
	"fmt"
	"unsafe"
)

type EmptyStruct struct {}

type Int32Struct struct {
	value int32
}

type Int64Struct struct {
	value int64
}

type StringStruct struct {
	value string
}

type ByteStruct struct {
	value byte
}

type fieldNum1Struct struct {
	field1 int32
}

type fieldNum2Struct struct {
	field1 int32
	field2 int32
}
type fieldNum3Struct struct {
	field1 int32
	field2 int32
	field3 int32
}
type fieldNum4Struct struct {
	field1 int32
	field2 int32
	field3 int32
	field4 int32
}

type fieldNum5Struct struct {
	field1 int32
	field2 int32
	field3 int32
	field4 int32
	field5 int32
}

type mixedFieldSizeOrder struct {
	field1 int32
	field2 int64
	field3 int32
	field4 int64
	field5 int32
}

type decrFieldSizeOrder struct {
	field1 int64
	field2 int64
	field3 int32
	field4 int32
	field5 int32
}

type incrFieldSizeOrder struct {
	field1 int32
	field2 int32
	field3 int32
	field4 int64
	field5 int64
}

func sizeof(inp interface{}) uint32 {
	return uint32(unsafe.Sizeof(inp)) // something wrong with converting uintptr to uint32
}

func sizeofInt32(inp Int32Struct) uint32 {
	return uint32(unsafe.Sizeof(inp))
}

func main() {

/*	
	var es EmptyStruct
	fmt.Println("Empty struct size is", unsafe.Sizeof(es))

	varInt32struct := Int32Struct{value: 1}
	fmt.Println("Int32 struct size (Sizeof) is", unsafe.Sizeof(varInt32struct)) // 4, looks reasonable
	fmt.Println("Int32 struct size (sizeof) is", sizeof(varInt32struct))        // 16, because casting anything to interface{} makes its Sizeof 16 (see below)
	fmt.Println("Int32 struct size (sizeof) is", sizeofInt32(varInt32struct))   
	var i32v int32
	fmt.Println("Int32 value size is", unsafe.Sizeof(i32v))

	var i64s Int64Struct
	fmt.Println("Int64 struct size is", unsafe.Sizeof(i64s))
	var i64v int64
	fmt.Println("Int64 value size is", unsafe.Sizeof(i64v))

	var ss StringStruct
	fmt.Println("String struct size is", unsafe.Sizeof(ss))
	var sv string
	fmt.Println("String struct value is", unsafe.Sizeof(sv))

	var bs ByteStruct
	fmt.Println("Byte struct size is", sizeof(bs)) */

	type iface interface{}
	/*
	var varInt32 int32 = 11
	fmt.Println(unsafe.Sizeof(varInt32), unsafe.Sizeof((iface)(varInt32)))
	var varInt64 int64 = 11
	fmt.Println(unsafe.Sizeof(varInt64), unsafe.Sizeof((iface)(varInt64))) */

	varFieldNum1 := fieldNum1Struct{field1: 1}
	fmt.Println(unsafe.Sizeof(varFieldNum1), unsafe.Sizeof((iface)(varFieldNum1)))
	varFieldNum2 := fieldNum2Struct{field1: 1, field2: 1}
	fmt.Println(unsafe.Sizeof(varFieldNum2), unsafe.Sizeof((iface)(varFieldNum2)))
	varFieldNum3 := fieldNum3Struct{field1: 1, field2: 1, field3: 1}
	fmt.Println(unsafe.Sizeof(varFieldNum3), unsafe.Sizeof((iface)(varFieldNum3)))
	varFieldNum4 := fieldNum4Struct{field1: 1, field2: 1, field3: 1, field4: 1}
	fmt.Println(unsafe.Sizeof(varFieldNum4), unsafe.Sizeof((iface)(varFieldNum4)))
	varFieldNum5 := fieldNum5Struct{field1: 1, field2: 1, field3: 1, field4: 1, field5: 1}
	fmt.Println(unsafe.Sizeof(varFieldNum5), unsafe.Sizeof((iface)(varFieldNum5)))

// https://stackoverflow.com/questions/39063530/optimising-datastructure-word-alignment-padding-in-golang

	varMixed := mixedFieldSizeOrder{field1: 1, field2: 1, field3: 1, field4: 1, field5: 1}
	fmt.Println(unsafe.Sizeof(varMixed), unsafe.Sizeof((iface)(varMixed)))
	varDecr := decrFieldSizeOrder{field1: 1, field2: 1, field3: 1, field4: 1, field5: 1}
	fmt.Println(unsafe.Sizeof(varDecr), unsafe.Sizeof((iface)(varDecr)))
	varIncr := incrFieldSizeOrder{field1: 1, field2: 1, field3: 1, field4: 1, field5: 1}
	fmt.Println(unsafe.Sizeof(varIncr), unsafe.Sizeof((iface)(varIncr)))

}