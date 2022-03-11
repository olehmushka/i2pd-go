package alder32

import "hash/adler32"

func Checksum(input []byte) uint32 {
	return adler32.Checksum(input)
}
