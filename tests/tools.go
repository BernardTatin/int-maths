/*
 * tools.go
 *
 * tests.tools.go
 *
 * project: int-maths
 * by:      bernard
 * created: 2022-06-29
 */

package tests

import (
	"fmt"
	"testing"
)

const i16_cpt = 0x40

func end_of_test(t *testing.T, goods int64, bads int64) {
	message := fmt.Sprintf("%s:  goods %5d, bads %5d", t.Name(), goods, bads)
	if bads != 0 {
		t.Errorf("%s\n", message)
	} else {
		t.Logf("%s\n", message)
	}
}
