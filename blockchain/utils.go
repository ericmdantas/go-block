package blockchain

import "strconv"

func intToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}
