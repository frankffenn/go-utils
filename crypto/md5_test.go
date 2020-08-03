package crypto

import (
	"reflect"
	"testing"
)

func TestMd5Hash(t *testing.T) {
	input := "abc"
	trueHash := "900150983cd24fb0d6963f7d28e17f72"

	if !reflect.DeepEqual(GenMd5(input), trueHash) {
		t.Errorf("true hash is : %s, but output is %s", trueHash, GenMd5(input))
	}
}
