/*
 * sadd_test.go
 *
 * tests.sadd_test.go
 *
 * project: int-maths
 * by:      bernard
 * created: 2022-06-25
 */

package tests

import (
	"fmt"
	int_maths "github.com/BernardTatin/int-maths"
	"math"
	"testing"
)

func TestSAddI8(t *testing.T) {
	var aFrom int16 = int16(math.MinInt8)
	var aTo int16 = int16(math.MaxInt8)
	var bFrom int16 = int16(math.MinInt8)
	var bTo int16 = int16(math.MaxInt8)
	var count, errors, goods = 0, 0, 0
	var errors1, errors2 = 0, 0
	var goods1, goods2 = 0, 0
	var i int16
	var j int16

	for i = aFrom; i <= aTo; i++ {
		for j = bFrom; j <= bTo; j++ {
			var computed int8
			var ok bool
			var c int16 = i + j

			count++
			computed, ok = int_maths.SAdd(int8(i), int8(j))
			computed16 := int16(computed)
			if c != computed16 {
				if ok {
					errors++
					// good = int8(a) <= math.MaxInt8-int8(b)
					fmt.Printf("BAD1 ok true  %4d != %4d %4d + %4d (%d <= %d)\n",
						computed16, c, i, j, i, math.MaxInt8-int8(j))
					errors1++
				} else {
					goods++
					goods1++
				}
			} else {
				if !ok {
					fmt.Printf("BAD2 ok false  %4d == %4d %4d + %4d\n",
						computed16, c, i, j)
				}
				goods++
				goods2++
			}

		}
	}
	fmt.Printf("Tests: %5d, goods: %5d, goods1: %5d, goods2: %5d\n",
		count, goods, goods1, goods2)
	fmt.Printf("Tests: %5d, errors: %5d, errors1: %5d, errors2: %5d (%5d)\n",
		count, errors, errors1, errors2, errors-errors1-errors2)
	fmt.Printf("Min int8: %d, Max int8: %d\n",
		math.MinInt8, math.MaxInt8)
}
func TestLimits(t *testing.T) {
	var low = int(math.MinInt8)
	var hih = int(math.MaxInt8)
	var low8 int8 = math.MinInt8
	var hih8 int8 = math.MaxInt8

	fmt.Printf("Min: %d/%02x, Min8: %d/%02x\n",
		low, low, low8, low8)
	fmt.Printf("Max: %d/%02x, Max8: %d/%02x\n",
		hih, hih, hih8, hih8)
}
