/*
 * sadd_test.go
 *
 * tests.sadd_test.go
 *
 * project: int-maths
 * by:      bernard
 * created: 2022-06-25
 *
 * Usage:
 * go tests -v -timeout 0 [-cpu 4]
 */

package tests

import (
	int_maths "github.com/BernardTatin/int-maths"
	"math"
	"reflect"
	"testing"
)

func test_add[V int8 | int16 | int32 | uint8 | uint16 | uint32](t *testing.T, from V, to V) (int64, int64) {
	var k = reflect.TypeOf(from).Name()
	var aFrom = int64(from)
	var aTo = int64(to)
	var bFrom = int64(from)
	var bTo = int64(to)
	var count, errors, goods int64 = 0, 0, 0
	var i, j int64

	for i = aFrom; i <= aTo; i++ {
		for j = bFrom; j <= bTo; j++ {
			var computed V
			var ok bool
			var c int64 = i + j

			count++
			computed, ok = int_maths.SAdd(V(i), V(j))
			computed64 := int64(computed)
			if c != computed64 {
				if ok {
					errors++
					t.Logf("%s BAD %s ok true  %4d != %4d %4d + %4d\n",
						t.Name(), k, computed64, c, i, j)
				} else {
					goods++
				}
			} else {
				goods++
			}
		}
	}
	return goods, errors
}

func TestSAddI8(t *testing.T) {
	var goods, bads = test_add(t, int8(math.MinInt8), int8(math.MaxInt8))

	end_of_test(t, goods, bads)
}

func TestSAddU8(t *testing.T) {
	var goods, bads = test_add(t, uint8(0), uint8(math.MaxUint8))

	end_of_test(t, goods, bads)
}

func TestSAddI16(t *testing.T) {
	// var goods, bads = test_add(t, int16(math.MinInt16), int16(math.MaxInt16))
	var goods, bads = test_add(t, int16(-i16_cpt), int16(i16_cpt))

	end_of_test(t, goods, bads)
}

func TestSAddU16(t *testing.T) {
	// var goods, bads = test_add(t, uint16(0), uint16(math.MaxUint16))
	var goods, bads = test_add(t, uint16(0), uint16(i16_cpt))

	end_of_test(t, goods, bads)
}
