package utils

import (
	"sync"

	"github.com/klauspost/compress/zstd"
)

func ClearPrefixZero(d []byte) []byte {

	var idx int

	for i := 31; i > 0; i-- {
		if d[i] == 0x00 {
			continue
		}
		idx = i
	}

	if idx == 0 {
		return []byte{0x00}
	}

	return d[idx:]
}

var (
	EncPool = sync.Pool{
		New: func() interface{} {
			enc, _ := zstd.NewWriter(nil)
			return enc
		},
	}
	DecPool = sync.Pool{
		New: func() interface{} {
			dec, _ := zstd.NewReader(nil)
			return dec
		},
	}
)
