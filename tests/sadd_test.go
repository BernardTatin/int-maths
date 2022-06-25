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
	var count, errors = 0, 0
	var i, j int16

	for i = aFrom; i < aTo; i++ {
		for j = bFrom; j < bTo; j++ {
			var computed int8
			var ok bool
			var c int16 = i + j
			var good bool = c >= int16(math.MinInt8) && c <= int16(math.MaxInt8)

			count++
			computed, ok = int_maths.SAdd(int8(i), int8(j))
			if !good {
				if ok {
					fmt.Printf("%d + %d gives %d instead of %d\n",
						i, j, computed, c)
					errors++
				}
			} else {
				if !ok {
					fmt.Printf("%d + %d gives %d instead of %d\n",
						i, j, computed, c)
					errors++
				}
			}
		}
	}
	fmt.Printf("Tests: %d, errors: %d\n", count, errors)
}
