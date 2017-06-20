package f

import (
	"reflect"
)

func FindBool(slice []bool, fn func(bool) bool) []bool {
	sliceCh := make(chan []bool, 1)
	sliceCh <- []bool{}
	for _, Bool := range slice {
		Bool := Bool
		go (func() {
			if fn(Bool) {
				sliceCh <- append(<-sliceCh, Bool)
			}
		})()
	}
	return <-sliceCh
}

func MapBool(slice []bool, fn func(bool) bool) []bool {
	sliceCh := make(chan []bool, 1)
	sliceCh <- []bool{}
	for _, Bool := range slice {
		Bool := Bool
		go (func() {
			v := fn(Bool)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachBool(slice []bool, fn func(bool)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Bool := range slice {
		Bool := Bool
		go (func() {
			fn(Bool)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindString(slice []string, fn func(string) bool) []string {
	sliceCh := make(chan []string, 1)
	sliceCh <- []string{}
	for _, String := range slice {
		String := String
		go (func() {
			if fn(String) {
				sliceCh <- append(<-sliceCh, String)
			}
		})()
	}
	return <-sliceCh
}

func MapString(slice []string, fn func(string) string) []string {
	sliceCh := make(chan []string, 1)
	sliceCh <- []string{}
	for _, String := range slice {
		String := String
		go (func() {
			v := fn(String)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachString(slice []string, fn func(string)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, String := range slice {
		String := String
		go (func() {
			fn(String)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindInt(slice []int, fn func(int) bool) []int {
	sliceCh := make(chan []int, 1)
	sliceCh <- []int{}
	for _, Int := range slice {
		Int := Int
		go (func() {
			if fn(Int) {
				sliceCh <- append(<-sliceCh, Int)
			}
		})()
	}
	return <-sliceCh
}

func MapInt(slice []int, fn func(int) int) []int {
	sliceCh := make(chan []int, 1)
	sliceCh <- []int{}
	for _, Int := range slice {
		Int := Int
		go (func() {
			v := fn(Int)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachInt(slice []int, fn func(int)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Int := range slice {
		Int := Int
		go (func() {
			fn(Int)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindInt8(slice []int8, fn func(int8) bool) []int8 {
	sliceCh := make(chan []int8, 1)
	sliceCh <- []int8{}
	for _, Int8 := range slice {
		Int8 := Int8
		go (func() {
			if fn(Int8) {
				sliceCh <- append(<-sliceCh, Int8)
			}
		})()
	}
	return <-sliceCh
}

func MapInt8(slice []int8, fn func(int8) int8) []int8 {
	sliceCh := make(chan []int8, 1)
	sliceCh <- []int8{}
	for _, Int8 := range slice {
		Int8 := Int8
		go (func() {
			v := fn(Int8)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachInt8(slice []int8, fn func(int8)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Int8 := range slice {
		Int8 := Int8
		go (func() {
			fn(Int8)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindInt16(slice []int16, fn func(int16) bool) []int16 {
	sliceCh := make(chan []int16, 1)
	sliceCh <- []int16{}
	for _, Int16 := range slice {
		Int16 := Int16
		go (func() {
			if fn(Int16) {
				sliceCh <- append(<-sliceCh, Int16)
			}
		})()
	}
	return <-sliceCh
}

func MapInt16(slice []int16, fn func(int16) int16) []int16 {
	sliceCh := make(chan []int16, 1)
	sliceCh <- []int16{}
	for _, Int16 := range slice {
		Int16 := Int16
		go (func() {
			v := fn(Int16)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachInt16(slice []int16, fn func(int16)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Int16 := range slice {
		Int16 := Int16
		go (func() {
			fn(Int16)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindInt32(slice []int32, fn func(int32) bool) []int32 {
	sliceCh := make(chan []int32, 1)
	sliceCh <- []int32{}
	for _, Int32 := range slice {
		Int32 := Int32
		go (func() {
			if fn(Int32) {
				sliceCh <- append(<-sliceCh, Int32)
			}
		})()
	}
	return <-sliceCh
}

func MapInt32(slice []int32, fn func(int32) int32) []int32 {
	sliceCh := make(chan []int32, 1)
	sliceCh <- []int32{}
	for _, Int32 := range slice {
		Int32 := Int32
		go (func() {
			v := fn(Int32)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachInt32(slice []int32, fn func(int32)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Int32 := range slice {
		Int32 := Int32
		go (func() {
			fn(Int32)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindInt64(slice []int64, fn func(int64) bool) []int64 {
	sliceCh := make(chan []int64, 1)
	sliceCh <- []int64{}
	for _, Int64 := range slice {
		Int64 := Int64
		go (func() {
			if fn(Int64) {
				sliceCh <- append(<-sliceCh, Int64)
			}
		})()
	}
	return <-sliceCh
}

func MapInt64(slice []int64, fn func(int64) int64) []int64 {
	sliceCh := make(chan []int64, 1)
	sliceCh <- []int64{}
	for _, Int64 := range slice {
		Int64 := Int64
		go (func() {
			v := fn(Int64)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachInt64(slice []int64, fn func(int64)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Int64 := range slice {
		Int64 := Int64
		go (func() {
			fn(Int64)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindUInt(slice []uint, fn func(uint) bool) []uint {
	sliceCh := make(chan []uint, 1)
	sliceCh <- []uint{}
	for _, Uint := range slice {
		Uint := Uint
		go (func() {
			if fn(Uint) {
				sliceCh <- append(<-sliceCh, Uint)
			}
		})()
	}
	return <-sliceCh
}

func MapUInt(slice []uint, fn func(uint) uint) []uint {
	sliceCh := make(chan []uint, 1)
	sliceCh <- []uint{}
	for _, Uint := range slice {
		Uint := Uint
		go (func() {
			v := fn(Uint)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachUInt(slice []uint, fn func(uint)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, UInt := range slice {
		UInt := UInt
		go (func() {
			fn(UInt)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindUInt8(slice []uint8, fn func(uint8) bool) []uint8 {
	sliceCh := make(chan []uint8, 1)
	sliceCh <- []uint8{}
	for _, Uint8 := range slice {
		Uint8 := Uint8
		go (func() {
			if fn(Uint8) {
				sliceCh <- append(<-sliceCh, Uint8)
			}
		})()
	}
	return <-sliceCh
}

func MapUInt8(slice []uint8, fn func(uint8) uint8) []uint8 {
	sliceCh := make(chan []uint8, 1)
	sliceCh <- []uint8{}
	for _, Uint8 := range slice {
		Uint8 := Uint8
		go (func() {
			v := fn(Uint8)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachUInt8(slice []uint8, fn func(uint8)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, UInt8 := range slice {
		UInt8 := UInt8
		go (func() {
			fn(UInt8)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindUInt16(slice []uint16, fn func(uint16) bool) []uint16 {
	sliceCh := make(chan []uint16, 1)
	sliceCh <- []uint16{}
	for _, Uint16 := range slice {
		Uint16 := Uint16
		go (func() {
			if fn(Uint16) {
				sliceCh <- append(<-sliceCh, Uint16)
			}
		})()
	}
	return <-sliceCh
}

func MapUInt16(slice []uint16, fn func(uint16) uint16) []uint16 {
	sliceCh := make(chan []uint16, 1)
	sliceCh <- []uint16{}
	for _, Uint16 := range slice {
		Uint16 := Uint16
		go (func() {
			v := fn(Uint16)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachUInt16(slice []uint16, fn func(uint16)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, UInt16 := range slice {
		UInt16 := UInt16
		go (func() {
			fn(UInt16)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindUInt32(slice []uint32, fn func(uint32) bool) []uint32 {
	sliceCh := make(chan []uint32, 1)
	sliceCh <- []uint32{}
	for _, Uint32 := range slice {
		Uint32 := Uint32
		go (func() {
			if fn(Uint32) {
				sliceCh <- append(<-sliceCh, Uint32)
			}
		})()
	}
	return <-sliceCh
}

func MapUInt32(slice []uint32, fn func(uint32) uint32) []uint32 {
	sliceCh := make(chan []uint32, 1)
	sliceCh <- []uint32{}
	for _, Uint32 := range slice {
		Uint32 := Uint32
		go (func() {
			v := fn(Uint32)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachUInt32(slice []uint32, fn func(uint32)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, UInt32 := range slice {
		UInt32 := UInt32
		go (func() {
			fn(UInt32)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindUInt64(slice []uint64, fn func(uint64) bool) []uint64 {
	sliceCh := make(chan []uint64, 1)
	sliceCh <- []uint64{}
	for _, Uint64 := range slice {
		Uint64 := Uint64
		go (func() {
			if fn(Uint64) {
				sliceCh <- append(<-sliceCh, Uint64)
			}
		})()
	}
	return <-sliceCh
}

func MapUInt64(slice []uint64, fn func(uint64) uint64) []uint64 {
	sliceCh := make(chan []uint64, 1)
	sliceCh <- []uint64{}
	for _, Uint64 := range slice {
		Uint64 := Uint64
		go (func() {
			v := fn(Uint64)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachUInt64(slice []uint64, fn func(uint64)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, UInt64 := range slice {
		UInt64 := UInt64
		go (func() {
			fn(UInt64)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindUIntPtr(slice []uintptr, fn func(uintptr) bool) []uintptr {
	sliceCh := make(chan []uintptr, 1)
	sliceCh <- []uintptr{}
	for _, Uintptr := range slice {
		Uintptr := Uintptr
		go (func() {
			if fn(Uintptr) {
				sliceCh <- append(<-sliceCh, Uintptr)
			}
		})()
	}
	return <-sliceCh
}

func MapUIntPtr(slice []uintptr, fn func(uintptr) uintptr) []uintptr {
	sliceCh := make(chan []uintptr, 1)
	sliceCh <- []uintptr{}
	for _, Uintptr := range slice {
		Uintptr := Uintptr
		go (func() {
			v := fn(Uintptr)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachUIntPtr(slice []uintptr, fn func(uintptr)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, UIntPtr := range slice {
		UIntPtr := UIntPtr
		go (func() {
			fn(UIntPtr)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindByte(slice []byte, fn func(byte) bool) []byte {
	sliceCh := make(chan []byte, 1)
	sliceCh <- []byte{}
	for _, Byte := range slice {
		Byte := Byte
		go (func() {
			if fn(Byte) {
				sliceCh <- append(<-sliceCh, Byte)
			}
		})()
	}
	return <-sliceCh
}

func MapByte(slice []byte, fn func(byte) byte) []byte {
	sliceCh := make(chan []byte, 1)
	sliceCh <- []byte{}
	for _, Byte := range slice {
		Byte := Byte
		go (func() {
			v := fn(Byte)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachByte(slice []byte, fn func(byte)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Byte := range slice {
		Byte := Byte
		go (func() {
			fn(Byte)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindRune(slice []rune, fn func(rune) bool) []rune {
	sliceCh := make(chan []rune, 1)
	sliceCh <- []rune{}
	for _, Rune := range slice {
		Rune := Rune
		go (func() {
			if fn(Rune) {
				sliceCh <- append(<-sliceCh, Rune)
			}
		})()
	}
	return <-sliceCh
}

func MapRune(slice []rune, fn func(rune) rune) []rune {
	sliceCh := make(chan []rune, 1)
	sliceCh <- []rune{}
	for _, Rune := range slice {
		Rune := Rune
		go (func() {
			v := fn(Rune)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachRune(slice []rune, fn func(rune)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Rune := range slice {
		Rune := Rune
		go (func() {
			fn(Rune)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindFloat32(slice []float32, fn func(float32) bool) []float32 {
	sliceCh := make(chan []float32, 1)
	sliceCh <- []float32{}
	for _, Float32 := range slice {
		Float32 := Float32
		go (func() {
			if fn(Float32) {
				sliceCh <- append(<-sliceCh, Float32)
			}
		})()
	}
	return <-sliceCh
}

func MapFloat32(slice []float32, fn func(float32) float32) []float32 {
	sliceCh := make(chan []float32, 1)
	sliceCh <- []float32{}
	for _, Float32 := range slice {
		Float32 := Float32
		go (func() {
			v := fn(Float32)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachFloat32(slice []float32, fn func(float32)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Float32 := range slice {
		Float32 := Float32
		go (func() {
			fn(Float32)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindFloat64(slice []float64, fn func(float64) bool) []float64 {
	sliceCh := make(chan []float64, 1)
	sliceCh <- []float64{}
	for _, Float64 := range slice {
		Float64 := Float64
		go (func() {
			if fn(Float64) {
				sliceCh <- append(<-sliceCh, Float64)
			}
		})()
	}
	return <-sliceCh
}

func MapFloat64(slice []float64, fn func(float64) float64) []float64 {
	sliceCh := make(chan []float64, 1)
	sliceCh <- []float64{}
	for _, Float64 := range slice {
		Float64 := Float64
		go (func() {
			v := fn(Float64)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachFloat64(slice []float64, fn func(float64)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Float64 := range slice {
		Float64 := Float64
		go (func() {
			fn(Float64)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindComplex64(slice []complex64, fn func(complex64) bool) []complex64 {
	sliceCh := make(chan []complex64, 1)
	sliceCh <- []complex64{}
	for _, Complex64 := range slice {
		Complex64 := Complex64
		go (func() {
			if fn(Complex64) {
				sliceCh <- append(<-sliceCh, Complex64)
			}
		})()
	}
	return <-sliceCh
}

func MapComplex64(slice []complex64, fn func(complex64) complex64) []complex64 {
	sliceCh := make(chan []complex64, 1)
	sliceCh <- []complex64{}
	for _, Complex64 := range slice {
		Complex64 := Complex64
		go (func() {
			v := fn(Complex64)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachComplex64(slice []complex64, fn func(complex64)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Complex64 := range slice {
		Complex64 := Complex64
		go (func() {
			fn(Complex64)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func FindComplex128(slice []complex128, fn func(complex128) bool) []complex128 {
	sliceCh := make(chan []complex128, 1)
	sliceCh <- []complex128{}
	for _, Complex128 := range slice {
		Complex128 := Complex128
		go (func() {
			if fn(Complex128) {
				sliceCh <- append(<-sliceCh, Complex128)
			}
		})()
	}
	return <-sliceCh
}

func MapComplex128(slice []complex128, fn func(complex128) complex128) []complex128 {
	sliceCh := make(chan []complex128, 1)
	sliceCh <- []complex128{}
	for _, Complex128 := range slice {
		Complex128 := Complex128
		go (func() {
			v := fn(Complex128)
			sliceCh <- append(<-sliceCh, v)
		})()
	}
	return <-sliceCh
}

func EachComplex128(slice []complex128, fn func(complex128)) {
	ch := make(chan bool, 1)
	ch <- true
	for _, Complex128 := range slice {
		Complex128 := Complex128
		go (func() {
			fn(Complex128)
			<-ch
			ch <- true
		})()
	}
	<-ch
}

func Find(slice interface{}, fn interface{}) interface{} {
	f := reflect.ValueOf(fn)
	a := reflect.ValueOf(slice)
	bCh := make(chan reflect.Value, 1)
	bCh <- reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)
	for i := 0; i < a.Len(); i++ {
		v := a.Index(i)
		go (func() {
			if f.Call([]reflect.Value{v})[0].Bool() {
				bCh <- reflect.Append(<-bCh, v)
			}
		})()
	}
	return (<-bCh).Interface()
}

func Map(slice interface{}, fn interface{}) interface{} {
	f := reflect.ValueOf(fn)
	a := reflect.ValueOf(slice)
	bCh := make(chan reflect.Value, 1)
	bCh <- reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(fn).Out(0)), 0, 0)
	for i := 0; i < a.Len(); i++ {
		i := i
		go (func() {
			v := f.Call([]reflect.Value{a.Index(i)})[0]
			bCh <- reflect.Append(<-bCh, v)
		})()
	}
	return (<-bCh).Interface()
}

func Each(slice interface{}, fn interface{}) {
	f := reflect.ValueOf(fn)
	s := reflect.ValueOf(slice)
	ch := make(chan bool, 1)
	ch <- true
	for i := 0; i < s.Len(); i++ {
		i := i
		go (func() {
			f.Call([]reflect.Value{s.Index(i)})
			<-ch
			ch <- true
		})()
	}
	<-ch
}
