package crypto

import (
	"reflect"
	"testing"
)

func TestSha1Hash(t *testing.T) {
	input := "abc"
	trueHash := "a9993e364706816aba3e25717850c26c9cd0d89d"

	if !reflect.DeepEqual(Sha1(input), trueHash) {
		t.Errorf("true hash is : %s, but output is %s", trueHash, Sha1(input))
	}
}
