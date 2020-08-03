package crypto

import (
	"testing"
)

func TestRsaHash(t *testing.T) {
	input := "abc"
	var publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCBx/R4BhA+GdcLFxvkrX2doXh5
1A7CQF9ojr8fEwyAzNmUS/kxfi8bJgQGeotSXMEmXZhVvA8ZYrIiB0b7/z8gVBQz
CGnf3wpyT9b91XmxQdAYcrB6FogZ57I+YulmE66ewRqPqPargYlPBesL7Wkg7QFq
8fD+dfTxBTqjnBsguwIDAQAB
-----END PUBLIC KEY-----`

	var privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCBx/R4BhA+GdcLFxvkrX2doXh51A7CQF9ojr8fEwyAzNmUS/kx
fi8bJgQGeotSXMEmXZhVvA8ZYrIiB0b7/z8gVBQzCGnf3wpyT9b91XmxQdAYcrB6
FogZ57I+YulmE66ewRqPqPargYlPBesL7Wkg7QFq8fD+dfTxBTqjnBsguwIDAQAB
AoGAakwKMhxVRWgIuC6vS/fFgUx1zuMfS39KJet1ItCJVx1xwIMRkbYWgxf6CAxq
/IS4NuJGGUK040fxOunW/COt4XKLM5C+lpeUSwzK0V/0HH7uWSOIL5BioqloqMgE
rbMsB4MJ2UBc/9OcfSzWBdftbSSXV6TuBBu8eoRNAhaQbFECQQCkVlXe4vmwbtSM
LjlvcBOaG/LZH5h2IJ1YH0al13xOiKr7/E1HtuHNS4IMLUUanj0At5IE+qAxn5oY
Zmp4ymuVAkEAyitarX+6vlszYykJyVWWM16uz0oI5K3J74/UJdAx+ef838CZzESz
Amggz7gZHG6hVexokm94yOUIoWvGvKbHDwJAAqxM1UVH3nDPEECoOoHOL9GIj7ON
3U4GgSpxEb2Bjh12+oZOE36MCbTStrP9zcMJJvvVajNRa7022nLABJr2SQJAPOmn
HpNXjrcn2n6o5jKVWND0Vpx15YC/USDCyMnb5PIS7M+5ByTD0rvQ9wT+++QRVDQw
pb5UZCMt4IT6kiil3QJAfdaunImTA3AFTBOg25k3RyZUlP6jOfiOGRHMv7yLiyvj
GJQQ0iCOBp31lAMwjOWBnmwoWzPMeM6OoExj8ISmaQ==
-----END RSA PRIVATE KEY-----`

	result, err := PrivateEncrypt(input, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	decrypt, err := PublicDecrypt(result, publicKey)
	if err != nil {
		t.Fatal(err)
	}

	if string(decrypt) != input {
		t.Errorf("private encrypt failed want: %s, got  %s", string(input), string(decrypt))
	}

	result2, err := PublicEncrypt(input, publicKey)
	if err != nil {
		t.Fatal(err)
	}

	decrypt2, err := PrivateDecrypt(result2, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	if string(decrypt2) != input {
		t.Errorf("publick encrypt want: %s, got  %s", string(input), string(decrypt2))
	}

	nonce := "XH0S52n5e1RqWiFdPJZVVQtXoQJDFKm7izeZHyk-E2jZS9i8jKJvoesMNPO8wnG3uhHSw66O592vuOJALduI_b5KxezszN3wGq0EdQmwBaqQ32yaJgYBqBhJtdT4sbr-zGZ77hic0ecKjs8xuyBGLhFkWzzwa-mYWqdOKTJcFxM="

	dec, err := PublicDecrypt(nonce, publicKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(dec))

	if string(dec) != "66RyTWoUJWxZCbzp" {
		t.FailNow()
	}
}
