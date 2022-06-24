/*
 * fibo.go
 *
 * int_maths.fibo.go
 *
 * project: int-maths
 * by:      bernard
 * created: 2022-06-24
 */

package int_maths

func Fibo[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](N V) (V, bool) {
	if N < 0 {
		return 0, false
	}
	switch N {
	case 0:
		return 0, true
	case 1:
		return 1, true
	case 2:
		return 1, true
	}
	var i, N1, N2 V
	N1 = 1
	N2 = 1
	for i = 2; i < N; i++ {
		result, ok := SAdd(N1, N2)
		if !ok {
			return result, false
		}
		N1 = N2
		N2 = result
	}
	return N2, true
}
