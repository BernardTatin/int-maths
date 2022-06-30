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

func is_signed[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a V) (bool, int64, uint64) {
	var k = reflect.TypeOf(a).Kind()
	switch k {
	case reflect.Int8:
		return true, math.MinInt8, math.MaxInt8
	case reflect.Int16:
		return true, math.MinInt16, math.MaxInt16
	case reflect.Int32:
		return true, math.MinInt32, math.MaxInt32
	case reflect.Int64:
		return true, math.MinInt64, math.MaxInt64

	case reflect.Uint8:
		return false, 0, math.MaxUint8
	case reflect.Uint16:
		return false, 0, math.MaxUint16
	case reflect.Uint32:
		return false, 0, math.MaxUint32
	case reflect.Uint64:
		return false, 0, math.MaxUint64
	}
	return false, 0, 0
}
func SSub[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a V, b V) (V, bool) {
	var isub func(a V, b V) (V, bool)
	var iadd func(a V, b V) (V, bool)
	var signed, smini0, smaxi0 = is_signed(a)
	var smini V = V(smini0)
	var smaxi V = V(smaxi0)
	var A, B V = -a, -b

	isub = func(a V, b V) (V, bool) {
		var good bool

		if signed {
			// something is misunderstood there!
			good = a-smini >= b
		} else {
			good = a >= b
		}
		return a - b, good
	}
	iadd = func(a V, b V) (V, bool) {
		var good bool

		good = a <= smaxi-b
		return a + b, good
	}
	/*
	   A = abs(a), B = abs(b)
	   ! ------ ! a == 0 ! a < 0     ! a > 0  !
	   ! b == 0 ! 0      ! a         ! a      !
	   ! b < 0  ! -b     ! B - A     ! A + B  !
	   ! b > 0  ! -b     ! -(A + b)  ! *****  !

	   a > 0 && b > 0
	   ! a > b ! a == b ! b > a    !
	   ! a - b ! 0      ! -(b - a) !
	*/
	// fmt.Printf("SSub %d, %d - %d\n", a, b, -b)
	if b == 0 {
		return a, true
	} else if signed {
		if b == smini {
			return 0, false
		}
		if a == smini && b > 0 {
			return 0, false
		}
		if a == smini && b < 0 {
			return iadd(a, B)
		}
		if a < 0 {
			if b < 0 {
				// -A - (-B) -> B - A
				return isub(B, A)
			} else {
				// -A - B -> -(A + B)
				c, ok := iadd(A, B)
				return -c, ok
			}
		} else if a > 0 {
			if b < 0 {
				// A - (-B) -> A + B
				return iadd(A, B)
			} else {
				// A - B
				return isub(A, B)
			}
		} else if a == 0 {
			return -b, true
		}
	} else {
		if a < b {
			return 0, false
		} else {
			return isub(a, b)
		}
	}
	return isub(a, b)
}

func SAdd[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](a V, b V) (V, bool) {
	var isub func(a V, b V) (V, bool)
	var iadd func(a V, b V) (V, bool)
	var signed, smini0, smaxi0 = is_signed(a)
	var smini V = V(smini0)
	var smaxi V = V(smaxi0)

	iadd = func(a V, b V) (V, bool) {
		var good bool

		good = a <= smaxi-b
		return a + b, good
	}
	isub = func(a V, b V) (V, bool) {
		var good bool

		good = a-smini >= b
		return a - b, good
	}
	if a == 0 {
		return b, true
	} else if b == 0 {
		return a, true
	} else if signed {
		if a > 0 && b < 0 {
			if smini == b {
				return 0, false
			}
			return isub(a, -b)
		} else if a < 0 && b > 0 {
			if smini == a {
				return 0, false
			}
			return isub(b, -a)
		} else if a < 0 && b < 0 {
			if smini == a || smini == b {
				return 0, false
			}
			c, ok := iadd(-a, -b)
			return -c, ok
		} else {
			return iadd(a, b)
		}
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
