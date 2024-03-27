package utils

import (
	"reflect"
	"unsafe"
)

// During deepValueEqual, must keep track of checks that are
// in progress. The comparison algorithm assumes that all
// checks in progress are true when it reencounters them.
// Visited comparisons are stored in a map indexed by visit.
type visit struct {
	a1  unsafe.Pointer
	a2  unsafe.Pointer
	typ reflect.Type
}

// ConvertStrSlice2Map 将字符串 slice 转为 map[string]struct{}。
func ConvertStrSlice2Map(sl []string) map[string]struct{} {
	set := make(map[string]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

// InMap 判断字符串是否在 map 中。
func InMap(m map[string]struct{}, s string) bool {
	_, ok := m[s]
	return ok
}

//package main
//
//import (
//"fmt"
//"reflect"
//)

// GenericSliceDiff 通用切片差集函数
func GenericSliceDiff[T any](slice1, slice2 []T) []T {
	var diff []T

	for _, elem1 := range slice1 {
		found := false

		for _, elem2 := range slice2 {
			// 使用 reflect.DeepEqual 进行元素比较
			if reflect.DeepEqual(elem1, elem2) {
				found = true
				break
			}
		}

		if !found {
			diff = append(diff, elem1)
		}
	}

	return diff
}

func HasField(obj interface{}, fieldName string) bool {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == fieldName {
			return true
		}
	}

	return false
}

// GenericSliceIntersection 通用切片交集函数
func GenericSliceIntersection[T any](slice1, slice2 []T) []T {
	var intersection []T

	for _, elem1 := range slice1 {

		for _, elem2 := range slice2 {
			// 使用 reflect.DeepEqual 进行元素比较
			if reflect.DeepEqual(elem1, elem2) {
				intersection = append(intersection, elem1)
				break
			}
		}
	}

	return intersection
}

// GenericSliceUnion 通用切片并集函数
func GenericSliceUnion[T any](slice1, slice2 []T) []T {
	union := append([]T{}, slice1...)
	for _, elem2 := range slice2 {
		found := false

		for _, elem1 := range slice1 {

			// 使用 reflect.DeepEqual 进行元素比较
			if reflect.DeepEqual(elem1, elem2) {
				found = true
				break
			}
		}

		if !found {
			union = append(union, elem2)
		}
	}

	return union
}
