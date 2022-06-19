package crypto

import (
	"encoding/hex"
)

func DecodeHex(hexString string) string {
	decoded, _ := hex.DecodeString(hexString)

	return string(decoded)
}
