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
				} else {
					goods++
				}
			} else {
				goods++
			}

		}
	}
	fmt.Printf("Tests TestSAddI8: %5d, goods: %5d errors: %5d\n",
		count, goods, errors)
}

func TestSAddU8(t *testing.T) {
	var aFrom uint16 = 0
	var aTo uint16 = uint16(math.MaxUint8)
	var bFrom uint16 = 0
	var bTo uint16 = uint16(math.MaxUint8)
	var count, errors, goods = 0, 0, 0
	var i uint16
	var j uint16

	for i = aFrom; i <= aTo; i++ {
		for j = bFrom; j <= bTo; j++ {
			var computed uint8
			var ok bool
			var c uint16 = i + j

			count++
			computed, ok = int_maths.SAdd(uint8(i), uint8(j))
			computed16 := uint16(computed)
			if c != computed16 {
				if ok {
					errors++
					// good = uint8(a) <= math.MaxUint8-uint8(b)
					fmt.Printf("BAD1 ok true  %4d != %4d %4d + %4d (%d <= %d)\n",
						computed16, c, i, j, i, math.MaxUint8-uint8(j))
				} else {
					goods++
				}
			} else {
				goods++
			}

		}
	}
	fmt.Printf("Tests TestSAddU8: %5d, goods: %5d errors: %5d\n",
		count, goods, errors)
}
