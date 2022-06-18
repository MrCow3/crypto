package crypto

import (
	"encoding/hex"
	"fmt"
)

func Decode(hexString string) {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error:", err, "make sure you import a hex string")
	}
	fmt.Println(decoded)
}
