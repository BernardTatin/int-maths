/*
 * Factorial.go
 *
 * int_maths.Factorial.go
 *
 * project: int-maths
 * by:      bernard
 * created: 2022-06-24
 */

package int_maths

func Fact[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](N V) (V, bool) {
	var inner_fact func(V, V) (V, bool)
	inner_fact = func(accum V, k V) (V, bool) {
		if k > 1 {
			var new_accum, ok = SMul(k, accum)
			if ok {
				return inner_fact(new_accum, k-1)
			} else {
				return 1, false
			}
		} else {
			return accum, true
		}
	}
	return inner_fact(1, N)
}
