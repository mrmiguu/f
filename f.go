package f

func MapBool(boolsA []bool, f func(bool) bool) []bool {
	boolsB := []bool{}
	for _, b := range boolsA {
		boolsB = append(boolsB, f(b))
	}
	return boolsB
}

func EachBool(bools []int, f func(int)) {
	for _, b := range bools {
		f(b)
	}
}

func MapString(stringsA []string, f func(string) string) []string {
	stringsB := []string{}
	for _, b := range stringsA {
		stringsB = append(stringsB, f(b))
	}
	return stringsB
}

func EachString(strings []int, f func(int)) {
	for _, b := range strings {
		f(b)
	}
}

func MapInt(intsA []int, f func(int) int) []int {
	intsB := []int{}
	for _, i := range intsA {
		intsB = append(intsB, f(i))
	}
	return intsB
}

func EachInt(ints []int, f func(int)) {
	for _, i := range ints {
		f(i)
	}
}

func MapInt8(int8sA []int8, f func(int8) int8) []int8 {
	int8sB := []int8{}
	for _, i := range int8sA {
		int8sB = append(int8sB, f(i))
	}
	return int8sB
}

func EachInt8(int8s []int8, f func(int8)) {
	for _, i := range int8s {
		f(i)
	}
}

func MapInt16(int16sA []int16, f func(int16) int16) []int16 {
	int16sB := []int16{}
	for _, i := range int16sA {
		int16sB = append(int16sB, f(i))
	}
	return int16sB
}

func EachInt16(int16s []int16, f func(int16)) {
	for _, i := range int16s {
		f(i)
	}
}

func MapInt32(int32sA []int32, f func(int32) int32) []int32 {
	int32sB := []int32{}
	for _, i := range int32sA {
		int32sB = append(int32sB, f(i))
	}
	return int32sB
}

func EachInt32(int32s []int32, f func(int32)) {
	for _, i := range int32s {
		f(i)
	}
}

func MapInt64(int64sA []int64, f func(int64) int64) []int64 {
	int64sB := []int64{}
	for _, i := range int64sA {
		int64sB = append(int64sB, f(i))
	}
	return int64sB
}

func EachInt64(int64s []int64, f func(int64)) {
	for _, i := range int64s {
		f(i)
	}
}

func MapUInt(uintsA []uint, f func(uint) uint) []uint {
	uintsB := []uint{}
	for _, i := range uintsA {
		uintsB = append(uintsB, f(i))
	}
	return uintsB
}

func EachUInt(uints []uint, f func(uint)) {
	for _, i := range uints {
		f(i)
	}
}

func MapUInt8(uint8sA []uint8, f func(uint8) uint8) []uint8 {
	uint8sB := []uint8{}
	for _, i := range uint8sA {
		uint8sB = append(uint8sB, f(i))
	}
	return uint8sB
}

func EachUInt8(uint8s []uint8, f func(uint8)) {
	for _, i := range uint8s {
		f(i)
	}
}

func MapUInt16(uint16sA []uint16, f func(uint16) uint16) []uint16 {
	uint16sB := []uint16{}
	for _, i := range uint16sA {
		uint16sB = append(uint16sB, f(i))
	}
	return uint16sB
}

func EachUInt16(uint16s []uint16, f func(uint16)) {
	for _, i := range uint16s {
		f(i)
	}
}

func MapUInt32(uint32sA []uint32, f func(uint32) uint32) []uint32 {
	uint32sB := []uint32{}
	for _, i := range uint32sA {
		uint32sB = append(uint32sB, f(i))
	}
	return uint32sB
}

func EachUInt32(uint32s []uint32, f func(uint32)) {
	for _, i := range uint32s {
		f(i)
	}
}

func MapUInt64(uint64sA []uint64, f func(uint64) uint64) []uint64 {
	uint64sB := []uint64{}
	for _, i := range uint64sA {
		uint64sB = append(uint64sB, f(i))
	}
	return uint64sB
}

func EachUInt64(uint64s []uint64, f func(uint64)) {
	for _, i := range uint64s {
		f(i)
	}
}

func MapUIntPtr(uintptrsA []uintptr, f func(uintptr) uintptr) []uintptr {
	uintptrsB := []uintptr{}
	for _, i := range uintptrsA {
		uintptrsB = append(uintptrsB, f(i))
	}
	return uintptrsB
}

func EachUIntPtr(uintptrs []uintptr, f func(uintptr)) {
	for _, i := range uintptrs {
		f(i)
	}
}

func MapByte(bytesA []byte, f func(byte) byte) []byte {
	bytesB := []byte{}
	for _, i := range bytesA {
		bytesB = append(bytesB, f(i))
	}
	return bytesB
}

func EachByte(bytes []byte, f func(byte)) {
	for _, i := range bytes {
		f(i)
	}
}

func MapRune(runesA []rune, f func(rune) rune) []rune {
	runesB := []rune{}
	for _, i := range runesA {
		runesB = append(runesB, f(i))
	}
	return runesB
}

func EachRune(runes []rune, f func(rune)) {
	for _, i := range runes {
		f(i)
	}
}

func MapFloat32(float32sA []float32, f func(float32) float32) []float32 {
	float32sB := []float32{}
	for _, i := range float32sA {
		float32sB = append(float32sB, f(i))
	}
	return float32sB
}

func EachFloat32(float32s []float32, f func(float32)) {
	for _, i := range float32s {
		f(i)
	}
}

func MapFloat64(float64sA []float64, f func(float64) float64) []float64 {
	float64sB := []float64{}
	for _, i := range float64sA {
		float64sB = append(float64sB, f(i))
	}
	return float64sB
}

func EachFloat64(float64s []float64, f func(float64)) {
	for _, i := range float64s {
		f(i)
	}
}

func MapComplex64(complex64sA []complex64, f func(complex64) complex64) []complex64 {
	complex64sB := []complex64{}
	for _, i := range complex64sA {
		complex64sB = append(complex64sB, f(i))
	}
	return complex64sB
}

func EachComplex64(complex64s []complex64, f func(complex64)) {
	for _, i := range complex64s {
		f(i)
	}
}

func MapComplex128(complex128sA []complex128, f func(complex128) complex128) []complex128 {
	complex128sB := []complex128{}
	for _, i := range complex128sA {
		complex128sB = append(complex128sB, f(i))
	}
	return complex128sB
}

func EachComplex128(complex128s []complex128, f func(complex128)) {
	for _, i := range complex128s {
		f(i)
	}
}

func Map(interfacesA []interface{}, f func(interface{}) interface{}) []interface{} {
	interfacesB := make([]interface{}, 0)
	for _, i := range interfacesA {
		interfacesB = append(interfacesB, f(i))
	}
	return interfacesB
}

func Each(interfaces []interface{}, f func(interface{})) {
	for _, i := range interfaces {
		f(i)
	}
}
