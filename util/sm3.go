package util

import (
	"encoding/hex"
	"hash"

	"github.com/tjfoc/gmsm/sm3"
)

var h hash.Hash

func init() {
	h = sm3.New()
}
func SM3(origin string, salt string) string {
	src := []byte(salt + origin)
	h.Write(src)
	hashed := h.Sum(nil)
	hashString := hex.EncodeToString(hashed)
	return hashString
}
