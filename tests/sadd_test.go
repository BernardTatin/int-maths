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
	var aFrom int16 = int16(math.MinInt8) + 1
	var aTo int16 = int16(math.MaxInt8) - 1
	var bFrom int16 = int16(math.MinInt8) + 1
	var bTo int16 = int16(math.MaxInt8) - 1
	// var aFrom, aTo, bFrom, bTo int16 = -4, 4, -4, 4
	var count, errors, goods = 0, 0, 0
	var errors1 = 0
	var errors2 = 0
	var goods1 = 0
	var goods2 = 0
	var i int16
	var j int16

	for i = aFrom; i < aTo; i++ {
		for j = bFrom; j < bTo; j++ {
			var computed int8
			var ok bool
			var c int16 = i + j

			count++
			computed, ok = int_maths.SAdd(int8(i), int8(j))
			computed16 := int16(computed)
			if c != computed16 {
				if ok {
					errors++
					fmt.Printf("BAD1 ok true  %4d != %4d %4d + %4d\n",
						computed16, c, i, j)
					errors1++
				} else {
					goods++
					goods1++
				}
			} else {
				if !ok {
					errors++
					fmt.Printf("BAD2 ok false  %4d == %4d %4d + %4d\n",
						computed16, c, i, j)
					errors2++
				} else {
					goods++
					goods2++
				}
			}

		}
	}
	fmt.Printf("Tests: %5d, goods: %5d, goods1: %5d, goods2: %5d\n",
		count, goods, goods1, goods2)
	fmt.Printf("Tests: %5d, errors: %5d, errors1: %5d, errors2: %5d (%5d)\n",
		count, errors, errors1, errors2, errors-errors1-errors2)
}
