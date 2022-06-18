package crypto

import (
	"encoding/hex"
	"fmt"
)

func DecodeHex(hexString string) string {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error:", err, "make sure you import a hex string")
	}
	return (string(decoded))
}
