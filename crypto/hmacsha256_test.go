package crypto

import (
	"reflect"
	"testing"
)

func TestHmacSha256(t *testing.T) {
	input := "abc"
	trueHash := "37c1a13ea02b79939287b3846aca7b6f42d05aabd8cbb857f165f564e1cf2476"
	key := "pnKPzUUNjqTMBe6IqJ8WTg=="

	if !reflect.DeepEqual(HmacSha256(input, key), trueHash) {
		t.Errorf("true hash is : %s, but output is %s", trueHash, HmacSha256(input,key))
	}
}