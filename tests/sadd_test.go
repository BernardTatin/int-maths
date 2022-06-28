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
	"reflect"
	"testing"
)

const verbose bool = false
const verbose_mask int64 = 0xffff

func test_add[V int8 | int16 | int32 | uint8 | uint16 | uint32](from V, to V) (int64, int64) {
	var k = reflect.TypeOf(from).Name()
	var aFrom = int64(from)
	var aTo = int64(to)
	var bFrom = int64(from)
	var bTo = int64(to)
	var count, errors, goods int64 = 0, 0, 0
	var i, j int64
	var cpt int64 = 0

	for i = aFrom; i <= aTo; i++ {
		for j = bFrom; j <= bTo; j++ {
			var computed V
			var ok bool
			var c int64 = i + j

			count++
			computed, ok = int_maths.SAdd(V(i), V(j))
			computed16 := int64(computed)
			if c != computed16 {
				if ok {
					errors++
					fmt.Printf("BAD %s ok true  %4d != %4d %4d + %4d\n",
						k, computed16, c, i, j)
				} else {
					goods++
				}
			} else {
				goods++
			}
			if verbose {
				cpt++
				if (cpt & verbose_mask) == verbose_mask {
					fmt.Printf(" + %s: %08xX - \n", k, cpt+1)
				}
			}
		}
	}
	return goods, errors
}

/*
	var goodu16, badu16 = test_add(uint16(0), uint16(0xe00))
	v0.1.1: min time 0.750 s
	when is_signed returns the Kind(), we get 0.380 s
*/

func TestSAddI8(t *testing.T) {
	var goodi8, badi8 = test_add(int8(math.MinInt8), int8(math.MaxInt8))

	fmt.Printf("int8:  goods %5d, bads %5d\n",
		goodi8, badi8)
}

func TestSAddU8(t *testing.T) {
	var goodu8, badu8 = test_add(uint8(0), uint8(math.MaxUint8))

	fmt.Printf("uint8: goods %5d, bads %5d\n",
		goodu8, badu8)
}

func TestSAddI16(t *testing.T) {
	// var goodi16, badi16 = test_add(int16(math.MinInt16), int16(math.MaxInt16))
	var goodi16, badi16 = test_add(int16(-512), int16(512))

	fmt.Printf("int16:  goods %5d, bads %5d\n",
		goodi16, badi16)
}

func TestSAddU16(t *testing.T) {
	// var goodu16, badu16 = test_add(uint16(0), uint16(math.MaxUint16))
	var goodu16, badu16 = test_add(uint16(0), uint16(1024))

	fmt.Printf("uint16:  goods %5d, bads %5d\n",
		goodu16, badu16)
}
