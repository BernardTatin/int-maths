/*
 * safe-ops.go
 *
 * int_maths.safe-ops.go
 *
 * project: int-maths
 * by:      bernard
 * created: 2022-06-23
 */

package int_maths

import (
	"math"
	"reflect"
)

func SSub[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a V, b V) (V, bool) {
	var isub func(a V, b V) (V, bool)

	isub = func(a V, b V) (V, bool) {
		var good bool
		var t = reflect.TypeOf(a).Kind()
		switch t {
		case reflect.Int8:
			good = int8(a) <= math.MaxInt8-int8(b)
		case reflect.Int16:
			good = int16(a) <= math.MaxInt16-int16(b)
		case reflect.Int32:
			good = int32(a) <= math.MaxInt32-int32(b)
		case reflect.Int64:
			good = int64(a) <= math.MaxInt64-int64(b)
		case reflect.Uint8:
			good = uint8(a) <= math.MaxUint8-uint8(b)
		case reflect.Uint16:
			good = uint16(a) <= math.MaxUint16-uint16(b)
		case reflect.Uint32:
			good = uint32(a) <= math.MaxUint32-uint32(b)
		case reflect.Uint64:
			good = uint64(a) <= math.MaxUint64-uint64(b)
		}
		return a + b, good
	}
	if a == 0 {
		return -b, true
	} else if b == 0 {
		return a, true
	} else if a < b {
		c, good := isub(b, a)
		return -c, good
	} else {
		return isub(a, b)
	}
}

func SAdd[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a V, b V) (V, bool) {
	var iadd func(a V, b V) (V, bool)

	iadd = func(a V, b V) (V, bool) {
		var good bool
		var t = reflect.TypeOf(a).Kind()
		switch t {
		case reflect.Int8:
			good = int8(a) <= math.MaxInt8-int8(b)
		case reflect.Int16:
			good = int16(a) <= math.MaxInt16-int16(b)
		case reflect.Int32:
			good = int32(a) <= math.MaxInt32-int32(b)
		case reflect.Int64:
			good = int64(a) <= math.MaxInt64-int64(b)
		case reflect.Uint8:
			good = uint8(a) <= math.MaxUint8-uint8(b)
		case reflect.Uint16:
			good = uint16(a) <= math.MaxUint16-uint16(b)
		case reflect.Uint32:
			good = uint32(a) <= math.MaxUint32-uint32(b)
		case reflect.Uint64:
			good = uint64(a) <= math.MaxUint64-uint64(b)
		}
		return a + b, good
	}
	if a == 0 {
		return b, true
	} else if b == 0 {
		return a, true
	} else if a > 0 && b < 0 {
		return SSub(a, b)
	} else if a < 0 && b > 0 {
		return SSub(a, b)
	} else {
		return iadd(a, b)
	}
}

func SMul[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a V, b V) (V, bool) {
	var imul func(a V, b V) (V, bool)

	imul = func(a V, b V) (V, bool) {
		var good bool
		var t = reflect.TypeOf(a).Kind()

		switch t {
		case reflect.Int8:
			good = int8(a) <= math.MaxInt8/int8(b)
		case reflect.Int16:
			good = int16(a) <= math.MaxInt16/int16(b)
		case reflect.Int32:
			good = int32(a) <= math.MaxInt32/int32(b)
		case reflect.Int64:
			good = int64(a) <= math.MaxInt64/int64(b)
		case reflect.Uint8:
			good = uint8(a) <= math.MaxUint8/uint8(b)
		case reflect.Uint16:
			good = uint16(a) <= math.MaxUint16/uint16(b)
		case reflect.Uint32:
			good = uint32(a) <= math.MaxUint32/uint32(b)
		case reflect.Uint64:
			good = uint64(a) <= math.MaxUint64/uint64(b)
		}
		return a * b, good
	}
	if a == 0 || b == 0 {
		return 0, true
	} else {
		return imul(a, b)
	}
}
