package fn

import "reflect"

func FindBool(sliceA []bool, fn func(bool) bool) []bool {
	sliceB := []bool{}
	for _, b := range sliceA {
		if fn(b) {
			sliceB = append(sliceB, b)
		}
	}
	return sliceB
}

func MapBool(sliceA []bool, fn func(bool) bool) []bool {
	sliceB := []bool{}
	for _, b := range sliceA {
		sliceB = append(sliceB, fn(b))
	}
	return sliceB
}

func EachBool(slice []int, fn func(int)) {
	for _, i := range slice {
		fn(i)
	}
}

func FindString(sliceA []string, fn func(string) bool) []string {
	sliceB := []string{}
	for _, s := range sliceA {
		if fn(s) {
			sliceB = append(sliceB, s)
		}
	}
	return sliceB
}

func MapString(sliceA []string, fn func(string) string) []string {
	sliceB := []string{}
	for _, s := range sliceA {
		sliceB = append(sliceB, fn(s))
	}
	return sliceB
}

func EachString(slice []int, fn func(int)) {
	for _, i := range slice {
		fn(i)
	}
}

func FindInt(sliceA []int, fn func(int) bool) []int {
	sliceB := []int{}
	for _, i := range sliceA {
		if fn(i) {
			sliceB = append(sliceB, i)
		}
	}
	return sliceB
}

func MapInt(sliceA []int, fn func(int) int) []int {
	sliceB := []int{}
	for _, i := range sliceA {
		sliceB = append(sliceB, fn(i))
	}
	return sliceB
}

func EachInt(slice []int, fn func(int)) {
	for _, i := range slice {
		fn(i)
	}
}

func FindInt8(sliceA []int8, fn func(int8) bool) []int8 {
	sliceB := []int8{}
	for _, i := range sliceA {
		if fn(i) {
			sliceB = append(sliceB, i)
		}
	}
	return sliceB
}

func MapInt8(sliceA []int8, fn func(int8) int8) []int8 {
	sliceB := []int8{}
	for _, i := range sliceA {
		sliceB = append(sliceB, fn(i))
	}
	return sliceB
}

func EachInt8(slice []int8, fn func(int8)) {
	for _, i := range slice {
		fn(i)
	}
}

func FindInt16(sliceA []int16, fn func(int16) bool) []int16 {
	sliceB := []int16{}
	for _, i := range sliceA {
		if fn(i) {
			sliceB = append(sliceB, i)
		}
	}
	return sliceB
}

func MapInt16(sliceA []int16, fn func(int16) int16) []int16 {
	sliceB := []int16{}
	for _, i := range sliceA {
		sliceB = append(sliceB, fn(i))
	}
	return sliceB
}

func EachInt16(slice []int16, fn func(int16)) {
	for _, i := range slice {
		fn(i)
	}
}

func FindInt32(sliceA []int32, fn func(int32) bool) []int32 {
	sliceB := []int32{}
	for _, i := range sliceA {
		if fn(i) {
			sliceB = append(sliceB, i)
		}
	}
	return sliceB
}

func MapInt32(sliceA []int32, fn func(int32) int32) []int32 {
	sliceB := []int32{}
	for _, i := range sliceA {
		sliceB = append(sliceB, fn(i))
	}
	return sliceB
}

func EachInt32(slice []int32, fn func(int32)) {
	for _, i := range slice {
		fn(i)
	}
}

func FindInt64(sliceA []int64, fn func(int64) bool) []int64 {
	sliceB := []int64{}
	for _, i := range sliceA {
		if fn(i) {
			sliceB = append(sliceB, i)
		}
	}
	return sliceB
}

func MapInt64(sliceA []int64, fn func(int64) int64) []int64 {
	sliceB := []int64{}
	for _, i := range sliceA {
		sliceB = append(sliceB, fn(i))
	}
	return sliceB
}

func EachInt64(slice []int64, fn func(int64)) {
	for _, i := range slice {
		fn(i)
	}
}

func FindUInt(sliceA []uint, fn func(uint) bool) []uint {
	sliceB := []uint{}
	for _, u := range sliceA {
		if fn(u) {
			sliceB = append(sliceB, u)
		}
	}
	return sliceB
}

func MapUInt(sliceA []uint, fn func(uint) uint) []uint {
	sliceB := []uint{}
	for _, u := range sliceA {
		sliceB = append(sliceB, fn(u))
	}
	return sliceB
}

func EachUInt(slice []uint, fn func(uint)) {
	for _, u := range slice {
		fn(u)
	}
}

func FindUInt8(sliceA []uint8, fn func(uint8) bool) []uint8 {
	sliceB := []uint8{}
	for _, u := range sliceA {
		if fn(u) {
			sliceB = append(sliceB, u)
		}
	}
	return sliceB
}

func MapUInt8(sliceA []uint8, fn func(uint8) uint8) []uint8 {
	sliceB := []uint8{}
	for _, u := range sliceA {
		sliceB = append(sliceB, fn(u))
	}
	return sliceB
}

func EachUInt8(slice []uint8, fn func(uint8)) {
	for _, u := range slice {
		fn(u)
	}
}

func FindUInt16(sliceA []uint16, fn func(uint16) bool) []uint16 {
	sliceB := []uint16{}
	for _, u := range sliceA {
		if fn(u) {
			sliceB = append(sliceB, u)
		}
	}
	return sliceB
}

func MapUInt16(sliceA []uint16, fn func(uint16) uint16) []uint16 {
	sliceB := []uint16{}
	for _, u := range sliceA {
		sliceB = append(sliceB, fn(u))
	}
	return sliceB
}

func EachUInt16(slice []uint16, fn func(uint16)) {
	for _, u := range slice {
		fn(u)
	}
}

func FindUInt32(sliceA []uint32, fn func(uint32) bool) []uint32 {
	sliceB := []uint32{}
	for _, u := range sliceA {
		if fn(u) {
			sliceB = append(sliceB, u)
		}
	}
	return sliceB
}

func MapUInt32(sliceA []uint32, fn func(uint32) uint32) []uint32 {
	sliceB := []uint32{}
	for _, u := range sliceA {
		sliceB = append(sliceB, fn(u))
	}
	return sliceB
}

func EachUInt32(slice []uint32, fn func(uint32)) {
	for _, u := range slice {
		fn(u)
	}
}

func FindUInt64(sliceA []uint64, fn func(uint64) bool) []uint64 {
	sliceB := []uint64{}
	for _, u := range sliceA {
		if fn(u) {
			sliceB = append(sliceB, u)
		}
	}
	return sliceB
}

func MapUInt64(sliceA []uint64, fn func(uint64) uint64) []uint64 {
	sliceB := []uint64{}
	for _, u := range sliceA {
		sliceB = append(sliceB, fn(u))
	}
	return sliceB
}

func EachUInt64(slice []uint64, fn func(uint64)) {
	for _, u := range slice {
		fn(u)
	}
}

func FindUIntPtr(sliceA []uintptr, fn func(uintptr) bool) []uintptr {
	sliceB := []uintptr{}
	for _, u := range sliceA {
		if fn(u) {
			sliceB = append(sliceB, u)
		}
	}
	return sliceB
}

func MapUIntPtr(sliceA []uintptr, fn func(uintptr) uintptr) []uintptr {
	sliceB := []uintptr{}
	for _, u := range sliceA {
		sliceB = append(sliceB, fn(u))
	}
	return sliceB
}

func EachUIntPtr(slice []uintptr, fn func(uintptr)) {
	for _, u := range slice {
		fn(u)
	}
}

func FindByte(sliceA []byte, fn func(byte) bool) []byte {
	sliceB := []byte{}
	for _, b := range sliceA {
		if fn(b) {
			sliceB = append(sliceB, b)
		}
	}
	return sliceB
}

func MapByte(sliceA []byte, fn func(byte) byte) []byte {
	sliceB := []byte{}
	for _, b := range sliceA {
		sliceB = append(sliceB, fn(b))
	}
	return sliceB
}

func EachByte(slice []byte, fn func(byte)) {
	for _, b := range slice {
		fn(b)
	}
}

func FindRune(sliceA []rune, fn func(rune) bool) []rune {
	sliceB := []rune{}
	for _, r := range sliceA {
		if fn(r) {
			sliceB = append(sliceB, r)
		}
	}
	return sliceB
}

func MapRune(sliceA []rune, fn func(rune) rune) []rune {
	sliceB := []rune{}
	for _, r := range sliceA {
		sliceB = append(sliceB, fn(r))
	}
	return sliceB
}

func EachRune(slice []rune, fn func(rune)) {
	for _, r := range slice {
		fn(r)
	}
}

func FindFloat32(sliceA []float32, fn func(float32) bool) []float32 {
	sliceB := []float32{}
	for _, f := range sliceA {
		if fn(f) {
			sliceB = append(sliceB, f)
		}
	}
	return sliceB
}

func MapFloat32(sliceA []float32, fn func(float32) float32) []float32 {
	sliceB := []float32{}
	for _, f := range sliceA {
		sliceB = append(sliceB, fn(f))
	}
	return sliceB
}

func EachFloat32(slice []float32, fn func(float32)) {
	for _, f := range slice {
		fn(f)
	}
}

func FindFloat64(sliceA []float64, fn func(float64) bool) []float64 {
	sliceB := []float64{}
	for _, f := range sliceA {
		if fn(f) {
			sliceB = append(sliceB, f)
		}
	}
	return sliceB
}

func MapFloat64(sliceA []float64, fn func(float64) float64) []float64 {
	sliceB := []float64{}
	for _, f := range sliceA {
		sliceB = append(sliceB, fn(f))
	}
	return sliceB
}

func EachFloat64(slice []float64, fn func(float64)) {
	for _, f := range slice {
		fn(f)
	}
}

func FindComplex64(sliceA []complex64, fn func(complex64) bool) []complex64 {
	sliceB := []complex64{}
	for _, c := range sliceA {
		if fn(c) {
			sliceB = append(sliceB, c)
		}
	}
	return sliceB
}

func MapComplex64(sliceA []complex64, fn func(complex64) complex64) []complex64 {
	sliceB := []complex64{}
	for _, c := range sliceA {
		sliceB = append(sliceB, fn(c))
	}
	return sliceB
}

func EachComplex64(slice []complex64, fn func(complex64)) {
	for _, c := range slice {
		fn(c)
	}
}

func FindComplex128(sliceA []complex128, fn func(complex128) bool) []complex128 {
	sliceB := []complex128{}
	for _, c := range sliceA {
		if fn(c) {
			sliceB = append(sliceB, c)
		}
	}
	return sliceB
}

func MapComplex128(sliceA []complex128, fn func(complex128) complex128) []complex128 {
	sliceB := []complex128{}
	for _, c := range sliceA {
		sliceB = append(sliceB, fn(c))
	}
	return sliceB
}

func EachComplex128(slice []complex128, fn func(complex128)) {
	for _, c := range slice {
		fn(c)
	}
}

func Map(slice interface{}, fn interface{}) interface{} {
	fun := reflect.ValueOf(fn)
	a := reflect.ValueOf(slice)
	b := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(fn).Out(0)), 0, a.Cap())
	for i := 0; i < a.Len(); i++ {
		b = reflect.Append(b, fun.Call([]reflect.Value{a.Index(i)})[0])
	}
	return b.Interface()
}

func Each(slice interface{}, fn interface{}) {
	fun := reflect.ValueOf(fn)
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		fun.Call([]reflect.Value{s.Index(i)})
	}
}
