package f

import "reflect"

func MapBool(sliceA []bool, f func(bool) bool) []bool {
	sliceB := []bool{}
	for _, b := range sliceA {
		sliceB = append(sliceB, f(b))
	}
	return sliceB
}

func EachBool(slice []int, f func(int)) {
	for _, b := range slice {
		f(b)
	}
}

func MapString(sliceA []string, f func(string) string) []string {
	sliceB := []string{}
	for _, b := range sliceA {
		sliceB = append(sliceB, f(b))
	}
	return sliceB
}

func EachString(slice []int, f func(int)) {
	for _, b := range slice {
		f(b)
	}
}

func MapInt(sliceA []int, f func(int) int) []int {
	sliceB := []int{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachInt(slice []int, f func(int)) {
	for _, i := range slice {
		f(i)
	}
}

func MapInt8(sliceA []int8, f func(int8) int8) []int8 {
	sliceB := []int8{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachInt8(slice []int8, f func(int8)) {
	for _, i := range slice {
		f(i)
	}
}

func MapInt16(sliceA []int16, f func(int16) int16) []int16 {
	sliceB := []int16{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachInt16(slice []int16, f func(int16)) {
	for _, i := range slice {
		f(i)
	}
}

func MapInt32(sliceA []int32, f func(int32) int32) []int32 {
	sliceB := []int32{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachInt32(slice []int32, f func(int32)) {
	for _, i := range slice {
		f(i)
	}
}

func MapInt64(sliceA []int64, f func(int64) int64) []int64 {
	sliceB := []int64{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachInt64(slice []int64, f func(int64)) {
	for _, i := range slice {
		f(i)
	}
}

func MapUInt(sliceA []uint, f func(uint) uint) []uint {
	sliceB := []uint{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachUInt(slice []uint, f func(uint)) {
	for _, i := range slice {
		f(i)
	}
}

func MapUInt8(sliceA []uint8, f func(uint8) uint8) []uint8 {
	sliceB := []uint8{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachUInt8(slice []uint8, f func(uint8)) {
	for _, i := range slice {
		f(i)
	}
}

func MapUInt16(sliceA []uint16, f func(uint16) uint16) []uint16 {
	sliceB := []uint16{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachUInt16(slice []uint16, f func(uint16)) {
	for _, i := range slice {
		f(i)
	}
}

func MapUInt32(sliceA []uint32, f func(uint32) uint32) []uint32 {
	sliceB := []uint32{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachUInt32(slice []uint32, f func(uint32)) {
	for _, i := range slice {
		f(i)
	}
}

func MapUInt64(sliceA []uint64, f func(uint64) uint64) []uint64 {
	sliceB := []uint64{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachUInt64(slice []uint64, f func(uint64)) {
	for _, i := range slice {
		f(i)
	}
}

func MapUIntPtr(sliceA []uintptr, f func(uintptr) uintptr) []uintptr {
	sliceB := []uintptr{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachUIntPtr(slice []uintptr, f func(uintptr)) {
	for _, i := range slice {
		f(i)
	}
}

func MapByte(sliceA []byte, f func(byte) byte) []byte {
	sliceB := []byte{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachByte(slice []byte, f func(byte)) {
	for _, i := range slice {
		f(i)
	}
}

func MapRune(sliceA []rune, f func(rune) rune) []rune {
	sliceB := []rune{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachRune(slice []rune, f func(rune)) {
	for _, i := range slice {
		f(i)
	}
}

func MapFloat32(sliceA []float32, f func(float32) float32) []float32 {
	sliceB := []float32{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachFloat32(slice []float32, f func(float32)) {
	for _, i := range slice {
		f(i)
	}
}

func MapFloat64(sliceA []float64, f func(float64) float64) []float64 {
	sliceB := []float64{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachFloat64(slice []float64, f func(float64)) {
	for _, i := range slice {
		f(i)
	}
}

func MapComplex64(sliceA []complex64, f func(complex64) complex64) []complex64 {
	sliceB := []complex64{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachComplex64(slice []complex64, f func(complex64)) {
	for _, i := range slice {
		f(i)
	}
}

func MapComplex128(sliceA []complex128, f func(complex128) complex128) []complex128 {
	sliceB := []complex128{}
	for _, i := range sliceA {
		sliceB = append(sliceB, f(i))
	}
	return sliceB
}

func EachComplex128(slice []complex128, f func(complex128)) {
	for _, i := range slice {
		f(i)
	}
}

func Map(slice interface{}, fn interface{}) interface{} {
	f := reflect.ValueOf(fn)
	a := reflect.ValueOf(slice)
	b := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(fn).Out(0)), 0, a.Cap())
	for i := 0; i < a.Len(); i++ {
		b = reflect.Append(b, f.Call([]reflect.Value{a.Index(i)})[0])
	}
	return b.Interface()
}

func Each(slice interface{}, fn interface{}) {
	f := reflect.ValueOf(fn)
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		f.Call([]reflect.Value{s.Index(i)})
	}
}
