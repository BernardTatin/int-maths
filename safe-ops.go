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

func SafeAdd[V int32 | int64 | uint32 | uint64](a V, b V) (V, bool) {
	c := a + b
	t := reflect.TypeOf(a).Kind()
	var good bool
	switch t {
	case reflect.Int32:
		good = int32(a) <= math.MaxInt32-int32(b)
	case reflect.Int64:
		good = int64(a) <= math.MaxInt64-int64(b)
	case reflect.Uint32:
		good = uint32(a) <= math.MaxUint32-uint32(b)
	case reflect.Uint64:
		good = uint64(a) <= math.MaxUint64-uint64(b)
	}
	return c, good
}

func SafeMul[V int32 | int64 | uint32 | uint64](a V, b V) (V, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	c := a * b
	t := reflect.TypeOf(a).Kind()
	var good bool
	switch t {
	case reflect.Int32:
		good = int32(a) <= math.MaxInt32/int32(b)
	case reflect.Int64:
		good = int64(a) <= math.MaxInt64/int64(b)
	case reflect.Uint32:
		good = uint32(a) <= math.MaxUint32/uint32(b)
	case reflect.Uint64:
		good = uint64(a) <= math.MaxUint64/uint64(b)
	}
	return c, good
}
