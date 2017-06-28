package f

import (
	"reflect"
	"sync"
)

func IsBool(slice []bool) bool {
	return len(slice) > 0
}

func FromBool(slice []bool, b bool) bool {
	if IsBool(slice) {
		return slice[0]
	}
	return b
}

func FindBool(slice []bool, fn func(bool) bool) []bool {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []bool{}
	for _, Bool := range slice {
		Bool := Bool
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Bool) {
				mux.Lock()
				buf = append(buf, Bool)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapBool(slice []bool, fn func(bool) bool) []bool {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []bool{}
	for _, Bool := range slice {
		Bool := Bool
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Bool)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachBool(slice []bool, fn func(bool)) {
	var wg sync.WaitGroup
	for _, Bool := range slice {
		Bool := Bool
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Bool)
		})()
	}
	wg.Wait()
}

func IsString(slice []string) bool {
	return len(slice) > 0
}

func FromString(slice []string, s string) string {
	if IsString(slice) {
		return slice[0]
	}
	return s
}

func FindString(slice []string, fn func(string) bool) []string {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []string{}
	for _, String := range slice {
		String := String
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(String) {
				mux.Lock()
				buf = append(buf, String)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapString(slice []string, fn func(string) string) []string {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []string{}
	for _, String := range slice {
		String := String
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(String)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachString(slice []string, fn func(string)) {
	var wg sync.WaitGroup
	for _, String := range slice {
		String := String
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(String)
		})()
	}
	wg.Wait()
}

func IsInt(slice []int) bool {
	return len(slice) > 0
}

func FromInt(slice []int, i int) int {
	if IsInt(slice) {
		return slice[0]
	}
	return i
}

func FindInt(slice []int, fn func(int) bool) []int {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int{}
	for _, Int := range slice {
		Int := Int
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Int) {
				mux.Lock()
				buf = append(buf, Int)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapInt(slice []int, fn func(int) int) []int {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int{}
	for _, Int := range slice {
		Int := Int
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Int)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachInt(slice []int, fn func(int)) {
	var wg sync.WaitGroup
	for _, Int := range slice {
		Int := Int
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Int)
		})()
	}
	wg.Wait()
}

func IsInt8(slice []int8) bool {
	return len(slice) > 0
}

func FromInt8(slice []int8, i int8) int8 {
	if IsInt8(slice) {
		return slice[0]
	}
	return i
}

func FindInt8(slice []int8, fn func(int8) bool) []int8 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int8{}
	for _, Int8 := range slice {
		Int8 := Int8
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Int8) {
				mux.Lock()
				buf = append(buf, Int8)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapInt8(slice []int8, fn func(int8) int8) []int8 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int8{}
	for _, Int8 := range slice {
		Int8 := Int8
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Int8)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachInt8(slice []int8, fn func(int8)) {
	var wg sync.WaitGroup
	for _, Int8 := range slice {
		Int8 := Int8
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Int8)
		})()
	}
	wg.Wait()
}

func IsInt16(slice []int16) bool {
	return len(slice) > 0
}

func FromInt16(slice []int16, i int16) int16 {
	if IsInt16(slice) {
		return slice[0]
	}
	return i
}

func FindInt16(slice []int16, fn func(int16) bool) []int16 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int16{}
	for _, Int16 := range slice {
		Int16 := Int16
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Int16) {
				mux.Lock()
				buf = append(buf, Int16)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapInt16(slice []int16, fn func(int16) int16) []int16 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int16{}
	for _, Int16 := range slice {
		Int16 := Int16
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Int16)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachInt16(slice []int16, fn func(int16)) {
	var wg sync.WaitGroup
	for _, Int16 := range slice {
		Int16 := Int16
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Int16)
		})()
	}
	wg.Wait()
}

func IsInt32(slice []int32) bool {
	return len(slice) > 0
}

func FromInt32(slice []int32, i int32) int32 {
	if IsInt32(slice) {
		return slice[0]
	}
	return i
}

func FindInt32(slice []int32, fn func(int32) bool) []int32 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int32{}
	for _, Int32 := range slice {
		Int32 := Int32
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Int32) {
				mux.Lock()
				buf = append(buf, Int32)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapInt32(slice []int32, fn func(int32) int32) []int32 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int32{}
	for _, Int32 := range slice {
		Int32 := Int32
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Int32)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachInt32(slice []int32, fn func(int32)) {
	var wg sync.WaitGroup
	for _, Int32 := range slice {
		Int32 := Int32
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Int32)
		})()
	}
	wg.Wait()
}

func IsInt64(slice []int64) bool {
	return len(slice) > 0
}

func FromInt64(slice []int64, i int64) int64 {
	if IsInt64(slice) {
		return slice[0]
	}
	return i
}

func FindInt64(slice []int64, fn func(int64) bool) []int64 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int64{}
	for _, Int64 := range slice {
		Int64 := Int64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Int64) {
				mux.Lock()
				buf = append(buf, Int64)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapInt64(slice []int64, fn func(int64) int64) []int64 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []int64{}
	for _, Int64 := range slice {
		Int64 := Int64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Int64)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachInt64(slice []int64, fn func(int64)) {
	var wg sync.WaitGroup
	for _, Int64 := range slice {
		Int64 := Int64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Int64)
		})()
	}
	wg.Wait()
}

func IsUInt(slice []uint) bool {
	return len(slice) > 0
}

func FromUInt(slice []uint, u uint) uint {
	if IsUInt(slice) {
		return slice[0]
	}
	return u
}

func FindUInt(slice []uint, fn func(uint) bool) []uint {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint{}
	for _, Uint := range slice {
		Uint := Uint
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Uint) {
				mux.Lock()
				buf = append(buf, Uint)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapUInt(slice []uint, fn func(uint) uint) []uint {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint{}
	for _, Uint := range slice {
		Uint := Uint
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Uint)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachUInt(slice []uint, fn func(uint)) {
	var wg sync.WaitGroup
	for _, UInt := range slice {
		UInt := UInt
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(UInt)
		})()
	}
	wg.Wait()
}

func IsUInt8(slice []uint8) bool {
	return len(slice) > 0
}

func FromUInt8(slice []uint8, u uint8) uint8 {
	if IsUInt8(slice) {
		return slice[0]
	}
	return u
}

func FindUInt8(slice []uint8, fn func(uint8) bool) []uint8 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint8{}
	for _, Uint8 := range slice {
		Uint8 := Uint8
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Uint8) {
				mux.Lock()
				buf = append(buf, Uint8)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapUInt8(slice []uint8, fn func(uint8) uint8) []uint8 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint8{}
	for _, Uint8 := range slice {
		Uint8 := Uint8
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Uint8)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachUInt8(slice []uint8, fn func(uint8)) {
	var wg sync.WaitGroup
	for _, UInt8 := range slice {
		UInt8 := UInt8
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(UInt8)
		})()
	}
	wg.Wait()
}

func IsUInt16(slice []uint16) bool {
	return len(slice) > 0
}

func FromUInt16(slice []uint16, u uint16) uint16 {
	if IsUInt16(slice) {
		return slice[0]
	}
	return u
}

func FindUInt16(slice []uint16, fn func(uint16) bool) []uint16 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint16{}
	for _, Uint16 := range slice {
		Uint16 := Uint16
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Uint16) {
				mux.Lock()
				buf = append(buf, Uint16)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapUInt16(slice []uint16, fn func(uint16) uint16) []uint16 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint16{}
	for _, Uint16 := range slice {
		Uint16 := Uint16
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Uint16)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachUInt16(slice []uint16, fn func(uint16)) {
	var wg sync.WaitGroup
	for _, UInt16 := range slice {
		UInt16 := UInt16
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(UInt16)
		})()
	}
	wg.Wait()
}

func IsUInt32(slice []uint32) bool {
	return len(slice) > 0
}

func FromUInt32(slice []uint32, u uint32) uint32 {
	if IsUInt32(slice) {
		return slice[0]
	}
	return u
}

func FindUInt32(slice []uint32, fn func(uint32) bool) []uint32 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint32{}
	for _, Uint32 := range slice {
		Uint32 := Uint32
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Uint32) {
				mux.Lock()
				buf = append(buf, Uint32)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapUInt32(slice []uint32, fn func(uint32) uint32) []uint32 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint32{}
	for _, Uint32 := range slice {
		Uint32 := Uint32
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Uint32)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachUInt32(slice []uint32, fn func(uint32)) {
	var wg sync.WaitGroup
	for _, UInt32 := range slice {
		UInt32 := UInt32
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(UInt32)
		})()
	}
	wg.Wait()
}

func IsUInt64(slice []uint64) bool {
	return len(slice) > 0
}

func FromUInt64(slice []uint64, u uint64) uint64 {
	if IsUInt64(slice) {
		return slice[0]
	}
	return u
}

func FindUInt64(slice []uint64, fn func(uint64) bool) []uint64 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint64{}
	for _, Uint64 := range slice {
		Uint64 := Uint64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Uint64) {
				mux.Lock()
				buf = append(buf, Uint64)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapUInt64(slice []uint64, fn func(uint64) uint64) []uint64 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uint64{}
	for _, Uint64 := range slice {
		Uint64 := Uint64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Uint64)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachUInt64(slice []uint64, fn func(uint64)) {
	var wg sync.WaitGroup
	for _, UInt64 := range slice {
		UInt64 := UInt64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(UInt64)
		})()
	}
	wg.Wait()
}

func IsUIntPtr(slice []uintptr) bool {
	return len(slice) > 0
}

func FromUIntPtr(slice []uintptr, u uintptr) uintptr {
	if IsUIntPtr(slice) {
		return slice[0]
	}
	return u
}

func FindUIntPtr(slice []uintptr, fn func(uintptr) bool) []uintptr {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uintptr{}
	for _, Uintptr := range slice {
		Uintptr := Uintptr
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Uintptr) {
				mux.Lock()
				buf = append(buf, Uintptr)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapUIntPtr(slice []uintptr, fn func(uintptr) uintptr) []uintptr {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []uintptr{}
	for _, Uintptr := range slice {
		Uintptr := Uintptr
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Uintptr)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachUIntPtr(slice []uintptr, fn func(uintptr)) {
	var wg sync.WaitGroup
	for _, UIntPtr := range slice {
		UIntPtr := UIntPtr
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(UIntPtr)
		})()
	}
	wg.Wait()
}

func IsByte(slice []byte) bool {
	return len(slice) > 0
}

func FromByte(slice []byte, b byte) byte {
	if IsByte(slice) {
		return slice[0]
	}
	return b
}

func FindByte(slice []byte, fn func(byte) bool) []byte {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []byte{}
	for _, Byte := range slice {
		Byte := Byte
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Byte) {
				mux.Lock()
				buf = append(buf, Byte)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapByte(slice []byte, fn func(byte) byte) []byte {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []byte{}
	for _, Byte := range slice {
		Byte := Byte
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Byte)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachByte(slice []byte, fn func(byte)) {
	var wg sync.WaitGroup
	for _, Byte := range slice {
		Byte := Byte
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Byte)
		})()
	}
	wg.Wait()
}

func IsRune(slice []rune) bool {
	return len(slice) > 0
}

func FromRune(slice []rune, r rune) rune {
	if IsRune(slice) {
		return slice[0]
	}
	return r
}

func FindRune(slice []rune, fn func(rune) bool) []rune {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []rune{}
	for _, Rune := range slice {
		Rune := Rune
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Rune) {
				mux.Lock()
				buf = append(buf, Rune)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapRune(slice []rune, fn func(rune) rune) []rune {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []rune{}
	for _, Rune := range slice {
		Rune := Rune
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Rune)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachRune(slice []rune, fn func(rune)) {
	var wg sync.WaitGroup
	for _, Rune := range slice {
		Rune := Rune
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Rune)
		})()
	}
	wg.Wait()
}

func IsFloat32(slice []float32) bool {
	return len(slice) > 0
}

func FromFloat32(slice []float32, f float32) float32 {
	if IsFloat32(slice) {
		return slice[0]
	}
	return f
}

func FindFloat32(slice []float32, fn func(float32) bool) []float32 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []float32{}
	for _, Float32 := range slice {
		Float32 := Float32
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Float32) {
				mux.Lock()
				buf = append(buf, Float32)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapFloat32(slice []float32, fn func(float32) float32) []float32 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []float32{}
	for _, Float32 := range slice {
		Float32 := Float32
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Float32)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachFloat32(slice []float32, fn func(float32)) {
	var wg sync.WaitGroup
	for _, Float32 := range slice {
		Float32 := Float32
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Float32)
		})()
	}
	wg.Wait()
}

func IsFloat64(slice []float64) bool {
	return len(slice) > 0
}

func FromFloat64(slice []float64, f float64) float64 {
	if IsFloat64(slice) {
		return slice[0]
	}
	return f
}

func FindFloat64(slice []float64, fn func(float64) bool) []float64 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []float64{}
	for _, Float64 := range slice {
		Float64 := Float64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Float64) {
				mux.Lock()
				buf = append(buf, Float64)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapFloat64(slice []float64, fn func(float64) float64) []float64 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []float64{}
	for _, Float64 := range slice {
		Float64 := Float64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Float64)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachFloat64(slice []float64, fn func(float64)) {
	var wg sync.WaitGroup
	for _, Float64 := range slice {
		Float64 := Float64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Float64)
		})()
	}
	wg.Wait()
}

func IsComplex64(slice []complex64) bool {
	return len(slice) > 0
}

func FromComplex64(slice []complex64, c complex64) complex64 {
	if IsComplex64(slice) {
		return slice[0]
	}
	return c
}

func FindComplex64(slice []complex64, fn func(complex64) bool) []complex64 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []complex64{}
	for _, Complex64 := range slice {
		Complex64 := Complex64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Complex64) {
				mux.Lock()
				buf = append(buf, Complex64)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapComplex64(slice []complex64, fn func(complex64) complex64) []complex64 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []complex64{}
	for _, Complex64 := range slice {
		Complex64 := Complex64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Complex64)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachComplex64(slice []complex64, fn func(complex64)) {
	var wg sync.WaitGroup
	for _, Complex64 := range slice {
		Complex64 := Complex64
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Complex64)
		})()
	}
	wg.Wait()
}

func IsComplex128(slice []complex128) bool {
	return len(slice) > 0
}

func FromComplex128(slice []complex128, c complex128) complex128 {
	if IsComplex128(slice) {
		return slice[0]
	}
	return c
}

func FindComplex128(slice []complex128, fn func(complex128) bool) []complex128 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []complex128{}
	for _, Complex128 := range slice {
		Complex128 := Complex128
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if fn(Complex128) {
				mux.Lock()
				buf = append(buf, Complex128)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return buf
}

func MapComplex128(slice []complex128, fn func(complex128) complex128) []complex128 {
	var mux sync.Mutex
	var wg sync.WaitGroup
	buf := []complex128{}
	for _, Complex128 := range slice {
		Complex128 := Complex128
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := fn(Complex128)
			mux.Lock()
			buf = append(buf, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return buf
}

func EachComplex128(slice []complex128, fn func(complex128)) {
	var wg sync.WaitGroup
	for _, Complex128 := range slice {
		Complex128 := Complex128
		wg.Add(1)
		go (func() {
			defer wg.Done()
			fn(Complex128)
		})()
	}
	wg.Wait()
}

func Is(slice interface{}) bool {
	a := reflect.ValueOf(slice)
	return a.Len() > 0
}

func From(slice interface{}, v interface{}) interface{} {
	if Is(slice) {
		return reflect.ValueOf(slice).Index(0).Interface()
	}
	return reflect.ValueOf(v).Interface()
}

func Find(slice interface{}, fn interface{}) interface{} {
	f := reflect.ValueOf(fn)
	a := reflect.ValueOf(slice)
	b := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)
	var mux sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < a.Len(); i++ {
		v := a.Index(i)
		wg.Add(1)
		go (func() {
			defer wg.Done()
			if f.Call([]reflect.Value{v})[0].Bool() {
				mux.Lock()
				b = reflect.Append(b, v)
				mux.Unlock()
			}
		})()
	}
	wg.Wait()
	return b.Interface()
}

func Map(slice interface{}, fn interface{}) interface{} {
	f := reflect.ValueOf(fn)
	a := reflect.ValueOf(slice)
	b := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(fn).Out(0)), 0, 0)
	var mux sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < a.Len(); i++ {
		i := i
		wg.Add(1)
		go (func() {
			defer wg.Done()
			v := f.Call([]reflect.Value{a.Index(i)})[0]
			mux.Lock()
			b = reflect.Append(b, v)
			mux.Unlock()
		})()
	}
	wg.Wait()
	return b.Interface()
}

func Each(slice interface{}, fn interface{}) {
	f := reflect.ValueOf(fn)
	s := reflect.ValueOf(slice)
	var wg sync.WaitGroup
	for i := 0; i < s.Len(); i++ {
		i := i
		wg.Add(1)
		go (func() {
			defer wg.Done()
			f.Call([]reflect.Value{s.Index(i)})
		})()
	}
	wg.Wait()
}
