package crypto

import "encoding/hex"

func decode(hexString string) {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	return decoded
}
